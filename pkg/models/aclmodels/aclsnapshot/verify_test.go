package aclsnapshot_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclgrants"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclsnapshot"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v4/jwk"
)

var testKeys = map[string]*rsa.PrivateKey{}

func testKey(t *testing.T, name string) *rsa.PrivateKey {
	t.Helper()
	if key, ok := testKeys[name]; ok {
		return key
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	testKeys[name] = key
	return key
}

func testSigner(t *testing.T, kid string, alg string) *tokenstoragehelper.KeyStorageProvider {
	t.Helper()
	return &tokenstoragehelper.KeyStorageProvider{Keys: map[int]tokenstoragehelper.Key{
		1: {KeyID: kid, PrivateKey: testKey(t, kid), AlgorithmKey: alg},
	}}
}

func testPayload(t *testing.T) []byte {
	t.Helper()
	payload, err := aclsnapshot.Encode(aclsnapshot.Payload{
		Schema:      aclsnapshot.Schema,
		Revision:    7,
		GeneratedAt: time.Date(2026, 11, 9, 10, 0, 0, 0, time.UTC),
		Rows: []aclsnapshot.Row{{Group: "readers@ror.system", Grant: aclgrants.Grant{
			Scope: "ror", Subject: "acl", Access: []aclcaps.AccessTypeV3{"ror:acl:snapshot:read"},
		}}},
	}, aclsnapshot.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	return payload
}

func sign(t *testing.T, signer *tokenstoragehelper.KeyStorageProvider, payload []byte, mediaType string) []byte {
	t.Helper()
	envelope, err := signer.SignPayload(payload, mediaType)
	if err != nil {
		t.Fatal(err)
	}
	return []byte(envelope)
}

func publicKeys(t *testing.T, signers ...*tokenstoragehelper.KeyStorageProvider) jwk.Set {
	t.Helper()
	set := jwk.NewSet()
	for _, signer := range signers {
		keys, err := signer.GetJwks()
		if err != nil {
			t.Fatal(err)
		}
		for _, key := range keys.All() {
			if err := set.AddKey(key); err != nil {
				t.Fatal(err)
			}
		}
	}
	return set
}

func TestVerifyAcceptsProducerSignature(t *testing.T) {
	signer := testSigner(t, "kid-a", "RS512")
	payload := testPayload(t)
	verified, err := aclsnapshot.Verify(sign(t, signer, payload, aclsnapshot.MediaType),
		publicKeys(t, signer), aclsnapshot.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	if string(verified.Bytes) != string(payload) || verified.Digest != sha256.Sum256(payload) ||
		verified.KeyID != "kid-a" || verified.Payload.Revision != 7 || len(verified.Payload.Rows) != 1 {
		t.Fatalf("unexpected verified snapshot: %+v", verified)
	}
}

func TestVerifyRejectsUntrustedOrAlteredSnapshots(t *testing.T) {
	limits := aclsnapshot.DefaultLimits()
	signer := testSigner(t, "kid-a", "RS512")
	other := testSigner(t, "kid-b", "RS512")
	payload := testPayload(t)
	valid := sign(t, signer, payload, aclsnapshot.MediaType)
	parts := strings.Split(string(valid), ".")
	header := func(json string) string { return base64.RawURLEncoding.EncodeToString([]byte(json)) }
	impostor := &tokenstoragehelper.KeyStorageProvider{Keys: map[int]tokenstoragehelper.Key{
		1: {KeyID: "kid-a", PrivateKey: testKey(t, "kid-b"), AlgorithmKey: "RS512"},
	}}
	hs512 := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{})
	hs512.Header["kid"] = "kid-a"
	hs512.Header["typ"] = aclsnapshot.MediaType
	hsToken, err := hs512.SignedString([]byte("secret"))
	if err != nil {
		t.Fatal(err)
	}
	invalidPayload := sign(t, signer, []byte(`{"schema":"other"}`), aclsnapshot.MediaType)

	for name, envelope := range map[string]string{
		"wrong key same kid":           string(sign(t, impostor, payload, aclsnapshot.MediaType)),
		"unknown kid":                  string(sign(t, other, payload, aclsnapshot.MediaType)),
		"rs256":                        string(sign(t, testSigner(t, "kid-a", "RS256"), payload, aclsnapshot.MediaType)),
		"hs512":                        hsToken,
		"wrong typ":                    string(sign(t, signer, payload, "application/jose")),
		"altered payload":              parts[0] + "." + base64.RawURLEncoding.EncodeToString(append(payload, ' ')) + "." + parts[2],
		"altered signature":            parts[0] + "." + parts[1] + "." + strings.Repeat("A", len(parts[2])),
		"missing kid":                  header(`{"alg":"RS512","typ":"`+aclsnapshot.MediaType+`"}`) + "." + parts[1] + "." + parts[2],
		"none":                         header(`{"alg":"none","kid":"kid-a","typ":"`+aclsnapshot.MediaType+`"}`) + "." + parts[1] + ".",
		"extra header":                 header(`{"alg":"RS512","kid":"kid-a","typ":"`+aclsnapshot.MediaType+`","crit":["x"]}`) + "." + parts[1] + "." + parts[2],
		"duplicate alg":                header(`{"alg":"none","alg":"RS512","kid":"kid-a","typ":"`+aclsnapshot.MediaType+`"}`) + "." + parts[1] + "." + parts[2],
		"trailing header data":         header(`{"alg":"RS512","kid":"kid-a","typ":"`+aclsnapshot.MediaType+`"}{}`) + "." + parts[1] + "." + parts[2],
		"padded header":                parts[0] + "=." + parts[1] + "." + parts[2],
		"json serialization":           `{"payload":"` + parts[1] + `","protected":"` + parts[0] + `","signature":"` + parts[2] + `"}`,
		"unsigned":                     parts[0] + "." + parts[1],
		"valid signature, bad payload": string(invalidPayload),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := aclsnapshot.Verify([]byte(envelope), publicKeys(t, signer), limits); err == nil {
				t.Fatal("expected verification failure")
			}
		})
	}
	if _, err := aclsnapshot.Verify(valid, nil, limits); err == nil {
		t.Fatal("expected missing key set to fail")
	}
}

