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
