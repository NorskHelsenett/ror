package tokenstoragehelper

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"strings"
	"sync"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestSignPayload(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, method := range []*jwt.SigningMethodRSA{jwt.SigningMethodRS256, jwt.SigningMethodRS512} {
		t.Run(method.Alg(), func(t *testing.T) {
			storage := &KeyStorageProvider{Keys: map[int]Key{1: {
				KeyID: "test-key", PrivateKey: privateKey, AlgorithmKey: method.Alg(),
			}}}
			payload := []byte("{ \"rows\": [] }\n")
			signed, err := storage.SignPayload(payload, "application/test+jws")
			if err != nil {
				t.Fatal(err)
			}
			parts := strings.Split(signed, ".")
			decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
			if err != nil || string(decoded) != string(payload) {
				t.Fatal("payload bytes changed")
			}
			header, _ := base64.RawURLEncoding.DecodeString(parts[0])
			var fields map[string]string
			if err := json.Unmarshal(header, &fields); err != nil {
				t.Fatal(err)
			}
			if len(fields) != 3 || fields["alg"] != method.Alg() || fields["kid"] != "test-key" || fields["typ"] != "application/test+jws" {
				t.Fatalf("unexpected protected header %s", header)
			}
			signature, _ := base64.RawURLEncoding.DecodeString(parts[2])
			input := strings.Join(parts[:2], ".")
			if err := method.Verify(input, signature, &privateKey.PublicKey); err != nil {
				t.Fatal(err)
			}
			if err := method.Verify(input+"x", signature, &privateKey.PublicKey); err == nil {
				t.Fatal("tampered payload accepted")
			}
		})
	}
	previous := keyStorage
	t.Cleanup(func() { keyStorage = previous })
	keyStorage = KeyStorageProvider{Keys: map[int]Key{1: {KeyID: "singleton", PrivateKey: privateKey, AlgorithmKey: "RS512"}}}
	if _, err := SignPayload([]byte("exact bytes"), "application/test"); err != nil {
		t.Fatal(err)
	}
}

func TestSignPayloadDuringRotation(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	key := Key{KeyID: "rotation-test", PrivateKey: privateKey, AlgorithmKey: "RS512"}
	previous := generateKeyFn
	generateKeyFn = func() (Key, error) { return key, nil }
	t.Cleanup(func() { generateKeyFn = previous })
	storage := &KeyStorageProvider{
		NumKeys: 3, Keys: map[int]Key{0: key, 1: key, 2: key},
	}
	var wg sync.WaitGroup
	errs := make(chan error, 60)
	for range 20 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			storage.Rotate(true)
		}()
		go func() {
			defer wg.Done()
			if _, err := storage.SignPayload([]byte("{}"), "application/test"); err != nil {
				errs <- err
			}
		}()
		go func() {
			defer wg.Done()
			if _, err := storage.GetJwks(); err != nil {
				errs <- err
			}
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Error(err)
	}
}

func TestSignPayloadFailures(t *testing.T) {
	for _, storage := range []*KeyStorageProvider{
		nil, {}, {Keys: map[int]Key{1: {KeyID: "missing-private-key"}}},
		{Keys: map[int]Key{2: {KeyID: "no-active-key"}}},
	} {
		if _, err := storage.SignPayload([]byte("{}"), "application/test"); err == nil {
			t.Fatal("invalid storage accepted")
		}
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, algorithm := range []string{"none", "HS256", "PS512", ""} {
		storage := &KeyStorageProvider{Keys: map[int]Key{1: {KeyID: "key", PrivateKey: privateKey, AlgorithmKey: algorithm}}}
		if _, err := storage.SignPayload([]byte("{}"), "application/test"); err == nil {
			t.Fatalf("accepted %s", algorithm)
		}
	}
	storage := &KeyStorageProvider{Keys: map[int]Key{1: {KeyID: "key", PrivateKey: privateKey, AlgorithmKey: "RS512"}}}
	if _, err := storage.SignPayload([]byte("{}"), ""); err == nil {
		t.Fatal("missing media type accepted")
	}
	for _, invalid := range []Key{
		{KeyID: "\xff", PrivateKey: privateKey, AlgorithmKey: "RS512"},
		{KeyID: "key", PrivateKey: &rsa.PrivateKey{PublicKey: rsa.PublicKey{N: big.NewInt(15), E: 3}}, AlgorithmKey: "RS512"},
	} {
		storage.Keys[1] = invalid
		if _, err := storage.SignPayload([]byte("{}"), "application/test"); err == nil {
			t.Fatal("invalid signing key accepted")
		}
	}
}
