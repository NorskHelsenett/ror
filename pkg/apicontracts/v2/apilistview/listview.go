package apilistview

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ListType string

const (
	ListTypeGrid  ListType = "grid"
	ListTypeTable ListType = "table"
	ListTypeChart ListType = "chart"
)

type ListFieldType string

const (
	ListFieldTypeString   ListFieldType = "string"
	ListFieldTypeNumber   ListFieldType = "number"
	ListFieldTypeDate     ListFieldType = "date"
	ListFieldTypeDateTime ListFieldType = "datetime"
	ListFieldTypeBoolean  ListFieldType = "boolean"
	ListFieldTypeEnum     ListFieldType = "enum"
)

type ListView struct {
	Type    ListType    `json:"type,omitempty"`
	Columns []ListField `json:"columns,omitempty"`
	Rows    []ListData  `json:"rows,omitempty"`
}

type ListField struct {
	Name             string          `json:"name"`  // Clustername
	Order            int             `json:"order"` //1
	Default          bool            `json:"default"`
	Writeable        bool            `json:"writeable,omitempty"` //true
	Type             ListFieldType   `json:"type,omitempty"`
	PossibleValues   []string        `json:"possible_values,omitempty"`
	ResourceType     metav1.TypeMeta `json:"resource_type,omitempty"`
	ResourceFieldRef string          `json:"resource_field_ref,omitempty"` //"spec.clusterdata.clustername"
}

type ListData struct {
	Name        string `json:"name"`                  // Clustername
	Value       string `json:"value"`                 // "test-001"
	Resourceuid string `json:"resourceuid,omitempty"` // dfafasdf-sdafasf-asdfadsf-dasf-sfasdf
}
