package taskservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	"github.com/NorskHelsenett/ror/cmd/operator/clients"
	"github.com/NorskHelsenett/ror/cmd/operator/clients/ror"
	v1alpha1taskclient "github.com/NorskHelsenett/ror/cmd/operator/clients/tasksclient/v1alpha1"
	"github.com/NorskHelsenett/ror/cmd/operator/models"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

func CreateTasks(k8sclient *kubernetes.Clientset, operatorTask apicontracts.OperatorTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	taskClient, err := v1alpha1taskclient.NewClient(clients.Kubernetes.GetConfig())
	if err != nil {
		return errors.New("could not create a taskClient")
	}

	operatorJob, err := ror.FetchTaskConfiguration(context.TODO(), operatorTask)
	if err != nil {
		rlog.Error("could not fetch ror-task configuration: ", err)
		return errors.New("could not fetch ror-task configuration")
	}

	if operatorJob == nil {
		return fmt.Errorf("could not find task config for %s", operatorTask.Name)
	}

	operatorJob.RunOnce = operatorTask.RunOnce
	task, taskStatus, err := GetTaskAndTaskStatus(ctx, taskClient, operatorTask, operatorJob)
	if err != nil {
		return err
	}

	if taskStatus.Installed && taskStatus.ConfigChanged && operatorTask.RunOnce {
		rlog.Debug("task already installed and run once is set to true", rlog.String("task name", task.Name))
		return nil
	}

	if taskStatus.Installed && !taskStatus.ConfigChanged {
		rlog.Debug("task already installed, no changes to config", rlog.String("task name", task.Name))
		return nil
	}

	taskNamespace := viper.GetString(configconsts.POD_NAMESPACE)

	if task == nil {
		taskDefinition := v1alpha1.Task{
			TypeMeta: metav1.TypeMeta{
				APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
				Kind:       "Task",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      operatorTask.Name,
				Namespace: taskNamespace,
				Labels: map[string]string{
					"app.kubernetes.io/name":       "task",
					"app.kubernetes.io/instance":   "ror_operator",
					"app.kubernetes.io/part-of":    "task",
					"app.kubernetes.io/managed-by": "ror_operator",
					"app.kubernetes.io/created-by": "ror_operator",
				},
			},
			Spec: *operatorJob,
			Status: v1alpha1.TaskStatus{
				Phase:   v1alpha1.PhasePending,
				Success: &v1alpha1.TaskSuccess{},
				Failure: &v1alpha1.TaskFailure{},
			},
		}

		_, err = taskClient.TaskConfigs(ctx, taskNamespace).Create(&taskDefinition)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("could not create ROR Task %s: ", operatorTask.Name), err)
		}
	}

	return nil
}

func GetTaskAndTaskStatus(ctx context.Context, taskClient *v1alpha1taskclient.TasksConfigV1Alpha1Client, operatorTask apicontracts.OperatorTask, operatorJob *apicontracts.OperatorJob) (*v1alpha1.Task, models.TaskStatus, error) {
	operatorNamespace := viper.GetString(configconsts.POD_NAMESPACE)
	task, err := taskClient.TaskConfigs(ctx, operatorNamespace).Get(operatorTask.Name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return nil, models.TaskStatus{}, nil
		}
		rlog.Error("could not get task config", err)
		return nil, models.TaskStatus{}, err
	}

	if task.Name == "" || task.Status.Success == nil {
		return task, models.TaskStatus{
			ConfigChanged: false,
			Installed:     true,
		}, nil
	}

	configBytes, err := json.Marshal(operatorJob)
	if err != nil {
		return task, models.TaskStatus{}, err
	}

	if task.Name == operatorTask.Version && operatorTask.Version == operatorJob.ImageTag && task.Status.Success.ConfigMd5 == stringhelper.GetSHA256Hash(configBytes) {
		return task, models.TaskStatus{
			ConfigChanged: false,
			Installed:     true,
		}, nil
	}

	if task.Name == operatorTask.Name && operatorTask.Version == task.Spec.ImageTag && task.Status.Success.ConfigMd5 != stringhelper.GetSHA256Hash(configBytes) {
		return task, models.TaskStatus{
			ConfigChanged: true,
			Installed:     true,
		}, nil
	}

	return task, models.TaskStatus{ConfigChanged: false, Installed: false}, nil
}
