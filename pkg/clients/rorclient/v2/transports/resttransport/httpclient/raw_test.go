package httpclient_test

import (
	"compress/gzip"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

func newRawClient(t *testing.T, handler http.HandlerFunc) *httpclient.HttpTransportClient {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg, err := httpclient.NewHttpTransportClientConfig(server.URL,
		httpauthprovider.NewAuthProvider(httpauthprovider.AuthPoviderTypeAPIKey, "test-key"), "test", rorversion.GetRorVersion())
	if err != nil {
		t.Fatal(err)
	}
	return httpclient.NewHttpTransportClient(&http.Client{}, cfg, httpclient.NewHttpTransportClientStatus())
}

func maxBytes(n int64) httpclient.HttpTransportClientParams {
	return httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsMaxBytes, Value: n}
}

func TestGetRawReturnsBodyAndHeaders(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "test-key" || r.Header.Values("Accept")[0] != "application/jose" ||
			len(r.Header.Values("Accept")) != 1 || r.Header.Get("Content-Type") != "" || r.Header.Get("If-None-Match") != `"a"` {
			http.Error(w, "unexpected request headers", http.StatusBadRequest)
			return
		}
		w.Header().Set("ETag", `"b"`)
		_, _ = w.Write([]byte("abc"))
	})
	res, err := client.GetRaw(context.Background(), "/raw", maxBytes(3), httpclient.HttpTransportClientParams{
		Key: httpclient.HttpTransportClientOptsHeaders, Value: map[string]string{"accept": "application/jose", "If-None-Match": `"a"`},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK || string(res.Body) != "abc" || res.Header.Get("ETag") != `"b"` {
		t.Fatalf("unexpected response: %+v", res)
	}
	if !client.Status.Established {
		t.Fatal("postflight status was not updated")
	}
}

func TestGetRawNoAuthAndDefaultAccept(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "" || r.Header.Get("Accept") != "application/json" {
			http.Error(w, "unexpected request headers", http.StatusBadRequest)
			return
		}
		_, _ = w.Write([]byte("{}"))
	})
	if _, err := client.GetRaw(context.Background(), "/raw",
		httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsNoAuth}); err != nil {
		t.Fatal(err)
	}
}

func TestGetRawNotModified(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", `"a"`)
		w.WriteHeader(http.StatusNotModified)
	})
	res, err := client.GetRaw(context.Background(), "/raw", maxBytes(1))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotModified || len(res.Body) != 0 || res.Header.Get("ETag") != `"a"` {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func TestGetRawBodyLimit(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("abcd"))
	})
	if _, err := client.GetRaw(context.Background(), "/raw", maxBytes(3)); err == nil || !strings.Contains(err.Error(), "exceeds") {
		t.Fatalf("expected body limit error, got %v", err)
	}
	if _, err := client.GetRaw(context.Background(), "/raw", maxBytes(0)); err == nil {
		t.Fatal("expected non-positive limit to be rejected")
	}
	if _, err := client.GetRaw(context.Background(), "/raw",
		httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsMaxBytes, Value: 3}); err == nil {
		t.Fatal("expected non-int64 limit to be rejected")
	}
}

func TestGetRawDefaultBodyLimit(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(make([]byte, httpclient.DefaultRawMaxBytes+1))
	})
	if _, err := client.GetRaw(context.Background(), "/raw"); err == nil || !strings.Contains(err.Error(), "exceeds") {
		t.Fatalf("expected default body limit error, got %v", err)
	}
	if _, err := client.GetRaw(context.Background(), "/raw", maxBytes(httpclient.DefaultRawMaxBytes+1)); err != nil {
		t.Fatal(err)
	}
}

func TestGetRawDecodesGzip(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			_, _ = w.Write([]byte("plain"))
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		zw := gzip.NewWriter(w)
		_, _ = zw.Write([]byte("decoded"))
		_ = zw.Close()
	})
	res, err := client.GetRaw(context.Background(), "/raw", maxBytes(7))
	if err != nil {
		t.Fatal(err)
	}
	if string(res.Body) != "decoded" {
		t.Fatalf("expected transparently decoded body, got %q", res.Body)
	}
}

func TestGetRawErrorsAndRetryAfter(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "120")
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	if _, err := client.GetRaw(context.Background(), "/raw"); err == nil {
		t.Fatal("expected error for 503")
	}
	if client.Status.RetryAfter.IsZero() {
		t.Fatal("retry-after was not recorded")
	}
	if _, err := client.GetRaw(context.Background(), "/raw"); err == nil || !strings.Contains(err.Error(), "preflight") {
		t.Fatalf("expected preflight error, got %v", err)
	}
}
