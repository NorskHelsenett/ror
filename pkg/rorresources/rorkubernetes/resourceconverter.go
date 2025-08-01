package rorkubernetes

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// convertUnstructuredToStruct converts a map[string]interface{} from unstructured
// to a target struct using reflection. This avoids JSON marshal/unmarshal operations
// for better memory efficiency.
func convertUnstructuredToStruct(src map[string]any, dst any) error {
	if src == nil {
		return nil
	}

	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		rlog.Error("Destination must be a pointer to struct", nil)
		return nil
	}

	dstValue = dstValue.Elem()
	dstType := dstValue.Type()

	return convertMapToStruct(src, dstValue, dstType)
}

// convertMapToStruct recursively converts a map to a struct
func convertMapToStruct(src map[string]interface{}, dstValue reflect.Value, dstType reflect.Type) error {
	for i := 0; i < dstType.NumField(); i++ {
		field := dstType.Field(i)
		fieldValue := dstValue.Field(i)

		// Skip unexported fields
		if !fieldValue.CanSet() {
			continue
		}

		// Get the JSON tag name or use the field name
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		if jsonTag != "" && jsonTag != "-" {
			// Parse JSON tag (e.g., "name,omitempty" -> "name")
			if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
				fieldName = jsonTag[:commaIdx]
			} else {
				fieldName = jsonTag
			}
		}

		// Convert field name to lowercase for unstructured object lookup
		srcValue, exists := src[fieldName]
		if !exists {
			// Try with lowercase first letter
			lowercaseFieldName := strings.ToLower(fieldName[:1]) + fieldName[1:]
			srcValue, exists = src[lowercaseFieldName]
		}
		if !exists {
			continue
		}

		if err := setFieldValue(fieldValue, srcValue); err != nil {
			rlog.Error("Failed to set field "+field.Name, err)
			continue
		}
	}

	return nil
}

// setFieldValue sets a field value from an interface{} value using reflection
func setFieldValue(fieldValue reflect.Value, srcValue interface{}) error {
	if srcValue == nil {
		return nil
	}

	srcVal := reflect.ValueOf(srcValue)
	fieldType := fieldValue.Type()

	// Handle different field types
	switch fieldType.Kind() {
	case reflect.String:
		if srcVal.Kind() == reflect.String {
			fieldValue.SetString(srcVal.String())
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch srcVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetInt(srcVal.Int())
		case reflect.Float32, reflect.Float64:
			fieldValue.SetInt(int64(srcVal.Float()))
		}

	case reflect.Float32, reflect.Float64:
		switch srcVal.Kind() {
		case reflect.Float32, reflect.Float64:
			fieldValue.SetFloat(srcVal.Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetFloat(float64(srcVal.Int()))
		}

	case reflect.Bool:
		if srcVal.Kind() == reflect.Bool {
			fieldValue.SetBool(srcVal.Bool())
		}

	case reflect.Slice:
		if srcVal.Kind() == reflect.Slice {
			srcSlice := srcVal
			newSlice := reflect.MakeSlice(fieldType, srcSlice.Len(), srcSlice.Len())
			for i := 0; i < srcSlice.Len(); i++ {
				elemValue := newSlice.Index(i)
				if err := setFieldValue(elemValue, srcSlice.Index(i).Interface()); err != nil {
					return err
				}
			}
			fieldValue.Set(newSlice)
		}

	case reflect.Map:
		if srcVal.Kind() == reflect.Map {
			if fieldType.Key().Kind() == reflect.String {
				newMap := reflect.MakeMap(fieldType)
				for _, key := range srcVal.MapKeys() {
					if key.Kind() == reflect.String {
						mapValue := reflect.New(fieldType.Elem()).Elem()
						if err := setFieldValue(mapValue, srcVal.MapIndex(key).Interface()); err != nil {
							return err
						}
						newMap.SetMapIndex(key, mapValue)
					}
				}
				fieldValue.Set(newMap)
			}
		}

	case reflect.Struct:
		if srcVal.Kind() == reflect.Map {
			// Convert map to struct recursively
			if srcMap, ok := srcValue.(map[string]interface{}); ok {
				return convertMapToStruct(srcMap, fieldValue, fieldType)
			}
		}
		if fieldType.Name() == "Time" && srcVal.Kind() == reflect.String {
			// Handle time conversion if the field is of type Time
			timeValue, err := time.Parse(time.RFC3339, srcVal.String())
			if err != nil {
				return fmt.Errorf("failed to parse time: %w", err)
			}
			fieldValue.Set(reflect.ValueOf(v1.Time{Time: timeValue}))
		}

	// Handle pointer types

	case reflect.Ptr:
		if fieldType.Elem().Kind() == reflect.Struct {
			// Handle pointer to struct
			newStruct := reflect.New(fieldType.Elem())
			if srcMap, ok := srcValue.(map[string]interface{}); ok {
				if err := convertMapToStruct(srcMap, newStruct.Elem(), fieldType.Elem()); err != nil {
					return err
				}
				fieldValue.Set(newStruct)
			}
		} else {
			// Handle pointer to primitive type
			newValue := reflect.New(fieldType.Elem())
			if err := setFieldValue(newValue.Elem(), srcValue); err != nil {
				return err
			}
			fieldValue.Set(newValue)
		}

	default:
		// For other types, try direct assignment if types are compatible
		if srcVal.Type().ConvertibleTo(fieldType) {
			fieldValue.Set(srcVal.Convert(fieldType))
		}
	}

	return nil
}
