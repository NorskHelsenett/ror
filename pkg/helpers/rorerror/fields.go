package rorerror

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type Field struct {
	Key   string
	Value string
}

func (f Field) ToRlog() rlog.Field {
	return rlog.String(f.Key, f.Value)
}

// Field functions

func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: fmt.Sprintf("%d", value)}
}
func ByteString(key string, value []byte) Field {
	return Field{Key: key, Value: string(value)}
}
func Int64(key string, value int64) Field {
	return Field{Key: key, Value: fmt.Sprintf("%d", value)}
}

func Uint(key string, value uint) Field {
	return Field{Key: key, Value: fmt.Sprintf("%d", value)}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Value: fmt.Sprintf("%f", value)}
}
func Float32(key string, value float32) Field {
	return Field{Key: key, Value: fmt.Sprintf("%f", value)}
}

func Stringp(key string, value *string) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: *value}
}

func Intp(key string, value *int) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: fmt.Sprintf("%d", *value)}
}
func ByteStringp(key string, value *[]byte) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: string(*value)}
}
func Int64p(key string, value *int64) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: fmt.Sprintf("%d", *value)}
}

func Uintp(key string, value *uint) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: fmt.Sprintf("%d", *value)}
}

func Float64p(key string, value *float64) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: fmt.Sprintf("%f", *value)}
}
func Float32p(key string, value *float32) Field {
	if value == nil {
		return Field{Key: key, Value: "nil"}
	}
	return Field{Key: key, Value: fmt.Sprintf("%f", *value)}
}
