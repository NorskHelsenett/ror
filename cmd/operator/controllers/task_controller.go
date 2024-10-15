/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	appsv1alpha1 "github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	"github.com/NorskHelsenett/ror/cmd/operator/clients"
	"github.com/NorskHelsenett/ror/cmd/operator/services/jobservice"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	valuesName = "values"
	scriptName = "scripts"
)

// TaskReconciler reconciles a Task object
type TaskReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ror.nhn.no,resources=tasks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ror.nhn.no,resources=tasks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ror.nhn.no,resources=tasks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Task object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *TaskReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	k8sLogger := log.FromContext(ctx)

	k8sLogger.Info("⚡️ Event received! ⚡️")
	k8sLogger.Info("Request: ", "req", req)

	task := &appsv1alpha1.Task{}
	err := r.Get(ctx, req.NamespacedName, task)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	// if no phase set, default to Pending
	if task.Status.Phase == "" {
		task.Status.Phase = appsv1alpha1.PhasePending
	}

	switch task.Status.Phase {
	case appsv1alpha1.PhasePending:
		k8sLogger.Info("Phase: Pending 🕹️")

		err = runK8sJob(ctx, r, req, task)
		if err != nil {
			_ = setErrorState(ctx, r, task, err)
			return ctrl.Result{}, err
		}

		task.Status.Phase = appsv1alpha1.PhaseRunning
	case appsv1alpha1.PhaseRunning:
		k8sLogger.Info("Phase: Running 🚙")
		err := listenForChangesInJobAndReport(ctx, r, req, task)
		if err != nil {
			_ = setErrorState(ctx, r, task, err)
			return ctrl.Result{}, err
		}

		task.Status.Phase = appsv1alpha1.PhaseDone

	case appsv1alpha1.PhaseDone:
		k8sLogger.Info("Phase: DONE ✅")
		// reconcile without requeue
		return ctrl.Result{}, nil
	default:
		k8sLogger.Info("Unknown phase ‼️")
		return ctrl.Result{}, nil
	}

	// update status
	err = r.Status().Update(ctx, task)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TaskReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Task{}).
		Owns(&batchv1.Job{}).
		Owns(&v1.Secret{}).
		Complete(r)
}

