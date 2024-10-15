// TODO: Remove before merge
package rortest

import (
	mongohelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DatabaseQuery struct {
	Sorts  []string            `json:"sorts,omitempty"`
	Query  *DatabaseQueryQuery `json:"query,omitempty"`
	Fields []string            `json:"fields,omitempty"`
}
type DatabaseQueryQuery struct {
	UID             string `json:"uid,omitempty"`
	Kind            string `json:"kind,omitempty"`
	metav1.TypeMeta `json:",inline"`
	OwnerRef        rortypes.RorResourceOwnerReference `json:"owner_ref,omitempty"`
	Name            string                             `json:"name,omitempty"`
}

func NewQuery() Queryinterface {
	return &DatabaseQuery{}
}

func (db *DatabaseQuery) SetSort(sort string) Queryinterface {
	db.Sorts = append(db.Sorts, sort)
	return db
}
func (db DatabaseQuery) GetSorts() []string {

	return db.Sorts
}
func (db *DatabaseQuery) SetQuery(query *DatabaseQueryQuery) Queryinterface {
	db.Query = query
	return db
}

func (db DatabaseQuery) GetQuery() *DatabaseQueryQuery {
	return db.Query
}

func (db *DatabaseQuery) SetField(sort string) Queryinterface {
	db.Fields = append(db.Fields, sort)
	return db
}

func (db DatabaseQuery) GetFields() []string {

	return db.Fields
}

type Queryinterface interface {
	SetSort(string) Queryinterface
	GetSorts() []string
	SetQuery(*DatabaseQueryQuery) Queryinterface
	GetQuery() *DatabaseQueryQuery
	SetField(string) Queryinterface
	GetFields() []string
}

type Databaseprovider interface {
	Query(Queryinterface) string
	Do(Queryinterface) []byte
}

type Mongoprovider struct {
	Client *mongo.Database
}

func NewMongodbProvider(client *mongo.Database) Databaseprovider {
	mongoret := Mongoprovider{Client: client}
	return mongoret
}
func (mdb Mongoprovider) Query(query Queryinterface) string {

	var mongoquery []bson.M
	for _, sort := range query.GetSorts() {
		mongoquery = append(mongoquery, bson.M{"sortby": sort})
	}
	mongohelper.PrettyprintBSON(mongoquery)
	return "OK"
}
func (mdb Mongoprovider) Do(query Queryinterface) []byte {
	return []byte("OK")
}
