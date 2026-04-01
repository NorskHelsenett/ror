package rortypes

import (
	"fmt"

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
