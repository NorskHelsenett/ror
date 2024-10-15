package tasks

import (
	"fmt"
	tasksservice "github.com/NorskHelsenett/ror/cmd/api/services/tasksService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init tasks controller")
	validate = validator.New()
}

// @Summary	Get a task
// @Schemes
// @Description	Get a task by id
// @Tags			tasks
// @Accept			application/json
// @Produce		application/json
// @Param			id		path		string				true	"id"
// @Param			task	body		apicontracts.Task	true	"Get a task"
// @Success		200		{object}	apicontracts.Task
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v1/tasks/:id [get]
// @Security		ApiKey || AccessToken
func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		_, err := gincontext.GetUserFromGinContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rorerror.RorError{
				Status:  http.StatusUnauthorized,
				Message: "Could not fetch user",
			})
			return
		}

		taskId := c.Param("id")
		if taskId == "" || len(taskId) == 0 {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "invalid task id",
			})
			return
		}

		result, err := tasksservice.GetById(ctx, taskId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "could not get task",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// @Summary	Get tasks
// @Schemes
// @Description	Get all tasks
// @Tags			tasks
// @Accept			application/json
// @Produce		application/json
// @Success		200			{array}		apicontracts.Task
// @Failure		403			{string}	Forbidden
// @Failure		401			{string}	Unauthorized
// @Failure		500			{string}	Failure	message
// @Router			/v1/tasks	[get]
// @Security		ApiKey || AccessToken
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: acl
		// Access: read
		// TODO: check if this is the right way to do it
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectAcl)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		tasks, err := tasksservice.GetAll(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not find tasks ...",
			})
		}

		c.JSON(http.StatusOK, tasks)
	}
}

// @Summary	Create a task
// @Schemes
// @Description	Create a task
// @Tags			tasks
// @Accept			application/json
// @Produce		application/json
// @Param			task	body		apicontracts.Task	true	"Add a task"
// @Success		200		{array}		apicontracts.Task
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v1/tasks [post]
// @Security		ApiKey || AccessToken
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: acl
		// Access: create
		// TODO: check if this is the right way to do it
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectAcl)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var task apicontracts.Task
		//validate the request body
		if err := c.BindJSON(&task); err != nil {
			rlog.Errorc(ctx, "could not bind JSON", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not validate task object",
			})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&task); err != nil {
			rlog.Errorc(ctx, "could not validate object", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf("Required fields are missing: %s", err),
			})
			return
		}

		createdTask, err := tasksservice.Create(ctx, &task)
		if err != nil {
			rlog.Errorc(ctx, "could not create task", err)
			if strings.Contains(err.Error(), "exists") {
				c.JSON(http.StatusBadRequest, rorerror.RorError{
					Status:  http.StatusBadRequest,
					Message: "Already exists",
				})
				return
			}

			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields are missing",
			})
			return
		}

		c.Set("newObject", createdTask)

		c.JSON(http.StatusOK, createdTask)
	}
}

// @Summary	Update a task
// @Schemes
// @Description	Update a task by id
// @Tags			tasks
// @Accept			application/json
// @Produce		application/json
// @Param			id		path		string				true	"id"
// @Param			task	body		apicontracts.Task	true	"Update task"
// @Success		200		{object}	apicontracts.Task
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v1/tasks/:id [put]
// @Security		ApiKey || AccessToken
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var taskInput apicontracts.Task

		taskId := c.Param("id")
		if taskId == "" || len(taskId) == 0 {
			rlog.Errorc(ctx, "invalid task id", fmt.Errorf("id is zero length"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid task id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: acl
		// Access: update
		// TODO: check if this is the right way to do it
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectAcl)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		//validate the request body
		if err := c.BindJSON(&taskInput); err != nil {
			rlog.Errorc(ctx, "could not bind json", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Object is not valid",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&taskInput); validationErr != nil {
			rlog.Errorc(ctx, "could not validate reqired fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Required fields missing",
			})
			return
		}

		updatedTask, originalTask, err := tasksservice.Update(ctx, taskId, &taskInput)
		if err != nil {
			rlog.Errorc(ctx, "could not update task", err)
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not update task",
			})
			return
		}

		if updatedTask == nil {
			rlog.Errorc(ctx, "Could not update task", fmt.Errorf("task does not exist"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not update task, does it exist?!",
			})
			return
		}

		c.Set("newObject", updatedTask)
		c.Set("oldObject", originalTask)

		c.JSON(http.StatusOK, updatedTask)
	}
}

// @Summary	Delete a task
// @Schemes
// @Description	Delete a task by id
// @Tags			tasks
// @Accept			application/json
// @Produce		application/json
// @Param			id	path		string	true	"id"
// @Success		200	{bool}		true
// @Failure		403	{string}	Forbidden
// @Failure		401	{string}	Unauthorized
// @Failure		500	{string}	Failure	message
// @Router			/v1/tasks/:id [delete]
// @Security		ApiKey || AccessToken
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		taskId := c.Param("taskId")
		if taskId == "" || len(taskId) == 0 {
			rlog.Errorc(ctx, "invalid id", fmt.Errorf("id is zero lenght"))
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: acl
		// Access: delete
		// TODO: check if this is the right way to do it
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectAcl)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		result, deletedTask, err := tasksservice.Delete(ctx, taskId)
		if err != nil {
			rlog.Errorc(ctx, "could not delete task", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Could not delete task",
			})
			return
		}

		c.Set("oldObject", deletedTask)

		c.JSON(http.StatusOK, result)
	}
}