func TestVerifyBounds(t *testing.T) {
	signer := testSigner(t, "kid-a", "RS512")
	payload := testPayload(t)
	envelope := sign(t, signer, payload, aclsnapshot.MediaType)
	keys := publicKeys(t, signer)
	header, err := aclsnapshot.ParseHeader(envelope, aclsnapshot.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	headerBytes := len(`{"alg":"RS512","kid":"kid-a","typ":"` + aclsnapshot.MediaType + `"}`)

	for name, tt := range map[string]struct {
		limits func(*aclsnapshot.Limits)
		ok     bool
	}{
		"envelope at limit":    {func(l *aclsnapshot.Limits) { l.MaxEnvelopeBytes = len(envelope) }, true},
		"envelope over limit":  {func(l *aclsnapshot.Limits) { l.MaxEnvelopeBytes = len(envelope) - 1 }, false},
		"payload at limit":     {func(l *aclsnapshot.Limits) { l.MaxPayloadBytes = len(payload) }, true},
		"payload over limit":   {func(l *aclsnapshot.Limits) { l.MaxPayloadBytes = len(payload) - 1 }, false},
		"header at limit":      {func(l *aclsnapshot.Limits) { l.MaxHeaderBytes = headerBytes }, true},
		"header over limit":    {func(l *aclsnapshot.Limits) { l.MaxHeaderBytes = headerBytes - 3 }, false},
		"kid at field limit":   {func(l *aclsnapshot.Limits) { l.MaxFieldBytes = len(header.KeyID) + 20 }, true},
		"kid over field limit": {func(l *aclsnapshot.Limits) { l.MaxFieldBytes = len(header.KeyID) - 1 }, false},
		"invalid limits":       {func(l *aclsnapshot.Limits) { l.MaxRows = 0 }, false},
	} {
		t.Run(name, func(t *testing.T) {
			limits := aclsnapshot.DefaultLimits()
			tt.limits(&limits)
			_, err := aclsnapshot.Verify(envelope, keys, limits)
			if (err == nil) != tt.ok {
				t.Fatalf("verify error = %v, want ok %v", err, tt.ok)
			}
		})
	}
}

func TestDefaultLimitsAreValid(t *testing.T) {
	if err := aclsnapshot.DefaultLimits().Validate(); err != nil {
		t.Fatal(err)
	}
}
