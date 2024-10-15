// TODO: Describe package
package configurationcontroller

import (
	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	configurationservice "github.com/NorskHelsenett/ror/cmd/api/services/configurationService"
	operatorconfigservice "github.com/NorskHelsenett/ror/cmd/api/services/operatorConfigService"
	tasksservice "github.com/NorskHelsenett/ror/cmd/api/services/tasksService"
	"net/http"

	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"strings"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
)

func init() {
	rlog.Debug("init m2m token controller")
}

// GetOperatorConfiguration returns a gin.HandlerFunc that handles GET requests for operator configuration.
// It checks the access of the user and retrieves the operator configuration based on the cluster ID.
// If the operator configuration is not found, it returns a 404 error.
// If the cluster configuration is found, it checks the versions of the tasks and returns the operator configuration.
func GetOperatorConfiguration() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		clusterId := identity.GetId()

		// Access check
		// Scope: cluster
		// Subject: clusterId
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeCluster, clusterId)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}
		operatorConfigs, err := operatorconfigservice.GetByFilter(ctx, &apicontracts.Filter{
			Filters: []apicontracts.FilterMetadata{
				{
					Field:     "kind",
					MatchMode: apicontracts.MatchModeEquals,
					Value:     "ror-operator",
				},
				{
					Field:     "apiversion",
					MatchMode: apicontracts.MatchModeEquals,
					Value:     "github.com/NorskHelsenett/ror/v1/config",
				},
			},
		})

		if err != nil {
			rlog.Errorc(ctx, "Error when fetching OperatorConfig", err)
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Could not find operator config",
			})
			return
		}

		if operatorConfigs.DataCount == 0 || operatorConfigs.DataCount >= 2 {
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Could not find operator config",
			})
			return
		}

		config := operatorConfigs.Data[0]
		cluster, _ := clustersservice.GetByClusterId(ctx, clusterId)
		if cluster.Workspace.Datacenter.Provider == providers.ProviderTypeKind {
			c.JSON(http.StatusOK, config)
			return
		} else if cluster == nil || cluster.Config.Versions == nil {
			c.JSON(http.StatusOK, config)
			return
		}

		c.JSON(http.StatusOK, config)
	}
}

// GetTaskConfiguration returns a gin.HandlerFunc that handles GET requests for task configuration.
// It checks the access of the user and retrieves the task configuration based on the cluster ID and task name.
// If the task configuration is not found, it returns a 404 error.
// If the task configuration is found, it returns the task configuration (apicontracts.OperatorJob).
func GetTaskConfiguration() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		name := c.Param("name")
		defer cancel()

		name = strings.TrimSpace(name)

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		clusterId := identity.GetId()

		// Access check
		// Scope: cluster
		// Subject: clusterId
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeCluster, clusterId)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		if name == "" {
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Missing name",
			})
			return
		}

		tasks, err := tasksservice.GetByFilter(ctx, &apicontracts.Filter{
			Filters: []apicontracts.FilterMetadata{
				{
					Field:     "name",
					MatchMode: apicontracts.MatchModeEquals,
					Value:     name,
				},
			},
		})

		if err != nil {
			rlog.Errorc(ctx, "Error when fetching Task", err)
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Task not found",
			})
			return
		}

		if tasks.DataCount == 0 || tasks.DataCount >= 2 {
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Task not found",
			})
			return
		}

		task := tasks.Data[0]
		if len(task.Name) == 0 {
			c.JSON(http.StatusNotFound, rorerror.RorError{
				Status:  http.StatusNotFound,
				Message: "Task not found",
			})
			return
		}

		taskSpec, err := configurationservice.GetTaskConfigByClusterIdAndTaskName(ctx, &task, clusterId)
		if err != nil {
			rlog.Errorc(ctx, "Error when fetching Task.Spec", err)
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Error when fetching Task.Spec",
			})
			return
		}

		c.JSON(http.StatusOK, taskSpec)
	}
}
