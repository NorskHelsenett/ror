package mapping

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Map[F any, D any](from F, destination *D) error {

	jsonByte, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("could not convert to json, from type %s: %v", reflect.TypeOf(from), err)
	}

	if err := json.Unmarshal(jsonByte, &destination); err != nil {
		return fmt.Errorf("could not convert from json to %s: %v", reflect.TypeOf(destination), err)
	}

	return nil
}

func InterfaceToInt64(i interface{}) (int64, error) {
	switch v := i.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, errors.New("type error")
	}
}
