package secretservice

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func ExtractKubeconfigSecretFromSupervisorCluster(ctx context.Context, k8sConfig *rest.Config, namespace, secretName string) (*v1.Secret, error) {
	if k8sConfig == nil {
		rlog.Error("k8sConfig is nil", nil)
		return nil, nil
	}

	k8sClient, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		rlog.Error("error in config", err)
		return nil, err
	}

	secret, err := k8sClient.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		rlog.Error("error in getting secret", err)
		return nil, err
	}

	return secret, nil
}
