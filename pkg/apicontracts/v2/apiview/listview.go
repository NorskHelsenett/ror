package apiview

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ViewType string

const (
	ViewTypeList  ViewType = "list"
	ViewTypeChart ViewType = "chart"
)

type ViewFieldType string

type ViewMetadata struct {
	Id          string   `json:"id"`
	Type        ViewType `json:"type"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Version     int      `json:"version,omitempty"`
}

const (
	ViewFieldTypeString   ViewFieldType = "string"
	ViewFieldTypeNumber   ViewFieldType = "number"
	ViewFieldTypeDate     ViewFieldType = "date"
	ViewFieldTypeDateTime ViewFieldType = "datetime"
	ViewFieldTypeBoolean  ViewFieldType = "boolean"
	ViewFieldTypeEnum     ViewFieldType = "enum"
	ViewFieldTypeArray    ViewFieldType = "array"
	ViewFieldTypeObject   ViewFieldType = "object"
)

type View struct {
	Type    ViewType     `json:"type,omitempty"`
	Columns []ViewColumn `json:"columns,omitempty"`
	Rows    []ViewRow    `json:"rows,omitempty"`
}

type ViewRow map[string]ViewValue

type ViewColumn struct {
	Name             string          `json:"name"` // name of the column, used as a key in the rows, must be unique within a view
	Description      string          `json:"description,omitempty"`
	Order            int             `json:"order"` // order of the column in the view, it may be holes and duplicates in the order
	Default          bool            `json:"default"`
	Writeable        bool            `json:"writeable,omitempty"`
	Type             ViewFieldType   `json:"type,omitempty"`
	PossibleValues   []string        `json:"possibleValues,omitempty"`
	ResourceType     metav1.TypeMeta `json:"resourceType,omitzero"`      // resourceType in ror
	ResourceFieldRef string          `json:"resourceFieldRef,omitempty"` //"spec.clusterdata.clustername"
}

type ViewValue struct {
	FieldValue  any    `json:"fieldValue"`
	Description string `json:"description,omitempty"`
	ResourceUid string `json:"resourceUid,omitempty"`
}
