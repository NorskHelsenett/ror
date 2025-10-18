package rorconfig

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/joho/godotenv"
)

func readDotEnv() map[string]string {
	enfilevar := strings.Split(os.Getenv("ENV_FILE"), ",")
	if len(enfilevar) == 1 && enfilevar[0] == "" {
		enfilevar = []string{".env"}
	}

	dotenvs, _ := godotenv.Read(enfilevar...)
	return dotenvs
}

func anyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case uint:
		return fmt.Sprintf("%d", v)
	case uint64:
		return fmt.Sprintf("%d", v)
	case uint32:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		rlog.Error("Unsupported type for config value", fmt.Errorf("type %T not supported", v))
		return ""
	}
}
