package stringhelper

import (
	"crypto/md5"
	"crypto/rand"
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

// GetMD5Hash returns the md5 hash of a byte array
func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data)
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
		log.Fatalf(err.Error())
	}
	_, err = fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))
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
