package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/operator/clients/ror"
	"github.com/NorskHelsenett/ror/cmd/operator/services/taskservice"
	"sort"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"k8s.io/client-go/kubernetes"
)

func CheckForConfigChanges(k8sClient *kubernetes.Clientset) error {
	rlog.Info("check for config changes")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := ror.FetchConfiguration(ctx)
	if err != nil {
		rlog.Fatal("could not fetch configuration", err)
		return err
	}

	var operatorConfig apicontracts.OperatorConfig
	err = json.Unmarshal(response, &operatorConfig)
	if err != nil {
		rlog.Error("could not unmarshal operator config", err)
		return err
	}

	if operatorConfig.ApiVersion == "" || operatorConfig.Kind == "" {
		rlog.Error("failed to load config", fmt.Errorf("ror-operator config is empty"))
		return errors.New("empty ror-operator configuration")
	}

	if operatorConfig.Spec == nil {
		rlog.Error("failed to lead config", fmt.Errorf("operator config is missing spec"))
		return errors.New("spec missing from ror-operator configuration")
	}

	if len(operatorConfig.Spec.Tasks) < 1 {
		rlog.Error("failed to load config", fmt.Errorf("operator config is missing spec.tasks"))
		return errors.New("task spec missing from ror-operator configuration")
	}

	sort.Slice(operatorConfig.Spec.Tasks, func(i, j int) bool {
		return operatorConfig.Spec.Tasks[i].Index < operatorConfig.Spec.Tasks[j].Index
	})

	for _, job := range operatorConfig.Spec.Tasks {
		if job.Name == "" {
			rlog.Debug("Task not defined, continuing")
		}

		rlog.Debug("Task defined", rlog.Uint("index", job.Index),
			rlog.String("name", job.Name),
			rlog.String("version", job.Version))

		// TODO: run in gorutine with same index, but abort if one failes
		err := taskservice.CreateTasks(k8sClient, job)
		if err != nil {
			rlog.Error("task failed", err, rlog.String("name", job.Name))
			return err
		}
	}

	return nil
}
