package jsonhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

// StringToJson converts a string to a json byte array.
func StringToJson(key string, value string) []byte {
	keys := strings.Split(key, ".")
	data := make(map[string]interface{})

	currentMap := data
	for i, key := range keys {
		if len(key) == 0 {
			continue
		}
		if i == len(keys)-1 {
			currentMap[key] = value
		} else {
			newMap := make(map[string]interface{})
			currentMap[key] = newMap
			currentMap = newMap
		}
	}
	retjson, err := json.Marshal(data)
	if err != nil {
		return []byte("{}")
	}
	return retjson
}

// MapToJson converts a map to a json byte array.
func MapToJson(input map[string]string) []byte {
	jsonret := []byte("{}")

	for key, value := range input {
		jsonpart := StringToJson(key, value)
		jsonret, _ = jsonpatch.MergePatch(jsonret, jsonpart)
	}
	return jsonret
}

// PrettyPrintJson prints the json byte array to the stdout formated.
func PrettyPrintJson(jsonByte []byte) error {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, jsonByte, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(prettyJson.String())
	return nil
}
