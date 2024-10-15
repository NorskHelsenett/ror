package jobservice

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/spf13/viper"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateK8sJob(ctx context.Context, job *v1alpha1.Task) (*batchv1.Job, error) {
	var backOffLimit int32 = viper.GetInt32(configconsts.OPERATOR_BACKOFF_LIMIT)
	var activeDeadlineSeconds int64 = viper.GetInt64(configconsts.OPERATOR_DEADLINE_SECONDS)

	jobNamespace := viper.GetString(configconsts.POD_NAMESPACE)

	containerRegPrefix := viper.GetString(configconsts.CONTAINER_REG_PREFIX)
	versionSeparator := ":"
	if strings.Contains(job.Spec.ImageName, "sha") {
		versionSeparator = "@"
	}

	jobImage := fmt.Sprintf("%s%s%s%s", containerRegPrefix, job.Spec.ImageName, versionSeparator, job.Spec.ImageTag)
	if job.Spec.BackOffLimit != backOffLimit {
		backOffLimit = job.Spec.BackOffLimit
	}

	k8sJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      job.Name,
			Namespace: jobNamespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            job.Name,
							Image:           jobImage,
							ImagePullPolicy: v1.PullAlways,
							Command:         []string{job.Spec.Cmd},
						},
					},
					ServiceAccountName: viper.GetString(configconsts.OPERATOR_JOB_SERVICE_ACCOUNT),
					RestartPolicy:      v1.RestartPolicyNever,
				},
			},
			BackoffLimit:          &backOffLimit,
			ActiveDeadlineSeconds: &activeDeadlineSeconds,
		},
	}

	for _, config := range job.Spec.Configs {
		if config.Type == apicontracts.OperatorJobConfigTypeEnv {
			addEnvSecretReference(k8sJob, job.Name, config)
		} else if config.Type == apicontracts.OperatorJobConfigTypeFile {
			addSecretToJob(k8sJob, job.Name, config)
		}
	}

	return k8sJob, nil
}

func addEnvSecretReference(k8sJob *batchv1.Job, jobName string, config apicontracts.OperatorJobConfig) {
	secretName := fmt.Sprintf("%s-%s-secret", jobName, config.Name)
	k8sJob.Spec.Template.Spec.Containers[0].EnvFrom = append(k8sJob.Spec.Template.Spec.Containers[0].EnvFrom, v1.EnvFromSource{
		SecretRef: &v1.SecretEnvSource{
			LocalObjectReference: v1.LocalObjectReference{
				Name: secretName,
			},
		},
	})
}

func addSecretToJob(k8sJob *batchv1.Job, jobName string, config apicontracts.OperatorJobConfig) {
	secretName := fmt.Sprintf("%s-%s-secret", jobName, config.Name)
	nameVolume := fmt.Sprintf("%s-volume", secretName)
	defaultMode := int32(0755)

	k8sJob.Spec.Template.Spec.Volumes = append(k8sJob.Spec.Template.Spec.Volumes, v1.Volume{
		Name: nameVolume,
		VolumeSource: v1.VolumeSource{
			Secret: &v1.SecretVolumeSource{
				SecretName:  secretName,
				DefaultMode: &defaultMode,
			},
		},
	})

	k8sJob.Spec.Template.Spec.Containers[0].VolumeMounts = append(k8sJob.Spec.Template.Spec.Containers[0].VolumeMounts, v1.VolumeMount{
		Name:      nameVolume,
		MountPath: config.Path,
	})
}
