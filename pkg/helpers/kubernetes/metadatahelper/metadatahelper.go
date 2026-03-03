package metadatahelper

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func GetAnnotationOrLabel(metadata metav1.ObjectMeta, key string) (string, bool) {
	value, ok := metadata.GetAnnotations()[key]
	if ok {
		return value, true
	}

	value, ok = metadata.GetLabels()[key]
	if ok {
		return value, true
	}
	return "", false
}

func CheckAnnotationOrLabel(metadata metav1.ObjectMeta, key string) bool {
	_, ok := GetAnnotationOrLabel(metadata, key)
	return ok
}
