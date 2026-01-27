package idhelper

import (
	"crypto/rand"
	"fmt"
)

// GetIdentifier returns a identifier from a name
//
// Parameters:
//
// - name: name to be turned into an identifier
func GetIdentifier(name string) string {
	idpostfix := RandomLowerAlphaNum(4)
	identifier := fmt.Sprintf("%s-%s", name, idpostfix)
	return identifier
}

// RandomLowerAlphaNum returns a random lowercase alphanumeric string with the requested length
func RandomLowerAlphaNum(length int) string {
	if length <= 0 {
		return ""
	}

	const dictionary = "0123456789abcdefghijklmnopqrstuvwxyz"

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return ""
	}

	for i, b := range buf {
		buf[i] = dictionary[int(b)%len(dictionary)]
	}

	return string(buf)
}
