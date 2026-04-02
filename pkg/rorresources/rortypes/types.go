package rortypes

import (
	"fmt"
	"math"

	"go.mongodb.org/mongo-driver/v2/x/bsonx/bsoncore"
	kubernetesresource "k8s.io/apimachinery/pkg/api/resource"
)

// Quantity embeds resource.Quantity to inherit all its methods while adding
// custom BSON marshaling so MongoDB stores it as a string instead of a struct.
type Quantity struct {
	kubernetesresource.Quantity
}

func (q Quantity) MarshalBSONValue() (byte, []byte, error) {
	return byte(bsoncore.TypeString), bsoncore.AppendString(nil, q.String()), nil
}

func (q *Quantity) UnmarshalBSONValue(t byte, data []byte) error {
	if bsoncore.Type(t) != bsoncore.TypeString {
		return fmt.Errorf("cannot unmarshal BSON type %v into Quantity", t)
	}
	str, _, ok := bsoncore.ReadString(data)
	if !ok {
		return fmt.Errorf("failed to read BSON string for Quantity")
	}
	rq, err := kubernetesresource.ParseQuantity(str)
	if err != nil {
		return err
	}
	q.Quantity = rq
	return nil
}

type BinarySIUnit string

const (
	BinarySIUnitB  BinarySIUnit = "B"
	BinarySIUnitKi BinarySIUnit = "Ki"
	BinarySIUnitMi BinarySIUnit = "Mi"
	BinarySIUnitGi BinarySIUnit = "Gi"
	BinarySIUnitTi BinarySIUnit = "Ti"
	BinarySIUnitPi BinarySIUnit = "Pi"
)

var binarySIUnitDivisors = map[BinarySIUnit]int64{
	BinarySIUnitB:  1,
	BinarySIUnitKi: 1024,
	BinarySIUnitMi: 1024 * 1024,
	BinarySIUnitGi: 1024 * 1024 * 1024,
	BinarySIUnitTi: 1024 * 1024 * 1024 * 1024,
	BinarySIUnitPi: 1024 * 1024 * 1024 * 1024 * 1024,
}

// GetMemoryAs returns the quantity as an integer in the specified binary SI unit.
func (q *Quantity) GetMemoryAs(unit BinarySIUnit, decimals int) float64 {
	divisor, ok := binarySIUnitDivisors[unit]
	if !ok {
		divisor = 1
	}
	return getRoundedValue(float64(q.Value())/float64(divisor), decimals)
}

func getRoundedValue(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// GetMemoryString returns the quantity formatted at the highest fitting binary SI unit.
func (q *Quantity) GetMemoryString() string {
	bytes := q.Value()
	switch {
	case bytes >= binarySIUnitDivisors[BinarySIUnitPi]:
		return fmt.Sprintf("%dPiB", bytes/binarySIUnitDivisors[BinarySIUnitPi])
	case bytes >= binarySIUnitDivisors[BinarySIUnitTi]:
		return fmt.Sprintf("%dTiB", bytes/binarySIUnitDivisors[BinarySIUnitTi])
	case bytes >= binarySIUnitDivisors[BinarySIUnitGi]:
		return fmt.Sprintf("%dGiB", bytes/binarySIUnitDivisors[BinarySIUnitGi])
	case bytes >= binarySIUnitDivisors[BinarySIUnitMi]:
		return fmt.Sprintf("%dMiB", bytes/binarySIUnitDivisors[BinarySIUnitMi])
	case bytes >= binarySIUnitDivisors[BinarySIUnitKi]:
		return fmt.Sprintf("%dKiB", bytes/binarySIUnitDivisors[BinarySIUnitKi])
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}
