package apiview

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ViewType string

const (
	ViewTypeGrid  ViewType = "grid"
	ViewTypeTable ViewType = "table"
	ViewTypeChart ViewType = "chart"
)

type ViewFieldType string

type ViewMetadata struct {
	Id          string
	Description string
	Name        string
	Version     string
}

const (
	ViewFieldTypeString   ViewFieldType = "string"
	ViewFieldTypeNumber   ViewFieldType = "number"
	ViewFieldTypeDate     ViewFieldType = "date"
	ViewFieldTypeDateTime ViewFieldType = "datetime"
	ViewFieldTypeBoolean  ViewFieldType = "boolean"
	ViewFieldTypeEnum     ViewFieldType = "enum"
)

type View struct {
	Type    ViewType    `json:"type,omitempty"`
	Columns []ViewField `json:"columns,omitempty"`
	Rows    []ViewRow   `json:"rows,omitempty"`
}

type ViewRow []ViewData

type ViewField struct {
	Name             string          `json:"name"`
	Description      string          `json:"description,omitempty"`
	Order            int             `json:"order"`
	Default          bool            `json:"default"`
	Writeable        bool            `json:"writeable,omitempty"`
	Type             ViewFieldType   `json:"type,omitempty"`
	PossibleValues   []string        `json:"possibleValues,omitempty"`
	ResourceType     metav1.TypeMeta `json:"resourceType,omitempty"`
	ResourceFieldRef string          `json:"resourceFieldRef,omitempty"` //"spec.clusterdata.clustername"
}

type ViewData struct {
	FieldName   string `json:"fieldName"`
	FieldValue  string `json:"fieldValue"`
	Description string `json:"description,omitempty"`
	ResourceUid string `json:"resourceUid,omitempty"`
}
