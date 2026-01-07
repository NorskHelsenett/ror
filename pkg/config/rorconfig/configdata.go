package rorconfig

import (
	"math"
	"strconv"
	"time"
)

type ConfigSource string

const (
	ConfigSourceUnknown    ConfigSource = "unknown"
	ConfigSourceEnv        ConfigSource = "env"
	ConfigSourceFlag       ConfigSource = "flag"
	ConfigSourceConfigFile ConfigSource = "configfile"
	ConfigSourceDefault    ConfigSource = "default"
)

func NewConfigSource(source string) ConfigSource {

	switch source {
	case "env":
		return ConfigSourceEnv
	case "flag":
		return ConfigSourceFlag
	case "configfile":
		return ConfigSourceConfigFile
	case "default":
		return ConfigSourceDefault
	default:
		return ConfigSourceUnknown
	}
}

type ConfigData struct {
	Value  string
	source ConfigSource
}

func (cd ConfigData) String() string {
	return string(cd.Value)
}

func (cd ConfigData) Bool() bool {
	b, _ := strconv.ParseBool(cd.Value)
	return b
}
func (cd ConfigData) Int() int {
	i, _ := strconv.Atoi(cd.Value)
	return i
}
func (cd ConfigData) Int64() int64 {
	i, _ := strconv.ParseInt(cd.Value, 10, 64)
	return i
}
func (cd ConfigData) Float64() float64 {
	f, _ := strconv.ParseFloat(cd.Value, 64)
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return 0
	}
	return f
}
func (cd ConfigData) Float32() float32 {
	f, _ := strconv.ParseFloat(cd.Value, 32)
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return 0
	}
	return float32(f)
}
func (cd ConfigData) Uint() uint {
	u, _ := strconv.ParseUint(cd.Value, 10, 64)
	return uint(u)
}
func (cd ConfigData) Uint64() uint64 {
	u, _ := strconv.ParseUint(cd.Value, 10, 64)
	return u
}
func (cd ConfigData) Uint32() uint32 {
	u, _ := strconv.ParseUint(cd.Value, 10, 32)
	return uint32(u)
}

func (cd ConfigData) Time() time.Time {
	t, _ := time.Parse(time.RFC3339Nano, cd.Value)
	return t
}
func (cd ConfigData) TimeDuration() time.Duration {
	t, _ := time.ParseDuration(cd.Value)
	return t
}
