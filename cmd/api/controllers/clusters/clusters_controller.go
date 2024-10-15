// The clusters controller package provides controller functions for the /cluster and /clusters endpoints in the api V1.
package clusters

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	aclrepository "github.com/NorskHelsenett/ror/internal/acl/repositories"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

var (
	validate *validator.Validate
)

// Init is called to initialize the clusters controller
func init() {
	rlog.Debug("init clusters controller")
	validate = validator.New()
	customvalidators.Setup(validate)
}

// Get a apicontracts.Cluster by its clusterid.
// Identity must be authorized to view the requested cluster
//
//	@Summary	Get a cluster
//	@Schemes
//	@Description	Get a cluster by id
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	apicontracts.Cluster
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/cluster/{clusterid} [get]
//	@Router			/v1/clusters/{clusterid} [get]
//	@Security		ApiKey || AccessToken
func ClusterGetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		clusterId := c.Param("clusterid")
		defer cancel()
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

		// importing apicontracts for swagger
		var _ apicontracts.Cluster
		cluster, err := clustersservice.GetByClusterId(ctx, clusterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": "error when fetching data centers"})
			return
		}

		if cluster == nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusOK, cluster)
	}
}

// Check if clusterid exists.
// Identity must be authenticated
//
//	@Summary	ClusterId exists
//	@Schemes
//	@Description	Check if clusterId exists
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{bool}		bool
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/cluster/{clusterid}/exists [get]
//	@Router			/v1/clusters/{clusterid}/exists [get]
//	@Security		ApiKey || AccessToken
func ClusterExistsById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		clusterId := c.Param("clusterid")
		defer cancel()

		if clusterId == "" {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Missing clusterId"})
			return
		}

		exists, err := clustersservice.Exists(ctx, clusterId)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{"exists": false})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"exists": exists})
	}
}

