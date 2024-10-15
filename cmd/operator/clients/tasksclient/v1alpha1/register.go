package v1alpha1

import (
	"github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubectl/pkg/scheme"

	"k8s.io/client-go/rest"
)

var SchemeGroupVersion = v1alpha1.GroupVersion
var Scheme runtime.Scheme
var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&v1alpha1.Task{},
		&v1alpha1.TaskList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	Scheme = *scheme
	return nil
}

func NewClient(cfg *rest.Config) (*TasksConfigV1Alpha1Client, error) {
	// scheme := runtime.NewScheme()
	// SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	// if err := SchemeBuilder.AddToScheme(scheme); err != nil {
	// 	return nil, err
	// }
	config := *cfg

	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupVersion.Group, Version: v1alpha1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &TasksConfigV1Alpha1Client{restClient: client}, nil
}
