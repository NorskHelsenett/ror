package initialservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/clients/rorapiclient"
	"github.com/NorskHelsenett/ror/internal/kubernetes/operator/initialize"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Deprecated: FetchConfiguration is deprecated
func FetchConfiguration() error {
	rorClient, err := rorapiclient.GetOrCreateRorRestyClient()
	if err != nil {
		rlog.Fatal("could not create ror api client", err)
	}

	requestUrl := "v1/configs/operator"
	response, err := rorClient.R().
		SetHeader("Content-Type", "application/json").
		Get(requestUrl)
	if err != nil {
		rlog.Error("Could not send data to ror-api", err)
		return errors.New("could not send data to ror-api")
	}

	if response.StatusCode() > 299 {
		rlog.Fatal("Error calling ror api", fmt.Errorf("non 200 errorcode"),
			rlog.String("request", fmt.Sprintf("%s/%s", viper.GetString(configconsts.API_ENDPOINT), requestUrl)),
			rlog.Int("code", response.StatusCode()))
	}

	return nil
}

// Deprecated: FetchApikey is deprecated
func FetchApikey(k8sClient *kubernetes.Clientset, metricsClient *metrics.Clientset) (string, error) {
	clusterInfo, err := initialize.GetClusterInfoFromNode(k8sClient, metricsClient)
	if err != nil {
		return "", errors.New("could not get identifier")
	}

	rorUrl := viper.GetString(configconsts.API_ENDPOINT)
	apikey, err := initialize.GetApikey(clusterInfo, rorUrl)
	if err != nil {
		rlog.Error("not able to get api key", err,
			rlog.String("identifier", clusterInfo.ClusterName),
			rlog.String("ror url", rorUrl))

		return "", fmt.Errorf("could not fetch api key from API (url: %s)", rorUrl)
	}

	viper.Set(configconsts.API_KEY, apikey)
	return apikey, nil
}

// Deprecated: ExtractApikeyOrDie is deprecated
func ExtractApikeyOrDie(k8sClient *kubernetes.Clientset, metricsClient *metrics.Clientset) error {
	secretName := viper.GetString(configconsts.API_KEY_SECRET)
	namespace := viper.GetString(configconsts.POD_NAMESPACE)
	secretApiKey := "APIKEY"
	secret, err := k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metaV1.GetOptions{})
	if err != nil {
		apikey, err := FetchApikey(k8sClient, metricsClient)
		if err != nil {
			rlog.Error("could not fetch api key: ", err)
			return errors.New("could not fetch api key")
		}
		secret, err = k8sClient.CoreV1().Secrets(namespace).Create(context.TODO(),
			&coreV1.Secret{
				ObjectMeta: metaV1.ObjectMeta{
					Name:      secretName,
					Namespace: namespace,
				},
				Type: "Opaque",
				StringData: map[string]string{
					secretApiKey: apikey,
				},
			},
			metaV1.CreateOptions{})
		if err != nil {
			rlog.Error("could not create k8s secret: ", err)
			return errors.New("could not create secret")
		}
	}

	apikey := string(secret.Data[secretApiKey])
	viper.Set(configconsts.API_KEY, apikey)

	return nil
}

// Deprecated: GetOrCreateNamespace is deprecated
func GetOrCreateNamespace(k8sClient *kubernetes.Clientset) error {
	namespace := viper.GetString(configconsts.POD_NAMESPACE)
	_, err := k8sClient.CoreV1().Namespaces().Get(context.TODO(), namespace, metaV1.GetOptions{})
	if err != nil {
		_, err = k8sClient.CoreV1().Namespaces().Create(context.TODO(),
			&coreV1.Namespace{
				ObjectMeta: metaV1.ObjectMeta{
					Name: namespace,
				},
			},
			metaV1.CreateOptions{})
		if err != nil {
			rlog.Error("could not create namespace: ", err)
			return errors.New("could not create namespace")
		}
	}

	return nil
}
