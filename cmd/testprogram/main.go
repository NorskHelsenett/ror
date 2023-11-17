package main

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
)

func main() {
	teststring := stringhelper.RandomString(10, stringhelper.StringTypeAlphaNum)
	stringHelperObject := StringHelperObject{
		Name:         "String helper test",
		RandomString: teststring,
	}

	stringhelper.PrettyprintStruct(stringHelperObject)

	hashSHA512Value := stringhelper.HashSHA512(stringHelperObject.Name, []byte(stringHelperObject.RandomString))
	_, _ = fmt.Printf("Hash (SHA512) value: %s\n", hashSHA512Value)

	md5HashValue := stringhelper.GetMD5Hash([]byte(stringHelperObject.RandomString))
	_, _ = fmt.Printf("Hash (MD5) value: %s\n", md5HashValue)
}

type StringHelperObject struct {
	Name         string `json:"name"`
	RandomString string `json:"test"`
}