func setErrorState(ctx context.Context, r *TaskReconciler, task *appsv1alpha1.Task, err error) error {
	rlog.Errorc(ctx, "error occurred", err)
	task.Status.Phase = appsv1alpha1.PhaseError

	if task.Status.Failure == nil {
		task.Status.Failure = &appsv1alpha1.TaskFailure{
			Reason: err.Error(),
		}
	} else {
		task.Status.Failure.Reason = err.Error()
	}

	err = r.Status().Update(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func runK8sJob(ctx context.Context, r *TaskReconciler, req ctrl.Request, task *v1alpha1.Task) error {
	k8sLogger := log.FromContext(ctx)

	var existingJob batchv1.Job
	err := r.Get(ctx, req.NamespacedName, &existingJob)
	if !errors.IsNotFound(err) {
		k8sLogger.Error(err, "unable to fetch Job")
		return err
	}

	k8sJobDefinition, err := jobservice.CreateK8sJob(ctx, task)
	if err != nil {
		return err
	}

	for _, config := range task.Spec.Configs {
		secret, err := createSecret(config, task)
		if err != nil {
			rlog.Error("failed to create secret", err)
			task.Status.Phase = appsv1alpha1.PhaseError
			return err
		}
		err = assosiateSecretWithJob(ctx, r, req, task, secret)
		if err != nil {
			rlog.Error("failed to associate secret with job", err)
			task.Status.Phase = appsv1alpha1.PhaseError
			return err
		}
	}

	err = r.Create(ctx, k8sJobDefinition, &client.CreateOptions{})
	if err != nil {
		return err
	}

	err = ctrl.SetControllerReference(task, k8sJobDefinition, r.Scheme)
	if err != nil {
		return err
	}

	return nil
}

func listenForChangesInJobAndReport(ctx context.Context, r *TaskReconciler, req ctrl.Request, task *v1alpha1.Task) error {
	k8sClient, err := clients.Kubernetes.GetKubernetesClientset()
	if err != nil {
		panic(err.Error())
	}

	jobs := k8sClient.BatchV1().Jobs(req.Namespace)
	watchList, err := jobs.Watch(context.TODO(), metav1.SingleObject(metav1.ObjectMeta{
		Namespace: req.Namespace,
		Name:      task.Name,
	}))
	if err != nil {
		return err
	}

	for {
		jobChanges := <-watchList.ResultChan()
		if jobChanges.Object == nil {
			return fmt.Errorf("could not watch job changes, job event: %s", jobChanges.Type)
		}

		activejob := jobChanges.Object.(*batchv1.Job)
		rlog.Debug("job changed", rlog.String("job name", activejob.Name), rlog.Any("type", jobChanges.Type))

		if jobChanges.Type == "DELETED" {
			rlog.Debug("Job is deleted", rlog.String("job name", activejob.Name))
			return nil
		}

		rlog.Debug("job status", rlog.String("job name", activejob.Name), rlog.Any("status", activejob.Status))
		if activejob.Status.Active == 0 && activejob.Status.Succeeded == 0 && activejob.Status.Failed == 0 {
			rlog.Debug("job hasn't started yet", rlog.String("job name", activejob.Name))
		}

		if activejob.Status.Active > 0 {
			rlog.Debug("job is still running", rlog.String("job name", activejob.Name))
		}

		failedCount := activejob.Status.Failed
		if failedCount > 0 {
			if failedCount >= task.Spec.BackOffLimit {
				err := "task failed"
				if activejob.Status.Conditions != nil {
					latestCondition := activejob.Status.Conditions[len(activejob.Status.Conditions)-1]
					rlog.Error("job failed", fmt.Errorf(err), rlog.String("job name", activejob.Name))
					err = latestCondition.Reason
				}

				task.Status = v1alpha1.TaskStatus{
					Phase: v1alpha1.PhaseError,
					Failure: &v1alpha1.TaskFailure{
						Reason:    err,
						Timestamp: metav1.Now(),
					},
				}
				_ = updateTask(ctx, r, task)

				return fmt.Errorf("job '%s' failed: %s", activejob.Name, err)
			}

			if failedCount == 1 && len(activejob.Status.Conditions) > 0 && activejob.Status.Conditions[len(activejob.Status.Conditions)-1].Reason == "DeadlineExceeded" {
				err := activejob.Status.Conditions[len(activejob.Status.Conditions)-1].Reason

				task.Status = v1alpha1.TaskStatus{
					Phase: v1alpha1.PhaseError,
					Failure: &v1alpha1.TaskFailure{
						Reason:    err,
						Timestamp: metav1.Now(),
					},
				}
				_ = updateTask(ctx, r, task)
				return fmt.Errorf("job '%s' failed: %s", activejob.Name, err)
			}
		}

		if activejob.Status.Succeeded > 0 {
			rlog.Debug("job succeeded",
				rlog.String("job name", activejob.Name),
				rlog.String("image name", task.Spec.ImageName),
				rlog.String("image tag", task.Spec.ImageTag),
				rlog.String("image", activejob.Spec.Template.Spec.Containers[len(activejob.Spec.Template.Spec.Containers)-1].Image))
			hash, _ := generateHashForConfig(task.Spec)
			task.Status = v1alpha1.TaskStatus{
				Phase: v1alpha1.PhaseDone,
				Success: &v1alpha1.TaskSuccess{
					Timestamp: metav1.Now(),
					ConfigMd5: hash,
				},
			}
			_ = updateTask(ctx, r, task)
			return nil
		}
	}
}

func generateHashForConfig(jobConfig apicontracts.OperatorJob) (string, error) {
	configBytes, err := json.Marshal(jobConfig)
	if err != nil {
		return "", err
	}

	return stringhelper.GetSHA256Hash(configBytes), nil
}

func updateTask(ctx context.Context, r *TaskReconciler, task *v1alpha1.Task) error {
	err := r.Status().Update(ctx, task)
	return err
}

func createSecret(jobConfig apicontracts.OperatorJobConfig, task *v1alpha1.Task) (v1.Secret, error) {
	secretName := fmt.Sprintf("%s-%s-secret", task.Name, jobConfig.Name)
	secretNamespace := viper.GetString(configconsts.POD_NAMESPACE)

	data := make(map[string][]byte)
	for i, d := range jobConfig.Data {
		data[i] = []byte(d)
	}

	secret := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Data: data,
	}

	return secret, nil
}

func assosiateSecretWithJob(ctx context.Context, r *TaskReconciler, req ctrl.Request, task *v1alpha1.Task, secret v1.Secret) error {
	k8sLogger := log.FromContext(ctx)

	err := r.Get(ctx, req.NamespacedName, &secret)
	if !errors.IsNotFound(err) {
		k8sLogger.Error(err, "unable to fetch secret")
		return err
	}

	err = r.Create(ctx, &secret, &client.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		err = r.Update(ctx, &secret)
		if err != nil {
			return err
		}
	} else if err == nil {
		return err
	}

	err = ctrl.SetControllerReference(task, &secret, r.Scheme)
	if err != nil {
		return err
	}

	return nil
}