// Get a []apicontracts.Cluster by a apicontracts.Filter object provided in the body.
// Will only provide clusters the identity is authorized to view
//
//	@Summary	Get clusters by filter
//	@Schemes
//	@Description	Get clusters by filter
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apicontracts.PaginatedResult[apicontracts.Cluster]
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/clusters/filter	[post]
//	@Param			filter				body	apicontracts.Filter	true	"Filter"
//	@Security		ApiKey || AccessToken
func ClusterByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var filter apicontracts.Filter

		//validate the request body
		if err := c.BindJSON(&filter); err != nil {
			rlog.Errorc(ctx, "missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&filter); validationErr != nil {
			rlog.Errorc(ctx, "could not validate required fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		// importing apicontracts for swagger
		var _ apicontracts.PaginatedResult[apicontracts.Cluster]

		paginatedResult, err := clustersservice.GetByFilter(ctx, &filter)
		if err != nil {
			rlog.Errorc(ctx, "could not get cluster service", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		if paginatedResult == nil {
			empty := apicontracts.PaginatedResult[apicontracts.Cluster]{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, paginatedResult)
	}
}

// Get a []apicontracts.Cluster by a workspaceName .
// Will only provide clusters the identity is authorized to view
//
//	@Summary	Get clusters by workspace
//	@Schemes
//	@Description	Get clusters by workspace
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200												{array}		apicontracts.Cluster
//	@Failure		403												{string}	Forbidden
//	@Failure		401												{string}	Unauthorized
//	@Failure		500												{string}	Failure	message
//	@Router			/v1/clusters/workspace/{workspaceName}/filter	[get]
//	@Param			filter											body	apicontracts.Filter	true	"Filter"
//	@Param			workspaceName									path	string				true	"workspaceName"
//	@Security		ApiKey || AccessToken
func ClusterGetByWorkspace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		workspaceName := c.Param("workspaceName")
		var filter apicontracts.Filter
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&filter); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&filter); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		for i := 0; i < len(filter.Sort); i++ {
			sort := filter.Sort[i]

			if validationErr := validate.Struct(sort); validationErr != nil {
				c.JSON(http.StatusBadRequest, rorerror.RorError{
					Status:  http.StatusBadRequest,
					Message: validationErr.Error(),
				})
				return
			}
		}

		// importing apicontracts for swagger
		var _ apicontracts.PaginatedResult[apicontracts.Cluster]

		result, err := clustersservice.GetByWorkspace(ctx, &filter, workspaceName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: "Could not get clusters by workspace",
			})
			return
		}

		if result == nil {
			empty := apicontracts.PaginatedResult[apicontracts.Cluster]{}
			c.JSON(http.StatusOK, empty)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// Get a list of metadata for clusters.
// TODO Verify: Will only provide clusters the identity is authorized to view
// TODO: Update swagger
// TODO: Check if this is used
//
//	@Summary	get metadata
//	@Schemes
//	@Description	Get cluster metadata
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{bool}		bool
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/metadata [get]
//	@Security		ApiKey || AccessToken
func GetMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		result, err := clustersservice.GetMetadata(ctx)
		if err != nil {
			rlog.Errorc(ctx, "could not get metadata", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// Update metadata for a cluster.
// Will only provide clusters the identity is authorized to view
//
//	@Summary	Update metadata
//	@Schemes
//	@Description	Update cluster metadata
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{bool}		bool
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/cluster/{clusterid}/metadata [patch]
//	@Router			/v1/clusters/{clusterid}/metadata [patch]
//	@Security		ApiKey || AccessToken
func UpdateMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		clusterId := c.Param("clusterid")
		var input apicontracts.ClusterMetadataModel
		defer cancel()

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

		if clusterId == "" {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Missing clusterId"})
			return
		}

		//validate the request body
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&input); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		exists, err := clustersservice.Exists(ctx, clusterId)
		if err != nil || !exists {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		clusters, err := clustersservice.GetByFilter(ctx, &apicontracts.Filter{
			Filters: []apicontracts.FilterMetadata{
				{
					Field:     "clusterid",
					MatchMode: apicontracts.MatchModeEquals,
					Value:     clusterId,
				},
			},
		})
		if err != nil || clusters.DataCount != 1 {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		err = clustersservice.UpdateMetadata(ctx, &input, clusters.Data[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "message": "could not update cluster metadata"})
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// Register a cluster hearbeat report, the hearbeat is in the payload.
// Parameter clusterid must match authorized clusterid
//
//	@Summary	Register cluster heartbeat
//	@Schemes
//	@Description	Registers a cluster heartbeat report
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{bool}		bool
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/cluster/heartbeat [post]
//	@Router			/v1/clusters/heartbeat [post]
//	@Param			heartbeat	body	apicontracts.Cluster	true	"Heartbeat"
//	@Security		ApiKey || AccessToken
func RegisterHeartbeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Heartbeat controller")
		defer span.End()
		var input apicontracts.Cluster

		_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Validate request")
		defer span1.End()

		//validate the request body
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&input); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Cluster{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		// Access check
		// Scope: cluster
		// Subject: input.ClusterId
		// Access: update
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeCluster, input.ClusterId)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Update {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}
		span1.End()

		_, span1 = otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "clustersservice.CreateOrUpdate")
		defer span1.End()
		err := clustersservice.CreateOrUpdate(ctx, &input, input.ClusterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		span1.End()
		_, span1 = otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Sending reply")
		defer span1.End()

		c.JSON(http.StatusCreated, responses.Cluster{Status: http.StatusCreated, Message: "success", Data: nil})
	}
}

// Get a list of control planes metadata for clusters.
//
//	@Summary	get control planes metadata
//	@Schemes
//	@Description	Get cluster control planes metadata
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{array}		apicontracts.ClusterControlPlaneMetadata
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/controlPlanesMetadata [get]
//	@Security		ApiKey || AccessToken
func GetControlPlanesMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var result []apicontracts.ClusterControlPlaneMetadata

		// Access check
		// Scope: ror
		// Subject: global
		// Access: delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		result, err := clustersservice.GetControlPlanesMetadata(ctx)
		if err != nil {
			rlog.Errorc(ctx, "could not get control planes", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// Get a kubeconfig by clusterid.
// Identity must be authorized to view the requested cluster
//
//	@Summary	Get kubeconfig for clusterId
//	@Schemes
//	@Description	Get a kubeconfig by clusterId
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Param			credentials				body	apicontracts.KubeconfigCredentials	true	"Credentials"
//	@Success		200	{object}	apicontracts.ClusterKubeconfig
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/login [post]
//	@Security		ApiKey || AccessToken
func GetKubeconfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		clusterid := c.Param("clusterid")
		if clusterid == "" {
			rlog.Errorc(ctx, "clusterid must be provided", nil)
			c.JSON(http.StatusBadRequest, "clusterid must be provided")
			return
		}
		defer cancel()
		// Access check
		// Scope: cluster
		// Subject: clusterId
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeCluster, clusterid)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			rlog.Errorc(ctx, "403: No access", nil)
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		//aclModel := aclmodels.NewAclV2QueryAccessScopeSubject(scope, clusterId)
		access := aclrepository.CheckAcl2ByCluster(ctx, accessQuery)
		var hasAccess bool = false
		for _, acl := range access {
			if acl.Kubernetes.Logon {
				hasAccess = true
			}
		}

		if !hasAccess {
			rlog.Errorc(ctx, "403: No access to login to cluster", nil)
			c.JSON(http.StatusForbidden, "403: No access to login to cluster")
			return
		}

		var clusterKubeConfigPayload apicontracts.KubeconfigCredentials
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&clusterKubeConfigPayload); err != nil {
			rlog.Errorc(ctx, "missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&clusterKubeConfigPayload); validationErr != nil {
			rlog.Errorc(ctx, "could not validate required fields", validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		var result apicontracts.ClusterKubeconfig
		kubeconfigString, err := clustersservice.GetKubeconfig(ctx, clusterid, clusterKubeConfigPayload)
		if err != nil {
			rlog.Errorc(ctx, "error when fetching kubeconfig", err)
			if strings.Contains(err.Error(), "is not supported") {
				rlog.Debugc(ctx, "provider not supported")
				result.Status = "error"
				result.Message = "provider not supported"
				c.JSON(http.StatusBadRequest, result)
			} else if strings.Contains(err.Error(), "could not find cluster") {
				rlog.Debugc(ctx, "cluster not found")
				result.Status = "error"
				result.Message = "cluster not found"
				c.JSON(http.StatusNotFound, result)
			} else {
				rlog.Errorc(ctx, "error when fetching kubeconfig", err)
				result.Status = "error"
				result.Message = "error when fetching kubeconfig"
				c.JSON(http.StatusInternalServerError, result)
			}
			return
		}

		if len(kubeconfigString) == 0 {
			rlog.Errorc(ctx, "error, since kubeconfig is empty", nil)
			result.Status = "error"
			result.Message = "error, since kubeconfig is empty"
			c.JSON(http.StatusNotFound, result)
			return
		}

		kubeConfigEncoded := base64.StdEncoding.EncodeToString([]byte(kubeconfigString))
		result.Status = "success"
		result.Message = ""
		result.Data = kubeConfigEncoded
		result.DataType = "base64"

		c.JSON(http.StatusOK, result)
	}
}

// Create a cluster
//
//	@Summary	Create a cluster
//	@Schemes
//	@Description	Create a cluster
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			credentials				body	apicontracts.Cluster	true	"Credentials"
//	@Success		200	{string}	ClusterId
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters [post]
//	@Security		ApiKey || AccessToken
func CreateCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: globalscope
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal) // TODO: what is correct here?
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		var clusterInput apicontracts.Cluster
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&clusterInput); err != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&clusterInput); validationErr != nil {
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		clusterId, err := clustersservice.Create(ctx, &clusterInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rorerror.RorError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Could not create cluster (%s)", err),
			})
			return
		}

		c.JSON(http.StatusOK, clusterId)
	}
}
