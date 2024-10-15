package auth

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/settings"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func Login() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()

	kubeconfigPath := viper.GetString(configconsts.TANZU_AGENT_KUBECONFIG)
	if viper.GetBool(configconsts.TANZU_AGENT_DELETE_KUBECONFIG) {
		deleteKubeconfigFile(kubeconfigPath)
	}

	datacenterUrl := viper.GetString(configconsts.TANZU_AGENT_DATACENTER_URL)
	rlog.Info("Authentication", rlog.String("time", time.Now().String()))
	rlog.Info("Connection to url ", rlog.String("url", datacenterUrl), rlog.String("username", viper.GetString(configconsts.TANZU_AGENT_USERNAME)), rlog.String("kubeconfig", kubeconfigPath))
	args := []string{
		"-c",
		fmt.Sprintf("%s login --server=%s -u %s --insecure-skip-tls-verify", viper.GetString(configconsts.TANZU_AGENT_KUBE_VSPHERE_PATH), datacenterUrl, viper.GetString(configconsts.TANZU_AGENT_USERNAME)),
	}
	env := make([]string, 0)
	fmt.Println("Command:", fmt.Sprintf("%s login --server=%s -u %s --insecure-skip-tls-verify", viper.GetString(configconsts.TANZU_AGENT_KUBE_VSPHERE_PATH), datacenterUrl, viper.GetString(configconsts.TANZU_AGENT_USERNAME)))

	env = append(env, "KUBECTL_VSPHERE_PASSWORD="+viper.GetString(configconsts.TANZU_AGENT_PASSWORD))
	env = append(env, "KUBECONFIG="+kubeconfigPath)

	directoryOfKubectl := filepath.Dir(viper.GetString(configconsts.TANZU_AGENT_KUBE_VSPHERE_PATH))
	development := viper.GetBool(configconsts.DEVELOPMENT)
	if !development {
		env = append(env, "PATH=/usr/bin:/app:/bin")
	} else {
		env = append(env, "PATH="+directoryOfKubectl)
	}

	command := exec.CommandContext(ctx, "/bin/sh", args...) // #nosec G204 - we are not using user input
	command.Env = env
	command.Dir = directoryOfKubectl

	output, err := command.CombinedOutput()
	if err != nil {
		rlog.Infof("failed to run vsphere command: %s", viper.GetString(configconsts.TANZU_AGENT_KUBE_VSPHERE_PATH))
		rlog.Error("failed to run vsphere command", err,
			rlog.String("vsphere out", string(output)))
		return err
	}
	rlog.Infof("Vsphere command did return ok: %s", viper.GetString(configconsts.TANZU_AGENT_KUBE_VSPHERE_PATH))
	rlog.Debug("Vsphere command", rlog.String("vsphere output", string(output)))
	return nil
}

func TokenHasExpired(k8sconfig *rest.Config) (bool, error) {
	if k8sconfig == nil {
		return true, nil
	}

	claims, err := getJWTokenClaims(k8sconfig.BearerToken)
	if err != nil {
		return true, err
	}

	expiry := int64(claims["exp"].(float64))
	now := time.Now().Unix()

	if now > expiry {
		return true, nil
	}

	return false, nil
}

func GetK8sConfig(access_to_tanzu bool) (*rest.Config, error) {
	k8sconfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	if !access_to_tanzu {
		return k8sconfig, nil
	}

	claims, err := getJWTokenClaims(k8sconfig.BearerToken)
	if err != nil {
		return nil, err
	}

	expiry := int64(claims["exp"].(float64))
	viper.Set(configconsts.TANZU_AGENT_TOKEN_EXPIRY, expiry)
	expiryDate := time.Unix(expiry, 0)
	rlog.Debug("Access token expires ", rlog.String("expire date", expiryDate.String()))

	settings.K8sConfig = k8sconfig
	return k8sconfig, nil
}

func getJWTokenClaims(rawToken string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(rawToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func deleteKubeconfigFile(path string) {
	_, err := os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			rlog.Fatal("Could not delete kubeconfig file ... ", err)
		}
	}
}
