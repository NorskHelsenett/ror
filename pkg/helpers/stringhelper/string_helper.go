package stringhelper

import (
	"crypto/md5" // #nosec G501 - MD5 is used for hashing, not for encryption
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

type StringType = string

const (
	StringTypeUnknown   StringType = "Unknown"
	StringTypeAlphaNum  StringType = "alphanum"
	StringTypeClusterId StringType = "clusterid"
	StringTypeAlpha     StringType = "alpha"
	StringTypeNumber    StringType = "number"
)

// RandomString returns a random string of the specified length
func RandomString(strSize int, randType StringType) string {
	var dictionary string

	if randType == StringTypeUnknown {
		return ""
	}

	if randType == StringTypeAlphaNum {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == StringTypeClusterId {
		dictionary = "0123456789abcdefghijklmnopqrstuvwxyz"
	}

	if randType == StringTypeAlpha {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == StringTypeNumber {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	_, _ = rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}

// Deprecated: Use GetSHA256Hash, GetSHA384Hash or GetSHA512Hash instead
// GetMD5Hash returns the md5 hash of a byte array
// WARNING: This function is not safe for hashing passwords
func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data) // #nosec G401 - MD5 is used for hashing, not for encryption
	return hex.EncodeToString(hash[:])
}

// GetSHA256Hash returns the SHA256 hash of a byte array
func GetSHA256Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// GetSHA384Hash returns the SHA384 hash of a byte array
func GetSHA384Hash(data []byte) string {
	hash := sha512.Sum384(data)
	return hex.EncodeToString(hash[:])
}

// GetSHA512Hash returns the SHA512 hash of a byte array
func GetSHA512Hash(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// JsonToMap converts a json string to a map
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}

// CompareLabels compares two maps of labels
func CompareLabels(search map[string]string, labels map[string]string) bool {
	i := 0
	for k, v := range search {
		if labels[k] == v {
			i = i + 1
		}
	}
	return i == len(search)
}

// PrettyprintStruct prints a struct in a pretty way
func PrettyprintStruct(obj interface{}) {
	empJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalf(err.Error(), nil)
	}
	_, err = fmt.Printf("MarshalIndent function output %s\n", string(empJSON))
	if err != nil {
		panic(err)
	}
}

// Hash a password with the SHA-512 algorithm
func HashSHA512(text string, salt []byte) string {
	// Convert password string to byte slice
	var textByte = []byte(text)

	// Create sha-512 hasher
	var sha512hasher = sha512.New()

	textByte = append(textByte, salt...)

	_, err := sha512hasher.Write(textByte)
	if err != nil {
		return ""
	}

	// Get the SHA-512 hashed password
	var hashed = sha512hasher.Sum(nil)

	// Convert the hashed to hex string
	var hashedHex = hex.EncodeToString(hashed)
	return hashedHex
}

// Check if two passwords match
func DoHashsMatchWithString(hash, text string, salt []byte) bool {
	var textHash = HashSHA512(text, salt)
	return hash == textHash
}

func EscapeString(str string) string {
	return fmt.Sprintf("%q", str)
}
