package configuration

import (
	"encoding/json"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type JsonParser struct {
	Data []byte
}

// NewJsonParser creates a new json parser as ConfigParserInterface
func NewJsonParser(data []byte) ConfigParserInterface {
	ret := JsonParser{Data: data}
	return &ret
}

func (c JsonParser) Parse() ([]byte, error) {

	if !json.Valid(c.Data) {
		err := errors.New("invalid json")
		rlog.Error("failed to parse json:", err, rlog.Any("data", c.Data))
		return []byte{}, err
	}

	return c.Data, nil
}
