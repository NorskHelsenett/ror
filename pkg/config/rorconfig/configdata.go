package rorconfig

import (
	"math"
	"strconv"
)

type ConfigData string

func (cd ConfigData) String() string {
	return string(cd)
}

func (cd ConfigData) Bool() bool {
	b, _ := strconv.ParseBool(cd.String())
	return b
}
func (cd ConfigData) Int() int {
	i, _ := strconv.Atoi(cd.String())
	return i
}
func (cd ConfigData) Int64() int64 {
	i, _ := strconv.ParseInt(cd.String(), 10, 64)
	return i
}
func (cd ConfigData) Float64() float64 {
	f, _ := strconv.ParseFloat(cd.String(), 64)
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return 0
	}
	return f
}
func (cd ConfigData) Float32() float32 {
	f, _ := strconv.ParseFloat(cd.String(), 32)
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return 0
	}
	return float32(f)
}
func (cd ConfigData) Uint() uint {
	u, _ := strconv.ParseUint(cd.String(), 10, 64)
	return uint(u)
}
func (cd ConfigData) Uint64() uint64 {
	u, _ := strconv.ParseUint(cd.String(), 10, 64)
	return u
}
func (cd ConfigData) Uint32() uint32 {
	u, _ := strconv.ParseUint(cd.String(), 10, 32)
	return uint32(u)
}
