package roroperatorservice

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/settings"
	"github.com/NorskHelsenett/ror/internal/kubernetes/secretservice"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func InstallRorOperatorInCluster(ctx context.Context, supervisorNamespace, clusterName string) error {
	kubeconfigSecretName := fmt.Sprintf("%s-kubeconfig", clusterName)
	result, err := secretservice.ExtractKubeconfigSecretFromSupervisorCluster(context.Background(), settings.K8sConfig, supervisorNamespace, kubeconfigSecretName)
	if err != nil {
		rlog.Errorc(ctx, "could not extract kubeconfig secret from supervisor cluster", err)
		return err
	}

	rlog.Debug("kubeconfig", rlog.Any("kubeconfig", result))

	kubeconfigBytes := result.Data["value"]
	rlog.Debug("kubeconfig", rlog.Any("kubeconfig", kubeconfigBytes))

	kubeconfig := string(kubeconfigBytes)
	rlog.Debug("kubeconfig", rlog.String("kubeconfig", kubeconfig))

	k8sconfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		rlog.Errorc(ctx, "could not create k8s config", err)
		return err
	}

	rlog.Debug("k8sconfig", rlog.Any("k8sconfig", k8sconfig))

	k8sClient, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		rlog.Errorc(ctx, "could not create k8s client", err)
		return err
	}

	namespace := viper.GetString(configconsts.ROR_OPERATOR_NAMESPACE)
	err = checkOrCreateNamespace(ctx, namespace, k8sClient)
	if err != nil {
		rlog.Errorc(ctx, "could not create namespace", err, rlog.String("namespace", namespace))
		return err
	}

	serviceAccountName := "ror-operator-installer-sa"
	err = checkOrCreateServiceAccount(ctx, namespace, serviceAccountName, k8sClient)
	if err != nil {
		rlog.Errorc(ctx, "could not create serviceaccount", err, rlog.String("serviceaccount", serviceAccountName))
		return err
	}

	rolebindingName := "ror-operator-installer-rolebinding"
	err = checkOrCreateRolebinding(ctx, namespace, rolebindingName, serviceAccountName, k8sClient)
	if err != nil {
		rlog.Errorc(ctx, "could not create rolebinding", err, rlog.String("rolebinding", rolebindingName))
		return err
	}

	clusterRoleBindingName := "ror-operator-installer-clusterrolebinding"
	err = checkOrCreateClusterRoleBinding(ctx, namespace, clusterRoleBindingName, serviceAccountName, k8sClient)
	if err != nil {
		rlog.Errorc(ctx, "could not create clusterrolebinding", err, rlog.String("clusterrolebinding", clusterRoleBindingName))
		return err
	}

	jobDefinition := createJobDefinition(ctx, namespace, serviceAccountName, k8sClient)
	k8sjob, err := k8sClient.BatchV1().Jobs(namespace).Get(context.Background(), jobDefinition.Name, metav1.GetOptions{})
	if err != nil {
		k8sjob, err = k8sClient.BatchV1().Jobs(namespace).Create(context.Background(), jobDefinition, metav1.CreateOptions{})
		if err != nil {
			rlog.Errorc(ctx, "could not create job", err, rlog.Any("job", jobDefinition))
			return err
		}
	}
	rlog.Debug("k8sjob", rlog.Any("k8sjob", k8sjob))

	return nil
}

func checkOrCreateNamespace(ctx context.Context, namespace string, k8sClient *kubernetes.Clientset) error {
	k8sNamespace, err := k8sClient.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		rlog.Debug("could not get namespace, creating it")
		k8sNamespace, err = k8sClient.CoreV1().Namespaces().Create(context.Background(), &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}, metav1.CreateOptions{})
		if err != nil {
			rlog.Errorc(ctx, "could not create namespace", err, rlog.String("namespace", namespace))
			return err
		}

	}
	rlog.Debugc(ctx, "k8snamespace", rlog.Any("k8snamespace", k8sNamespace))
	return nil
}

