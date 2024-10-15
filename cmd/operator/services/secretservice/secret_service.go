package secretservice

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/operator/models"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	AppLogPath  = "AppLog"
	FailurePath = "Failures"
)

func GetSecretContent[T any](k8sClient *kubernetes.Clientset, contentPath string) (*T, error) {
	secret, err := GetAppSecret(k8sClient)
	if err != nil {
		return nil, errors.New("could not extract applog secret")
	}

	content := secret.Data[contentPath]
	var installedState T
	err = json.Unmarshal(content, &installedState)
	if err != nil {
		return nil, errors.New("could not unmarshal secret data")
	}

	return &installedState, nil
}

func SetSecretContent[T any](k8sClient *kubernetes.Clientset, contentPath string, content *T) (bool, error) {
	secret, err := GetAppSecret(k8sClient)
	if err != nil {
		return false, errors.New("could not extract applog secret")
	}

	contentByteArray, err := json.Marshal(content)
	if err != nil {
		return false, errors.New("could not marshal content")
	}

	secret.Data[contentPath] = contentByteArray

	_, err = k8sClient.CoreV1().Secrets(viper.GetString(configconsts.POD_NAMESPACE)).Update(context.TODO(), secret, metaV1.UpdateOptions{})
	if err != nil {
		return false, errors.New("could not update secret")
	}

	return true, nil
}

func CreateAppSecret(k8sClient *kubernetes.Clientset, secretName string, secretNamespace string) (bool, error) {
	if k8sClient == nil {
		return false, errors.New("kubernetes client is nil")
	}

	if secretName == "" {
		return false, errors.New("secret name is empty")
	}

	if secretNamespace == "" {
		return false, errors.New("secret namespace is empty")
	}

	installedState, err := json.Marshal(models.InstalledState{
		AppLog: make([]models.App, 0),
	})
	if err != nil {
		return false, errors.New("unable to marshal data")
	}

	failures, err := json.Marshal(models.ApplicationFailLog{
		Failures: make([]models.Failure, 0),
	})
	if err != nil {
		return false, errors.New("unable to marshal data")
	}

	existingNamespace, _ := k8sClient.CoreV1().Namespaces().Get(context.TODO(), secretNamespace, metaV1.GetOptions{})
	if existingNamespace == nil || len(existingNamespace.Name) < 1 {
		_, err := k8sClient.CoreV1().Namespaces().Create(context.TODO(), &coreV1.Namespace{
			ObjectMeta: metaV1.ObjectMeta{
				Name: secretNamespace,
			},
		}, metaV1.CreateOptions{})
		if err != nil {
			rlog.Error("Could not create applog namespace", err)
			return false, errors.New("could not create applog namespace ")
		}
	}

	_, err = k8sClient.CoreV1().Secrets(secretNamespace).Create(context.TODO(),
		&coreV1.Secret{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      secretName,
				Namespace: secretNamespace,
			},
			Type: "Opaque",
			Data: map[string][]byte{
				AppLogPath:  installedState,
				FailurePath: failures,
			},
		},
		metaV1.CreateOptions{})
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAppSecret(k8sclient *kubernetes.Clientset) (*coreV1.Secret, error) {
	namespace := viper.GetString(configconsts.POD_NAMESPACE)
	name := viper.GetString(configconsts.OPERATOR_APPLOG_SECRET_NAME)
	secret, err := k8sclient.CoreV1().Secrets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		rlog.Error("Could not get applog secret", err)
		return nil, errors.New("could not get secret")
	}

	return secret, nil
}
