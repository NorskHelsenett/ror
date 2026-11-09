package acl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/v2/acl"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclsnapshot"
)

func newSnapshotClient(t *testing.T, handler http.HandlerFunc) *acl.V2Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg, err := httpclient.NewHttpTransportClientConfig(server.URL,
		httpauthprovider.NewAuthProvider(httpauthprovider.AuthPoviderTypeAPIKey, "key"), "test", rorversion.GetRorVersion())
	if err != nil {
		t.Fatal(err)
	}
	return acl.NewV2Client(httpclient.NewHttpTransportClient(&http.Client{}, cfg, httpclient.NewHttpTransportClientStatus()))
}

func TestGetSnapshot(t *testing.T) {
	client := newSnapshotClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/acl/snapshot" || r.Header.Get("Accept") != aclsnapshot.ContentType || r.Header.Get("X-API-KEY") != "key" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		w.Header().Set("ETag", `"new"`)
		if r.Header.Get("If-None-Match") == `"new"` {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("Content-Type", "application/jose; charset=utf-8")
		_, _ = w.Write([]byte("a.b.c"))
	})
	limits := aclsnapshot.DefaultLimits()
	res, err := client.GetSnapshot(context.Background(), "", limits)
	if err != nil || res.NotModified || string(res.Envelope) != "a.b.c" || res.ETag != `"new"` {
		t.Fatalf("unexpected response %+v, %v", res, err)
	}
	res, err = client.GetSnapshot(context.Background(), `"new"`, limits)
	if err != nil || !res.NotModified || len(res.Envelope) != 0 || res.ETag != `"new"` {
		t.Fatalf("unexpected not-modified response %+v, %v", res, err)
	}
}

func TestGetSnapshotRejectsInvalidResponses(t *testing.T) {
	for name, handler := range map[string]http.HandlerFunc{
		"missing etag": func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", aclsnapshot.ContentType)
			_, _ = w.Write([]byte("a.b.c"))
		},
		"wrong content type": func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("ETag", `"x"`)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("{}"))
		},
		"too large": func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("ETag", `"x"`)
			w.Header().Set("Content-Type", aclsnapshot.ContentType)
			_, _ = w.Write([]byte(strings.Repeat("a", 2049)))
		},
		"server error": func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("ETag", `"x"`)
			http.Error(w, "down", http.StatusInternalServerError)
		},
	} {
		t.Run(name, func(t *testing.T) {
			limits := aclsnapshot.DefaultLimits()
			limits.MaxEnvelopeBytes = 2048
			limits.MaxPayloadBytes = 1024
			if _, err := newSnapshotClient(t, handler).GetSnapshot(context.Background(), "", limits); err == nil {
				t.Fatal("expected error")
			}
		})
	}
	if _, err := newSnapshotClient(t, nil).GetSnapshot(context.Background(), "", aclsnapshot.Limits{}); err == nil {
		t.Fatal("expected invalid limits error")
	}
}
