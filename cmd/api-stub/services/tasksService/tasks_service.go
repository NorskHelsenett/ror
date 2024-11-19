package tasksservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"
	tasksrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/tasksRepo"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	semverv4 "github.com/blang/semver/v4"
)

func GetAll(ctx context.Context) (*[]apicontracts.Task, error) {
	tasks, err := tasksrepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("could not get tasks")
	}

	return tasks, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.Task, error) {
	tasks, err := tasksrepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("could not get task by id")
	}

	return tasks, nil
}

func GetByProperty(ctx context.Context, propertyName string, propertyValue string) (*apicontracts.Task, error) {
	task, err := tasksrepo.FindOne(ctx, propertyName, propertyValue)
	if err != nil {
		return nil, fmt.Errorf("could not get task by property %s, and value: %s", propertyName, propertyValue)
	}

	return task, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.Task], error) {
	result, err := tasksrepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not get tasks: %v", err)
	}

	return result, nil
}

func Create(ctx context.Context, taskInput *apicontracts.Task) (*apicontracts.Task, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	exists, err := tasksrepo.FindOne(ctx, "name", taskInput.Name)
	if err != nil {
		return exists, fmt.Errorf("could not check if task exists: %v", err)
	}

	if exists != nil {
		return exists, fmt.Errorf("task already exists")
	}

	if !(strings.HasPrefix(taskInput.Config.Version, "sha") && strings.Contains(taskInput.Config.Version, ":")) {
		version, err := semverv4.Parse(taskInput.Config.Version)
		rlog.Debugc(ctx, "version", rlog.Any("version", version))
		if err != nil {
			return nil, fmt.Errorf("task.config.version is not valid: %s", taskInput.Config.Version)
		}
	}

	createdTask, err := tasksrepo.Create(ctx, taskInput)
	if err != nil {
		return nil, fmt.Errorf("could not create task: %v", err)
	}

	_, err = auditlog.Create(ctx, "New task created", models.AuditCategoryConfiguration, models.AuditActionCreate, identity.User, createdTask, nil)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	return createdTask, nil
}

func Update(ctx context.Context, taskId string, taskInput *apicontracts.Task) (*apicontracts.Task, *apicontracts.Task, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)

	if !(strings.HasPrefix(taskInput.Config.Version, "sha") && strings.Contains(taskInput.Config.Version, ":")) {
		_, err := semverv4.Parse(taskInput.Config.Version)
		if err != nil {
			return nil, nil, fmt.Errorf("task.version is not valid: %s", taskInput.Config.Version)
		}
	}

	newObject, oldObject, err := tasksrepo.Update(ctx, taskId, *taskInput)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update task: %v", err)
	}

	_, err = auditlog.Create(ctx, "Task updated", models.AuditCategoryConfiguration, models.AuditActionUpdate, identity.User, newObject, oldObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not audit log update action: %v", err)
	}

	return newObject, oldObject, nil
}

func Delete(ctx context.Context, taskId string) (bool, *apicontracts.Task, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	deleted, deletedTask, err := tasksrepo.Delete(ctx, taskId)
	if err != nil {
		return false, nil, fmt.Errorf("could not delete task: %v", err)
	}

	_, err = auditlog.Create(ctx, "Task deleted", models.AuditCategoryConfiguration, models.AuditActionDelete, identity.User, nil, deleted)
	if err != nil {
		return false, nil, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return deleted, deletedTask, nil
}
