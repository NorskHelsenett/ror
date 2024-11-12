package configuration

import (
	"encoding/json"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	k8syaml "sigs.k8s.io/yaml"
)

type YamlParser struct {
	Data []byte
}

// NewYamlParser creates a new yaml parser as ConfigParserInterface
func NewYamlParser(data []byte) ConfigParserInterface {
	ret := YamlParser{Data: data}
	return &ret
}

func (c YamlParser) Parse() ([]byte, error) {

	jsonBytes, err := k8syaml.YAMLToJSON(c.Data)

	if err != nil {
		rlog.Error("failed to parse yaml:", err, rlog.Any("data", c.Data))
		return []byte{}, err
	}

	if !json.Valid(jsonBytes) {
		return nil, errors.New("invalid converted json")
	}

	return jsonBytes, err
}
