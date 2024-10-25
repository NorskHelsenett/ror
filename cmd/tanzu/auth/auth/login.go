package auth

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/google/uuid"

	"github.com/spf13/viper"

	"github.com/alessio/shellescape"
)

func Login(credentials apicontracts.TanzuKubeConfigPayload) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()

	configFolderPath := viper.GetString(configconsts.TANZU_AUTH_CONFIG_FOLDER_PATH)
	uniqueId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	kubeconfigPath := fmt.Sprintf("%s/kubeconfig-%s.config", configFolderPath, uniqueId.String())
	_, err = os.Create(kubeconfigPath) // #nosec G304 - kubeconfigPath is a constant string
	if err != nil {
		rlog.Errorc(ctx, "Failed to create kubeconfig file", err)
	} //

	quotedDatacenterurl := shellescape.Quote(credentials.DatacenterUrl)
	quotedWorkspaceName := shellescape.Quote(credentials.WorkspaceName)
	quotedUser := shellescape.Quote(credentials.User)
	quotedPassword := shellescape.Quote(credentials.Password)
	kubectlVspherePath := viper.GetString(configconsts.TANZU_AUTH_KUBE_VSPHERE_PATH)
	rlog.Info("Authentication", rlog.String("time", time.Now().String()))
	rlog.Info("Connection to url ", rlog.String("url", quotedDatacenterurl), rlog.String("username", quotedUser), rlog.String("kubeconfig", kubeconfigPath))
	args := []string{"-c"}
	env := make([]string, 0)

	var loginCmd string

	if credentials.WorkspaceOnly {
		loginCmd = fmt.Sprintf("%s login --server=%s -u %s --insecure-skip-tls-verify --tanzu-kubernetes-cluster-namespace %s --request-timeout=16s",
			kubectlVspherePath, quotedDatacenterurl, quotedUser, quotedWorkspaceName)
	} else {
		quotedClusterName := shellescape.Quote(credentials.ClusterName)
		loginCmd = fmt.Sprintf("%s login --server=%s -u %s --insecure-skip-tls-verify --tanzu-kubernetes-cluster-namespace %s --tanzu-kubernetes-cluster-name %s --request-timeout=16s",
			kubectlVspherePath, quotedDatacenterurl, quotedUser, quotedWorkspaceName, quotedClusterName)
	}

	if loginCmd == "" {
		return "", fmt.Errorf("failed to create login command")
	}

	args = append(args, loginCmd)
	_, _ = fmt.Println(loginCmd)

	env = append(env, "KUBECTL_VSPHERE_PASSWORD="+quotedPassword)
	env = append(env, "KUBECONFIG="+kubeconfigPath)

	directoryOfKubectl := filepath.Dir(kubectlVspherePath)
	development := viper.GetBool(configconsts.DEVELOPMENT)
	if !development {
		env = append(env, "PATH=/usr/bin:/app:/bin")
	} else {
		env = append(env, "PATH="+directoryOfKubectl)
	}

	command := exec.CommandContext(ctx, "/bin/sh", args...)
	command.Env = env
	command.Dir = directoryOfKubectl

	output, err := command.CombinedOutput()
	if err != nil {
		rlog.Infof("failed to run vsphere command: %s", kubectlVspherePath)
		rlog.Error("failed to run vsphere command", err,
			rlog.String("vsphere out", string(output)))

		println("Retry with more logging: ")
		loginCmd = loginCmd + " -v=10"
		args2 := []string{"-c"}
		args2 = append(args2, loginCmd)
		command2 := exec.CommandContext(ctx, "/bin/sh", args2...)
		command2.Env = env
		command2.Dir = directoryOfKubectl

		output2, err := command2.CombinedOutput()
		if err != nil {
			rlog.Infof("failed to run vsphere command with more logging: %s", kubectlVspherePath)
			rlog.Error("failed to run vsphere command with more logging", err,
				rlog.String("vsphere out", string(output2)))
			return "", err
		}

		return "", err
	}
	rlog.Infof("Vsphere command did return ok: %s", kubectlVspherePath)
	rlog.Debug("Vsphere command", rlog.String("vsphere output", string(output)))
	return kubeconfigPath, nil
}
