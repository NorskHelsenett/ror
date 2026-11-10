package rorkubernetes

import (
	"reflect"
	"testing"

	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestSetFieldValueStringMapKeys(t *testing.T) {
	tests := []struct {
		name string
		dst  any
		src  any
		want any
	}{
		{
			name: "plain strings",
			dst:  new(map[string]string),
			src:  map[string]any{"example.com/key": "value"},
			want: map[string]string{"example.com/key": "value"},
		},
		{
			name: "annotation keys and values",
			dst:  new(map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue),
			src:  map[string]any{"example.com/key": "value"},
			want: map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue{"example.com/key": "value"},
		},
		{
			name: "label keys and values",
			dst:  new(map[gatewayv1.LabelKey]gatewayv1.LabelValue),
			src:  map[string]any{"example.com/key": "value"},
			want: map[gatewayv1.LabelKey]gatewayv1.LabelValue{"example.com/key": "value"},
		},
		{
			name: "typed source keys",
			dst:  new(map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue),
			src:  map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue{"example.com/key": "value"},
			want: map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue{"example.com/key": "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := reflect.ValueOf(tt.dst).Elem()
			if err := setFieldValue(dst, tt.src); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(dst.Interface(), tt.want) {
				t.Fatalf("got %#v, want %#v", dst.Interface(), tt.want)
			}
		})
	}
}

func TestGatewayInfrastructureMapConversion(t *testing.T) {
	resource := newGatewayFromMapInterface(map[string]any{
		"spec": map[string]any{
			"gatewayClassName": "example",
			"infrastructure": map[string]any{
				"annotations": map[string]any{"example.com/annotation": "annotation-value"},
				"labels":      map[string]any{"example.com/label": "label-value"},
			},
		},
	})
	if resource == nil || resource.Spec.Infrastructure == nil {
		t.Fatal("gateway infrastructure was not converted")
	}
	want := &gatewayv1.GatewayInfrastructure{
		Annotations: map[gatewayv1.AnnotationKey]gatewayv1.AnnotationValue{
			"example.com/annotation": "annotation-value",
		},
		Labels: map[gatewayv1.LabelKey]gatewayv1.LabelValue{
			"example.com/label": "label-value",
		},
	}
	if !reflect.DeepEqual(resource.Spec.Infrastructure, want) {
		t.Fatalf("got %#v, want %#v", resource.Spec.Infrastructure, want)
	}
}

func TestHTTPRouteMapConversionIncludesEmbeddedFields(t *testing.T) {
	resource := newHTTPRouteFromMapInterface(map[string]any{
		"spec": map[string]any{
			"parentRefs": []any{
				map[string]any{"name": "gateway", "sectionName": "https"},
			},
			"rules": []any{
				map[string]any{
					"backendRefs": []any{
						map[string]any{"name": "backend", "port": 8080, "weight": 7},
					},
				},
			},
		},
		"status": map[string]any{
			"parents": []any{
				map[string]any{
					"parentRef":      map[string]any{"name": "gateway"},
					"controllerName": "example.com/controller",
				},
			},
		},
	})
	if resource == nil {
		t.Fatal("HTTPRoute was not converted")
	}
	if len(resource.Spec.ParentRefs) != 1 || string(resource.Spec.ParentRefs[0].Name) != "gateway" ||
		resource.Spec.ParentRefs[0].SectionName == nil || string(*resource.Spec.ParentRefs[0].SectionName) != "https" {
		t.Fatalf("HTTPRoute parentRefs were not converted: %#v", resource.Spec.ParentRefs)
	}
	if len(resource.Spec.Rules) != 1 || len(resource.Spec.Rules[0].BackendRefs) != 1 {
		t.Fatalf("HTTPRoute backendRefs were not converted: %#v", resource.Spec.Rules)
	}
	backend := resource.Spec.Rules[0].BackendRefs[0]
	if string(backend.Name) != "backend" || backend.Port == nil || int32(*backend.Port) != 8080 ||
		backend.Weight == nil || *backend.Weight != 7 {
		t.Fatalf("HTTPRoute backend ref fields were not converted: %#v", backend)
	}
	if len(resource.Status.Parents) != 1 || string(resource.Status.Parents[0].ParentRef.Name) != "gateway" ||
		string(resource.Status.Parents[0].ControllerName) != "example.com/controller" {
		t.Fatalf("HTTPRoute status parents were not converted: %#v", resource.Status.Parents)
	}
}
