package aclsnapshot

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/lestrrat-go/jwx/v4/jwk"
	"github.com/lestrrat-go/jwx/v4/jws"
)

// Algorithm is the only accepted snapshot signature algorithm.
const Algorithm = "RS512"

// Header is the exact protected header of a signed snapshot.
type Header struct {
	Algorithm string
	KeyID     string
	Type      string
}

// Verified is a snapshot whose signature and payload have been verified.
type Verified struct {
	Payload Payload
	// Bytes are the exact signed payload bytes.
	Bytes  []byte
	Digest [sha256.Size]byte
	KeyID  string
}

var strictBase64 = base64.RawURLEncoding.Strict()

// ParseHeader validates the compact JWS form and protected header of a signed
// snapshot without verifying its signature. Use it to select verification keys
// or to check producer output; never treat its result as trusted policy.
func ParseHeader(envelope []byte, l Limits) (Header, error) {
	if err := l.Validate(); err != nil {
		return Header{}, err
	}
	if len(envelope) > l.MaxEnvelopeBytes {
		return Header{}, fmt.Errorf("signed snapshot exceeds envelope byte bound")
	}
	parts := bytes.Split(envelope, []byte("."))
	if len(parts) != 3 || len(parts[0]) == 0 || len(parts[1]) == 0 || len(parts[2]) == 0 {
		return Header{}, fmt.Errorf("signed snapshot must be compact JWS")
	}
	if len(parts[0]) > strictBase64.EncodedLen(l.MaxHeaderBytes) {
		return Header{}, fmt.Errorf("protected header exceeds byte bound")
	}
	if len(parts[1]) > strictBase64.EncodedLen(l.MaxPayloadBytes) {
		return Header{}, fmt.Errorf("snapshot payload exceeds byte bound")
	}
	raw := make([]byte, strictBase64.DecodedLen(len(parts[0])))
	n, err := strictBase64.Decode(raw, parts[0])
	if err != nil {
		return Header{}, fmt.Errorf("invalid protected header encoding")
	}
	header, err := decodeHeader(raw[:n])
	if err != nil {
		return Header{}, err
	}
	if header.Algorithm != Algorithm || header.Type != MediaType ||
		header.KeyID == "" || len(header.KeyID) > l.MaxFieldBytes {
		return Header{}, fmt.Errorf("signing header does not match snapshot contract")
	}
	return header, nil
}

// decodeHeader accepts only one JSON object with unique alg, kid and typ strings.
func decodeHeader(data []byte) (Header, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if token, err := decoder.Token(); err != nil || token != json.Delim('{') {
		return Header{}, fmt.Errorf("protected header must be a JSON object")
	}
	var header Header
	seen := map[string]bool{}
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return Header{}, fmt.Errorf("decode protected header: %w", err)
		}
		name, _ := token.(string)
		var target *string
		switch name {
		case "alg":
			target = &header.Algorithm
		case "kid":
			target = &header.KeyID
		case "typ":
			target = &header.Type
		default:
			return Header{}, fmt.Errorf("unsupported protected header field %q", name)
		}
		if seen[name] {
			return Header{}, fmt.Errorf("duplicate protected header field %q", name)
		}
		seen[name] = true
		if err := decoder.Decode(target); err != nil {
			return Header{}, fmt.Errorf("protected header field %q must be a string", name)
		}
	}
	if token, err := decoder.Token(); err != nil || token != json.Delim('}') {
		return Header{}, fmt.Errorf("protected header must be a JSON object")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return Header{}, fmt.Errorf("protected header has trailing data")
	}
	return header, nil
}

// Verify checks a signed snapshot against public verification keys and returns
// its decoded payload. The key set must contain the signing key with a matching
// kid and alg. Unverified content is never returned.
func Verify(envelope []byte, keys jwk.Set, l Limits) (Verified, error) {
	header, err := ParseHeader(envelope, l)
	if err != nil {
		return Verified{}, err
	}
	if keys == nil {
		return Verified{}, fmt.Errorf("verification keys are required")
	}
	if _, ok := keys.LookupKeyID(header.KeyID); !ok {
		return Verified{}, fmt.Errorf("unknown snapshot signing key %q", header.KeyID)
	}
	payload, err := jws.Verify(envelope, jws.WithCompact(), jws.WithKeySet(keys))
	if err != nil {
		return Verified{}, fmt.Errorf("verify snapshot signature: %w", err)
	}
	if len(payload) > l.MaxPayloadBytes {
		return Verified{}, fmt.Errorf("snapshot payload exceeds byte bound")
	}
	decoded, err := Decode(payload, l)
	if err != nil {
		return Verified{}, err
	}
	return Verified{
		Payload: decoded,
		Bytes:   payload,
		Digest:  sha256.Sum256(payload),
		KeyID:   header.KeyID,
	}, nil
}