func checkOrCreateServiceAccount(ctx context.Context, namespace, saName string, k8sClient *kubernetes.Clientset) error {
	sa, err := k8sClient.CoreV1().ServiceAccounts(namespace).Get(context.Background(), saName, metav1.GetOptions{})
	if err != nil {
		sa, err = k8sClient.CoreV1().ServiceAccounts(namespace).Create(context.Background(), &v1.ServiceAccount{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServiceAccount",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      saName,
				Namespace: namespace,
			},
		}, metav1.CreateOptions{})
		if err != nil {
			rlog.Errorc(ctx, "could not create serviceaccount", err, rlog.String("serviceaccount", saName))
			return err
		}
	}
	rlog.Debugc(ctx, "sa", rlog.Any("sa", sa))
	return nil
}

func checkOrCreateRolebinding(ctx context.Context, namespace string, rbName string, serviceAccountName string, k8sClient *kubernetes.Clientset) error {
	rolebinding, err := k8sClient.RbacV1().RoleBindings(namespace).Get(context.Background(), rbName, metav1.GetOptions{})
	if err != nil {
		rolebinding, err = k8sClient.RbacV1().RoleBindings(namespace).Create(context.Background(), &rbacV1.RoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "RoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      rbName,
				Namespace: namespace,
			},
			RoleRef: rbacV1.RoleRef{
				Kind:     "ClusterRole",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     "psp:vmware-system-restricted",
			},
			Subjects: []rbacV1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      serviceAccountName,
					Namespace: namespace,
				},
			},
		}, metav1.CreateOptions{})
		if err != nil {
			rlog.Errorc(ctx, "could not create rolebinding", err, rlog.String("rolebinding", rbName))
			return err
		}
	}
	rlog.Debug("rolebinding", rlog.Any("rolebinding", rolebinding))
	return nil
}

func checkOrCreateClusterRoleBinding(ctx context.Context, namespace string, clusterRoleName string, serviceAccountName string, k8sClient *kubernetes.Clientset) error {
	clusterRoleBinding, err := k8sClient.RbacV1().ClusterRoleBindings().Get(context.Background(), clusterRoleName, metav1.GetOptions{})
	if err != nil {
		clusterRoleBinding, err = k8sClient.RbacV1().ClusterRoleBindings().Create(context.Background(), &rbacV1.ClusterRoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: clusterRoleName,
			},
			RoleRef: rbacV1.RoleRef{
				Kind:     "ClusterRole",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     "cluster-admin",
			},
			Subjects: []rbacV1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      serviceAccountName,
					Namespace: namespace,
				},
			},
		}, metav1.CreateOptions{})
		if err != nil {
			rlog.Errorc(ctx, "could not create clusterrolebinding", err, rlog.String("clusterrolebinding", clusterRoleName))
			return err
		}
	}
	rlog.Debug("clusterRoleBinding", rlog.Any("clusterRoleBinding", clusterRoleBinding))

	return nil
}

func createJobDefinition(ctx context.Context, namespace string, serviceAccountName string, k8sClient *kubernetes.Clientset) *batchv1.Job {
	var backoffLimit int32 = 3
	var activeDeadlineSeconds int64 = 180
	cmd := "helm upgrade --install ror-operator oci://ncr.sky.nhn.no/nhn-helm/ror-operator -n nhn-ror --create-namespace"
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "install-ror-operator",
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			BackoffLimit:          &backoffLimit,
			ActiveDeadlineSeconds: &activeDeadlineSeconds,
			Template: v1.PodTemplateSpec{

				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "install-ror-operator",
							Image:           "ncr.sky.nhn.no/ror/tasks/devops-base:1.0.0",
							ImagePullPolicy: v1.PullAlways,
							Command:         []string{"sh", "-c", cmd},
						},
					},
					RestartPolicy:      v1.RestartPolicyNever,
					ServiceAccountName: serviceAccountName,
				},
			},
		},
	}
	return job
}
