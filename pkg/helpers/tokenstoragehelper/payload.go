package tokenstoragehelper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
)

// SignPayload signs exact payload bytes as compact JWS, without adding JWT
// claims. It uses the singleton signing store also exposed through GetJwks.
func SignPayload(payload []byte, mediaType string) (string, error) {
	return keyStorage.SignPayload(payload, mediaType)
}

// SignPayload uses the active RSA key and protects alg, kid and typ.
// It does not interpret or change payload bytes.
func (k *KeyStorageProvider) SignPayload(payload []byte, mediaType string) (string, error) {
	if k == nil {
		return "", fmt.Errorf("no signing keys available")
	}
	key := k.getCurrentKey()
	if key.KeyID == "" || key.PrivateKey == nil || mediaType == "" ||
		!utf8.ValidString(key.KeyID) || !utf8.ValidString(mediaType) {
		return "", fmt.Errorf("signing requires a key ID, private key and media type")
	}
	if err := key.PrivateKey.Validate(); err != nil {
		return "", fmt.Errorf("invalid payload signing key: %w", err)
	}
	var method *jwt.SigningMethodRSA
	switch key.AlgorithmKey {
	case "RS256":
		method = jwt.SigningMethodRS256
	case "RS512":
		method = jwt.SigningMethodRS512
	default:
		return "", fmt.Errorf("unsupported payload signing algorithm %q", key.AlgorithmKey)
	}
	header, err := json.Marshal(struct {
		Algorithm string `json:"alg"`
		KeyID     string `json:"kid"`
		Type      string `json:"typ"`
	}{method.Alg(), key.KeyID, mediaType})
	if err != nil {
		return "", fmt.Errorf("encode protected signing header: %w", err)
	}
	input := base64.RawURLEncoding.EncodeToString(header) + "." + base64.RawURLEncoding.EncodeToString(payload)
	signature, err := method.Sign(input, key.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("sign payload: %w", err)
	}
	return input + "." + base64.RawURLEncoding.EncodeToString(signature), nil
}
