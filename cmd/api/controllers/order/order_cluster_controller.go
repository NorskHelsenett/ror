package order

import (
	"github.com/NorskHelsenett/ror/cmd/api/customvalidators"
	orderservice "github.com/NorskHelsenett/ror/cmd/api/services/orderService"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate *validator.Validate
)

func init() {
	rlog.Debug("init cluster controller")
	validate = validator.New()
	customvalidators.Setup(validate)
}

// Order a kubernetes cluster by a apiresourcecontracts.ResourceClusterOrderSpec in the body
// Will only provide clusters the identity is authorized to views
//
//	@Summary	Order a kubernetes cluster
//	@Schemes
//	@Description	Order a kubernetes cluster
//	@Tags			orders
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apicontracts.PaginatedResult[apicontracts.Cluster]
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/orders/cluster	[post]
//	@Param			filter				body	apiresourcecontracts.ResourceClusterOrderSpec	true	"Filter"
//	@Security		ApiKey || AccessToken
func OrderCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = apicontracts.PaginatedResult[apicontracts.Cluster]{}
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var order apiresourcecontracts.ResourceClusterOrderSpec
		// Access check
		// Scope: ror
		// Subject: global
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			rlog.Errorc(ctx, "Missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&order); validationErr != nil {
			rlog.Errorc(ctx, validationErr.Error(), validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		rlog.Debugc(ctx, "cluster order request", rlog.Any("order", order))
		err := orderservice.OrderCluster(ctx, order)
		if err != nil {
			rlog.Errorc(ctx, "error ordering cluster", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}

// Order deletion of a kubernetes cluster by a apiresourcecontracts.ResourceClusterOrderSpec in the body
// Will only provide clusters the identity is authorized to view
//
//	@Summary	Order deletion a kubernetes cluster
//	@Schemes
//	@Description	Order a kubernetes cluster
//	@Tags			orders
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apicontracts.PaginatedResult[apicontracts.Cluster]
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/orders/cluster	[post]
//	@Param			filter				body	apiresourcecontracts.ResourceClusterOrderSpec	true	"Filter"
//	@Security		ApiKey || AccessToken
func DeleteCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		var order apiresourcecontracts.ResourceClusterOrderSpec
		// Access check
		// Scope: ror
		// Subject: global
		// Access: create
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Create {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		//validate the request body
		if err := c.BindJSON(&order); err != nil {
			rlog.Errorc(ctx, "Missing parameter", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Missing parameter",
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&order); validationErr != nil {
			rlog.Errorc(ctx, validationErr.Error(), validationErr)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: validationErr.Error(),
			})
			return
		}

		rlog.Debugc(ctx, "cluster order request", rlog.Any("order", order))
		// err := orderservice.OrderCluster(ctx, order)
		// if err != nil {
		// 	rlog.Errorc(ctx, "error ordering cluster", err)
		// 	c.JSON(http.StatusBadRequest, rorerror.RorError{
		// 		Status:  http.StatusBadRequest,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		c.JSON(http.StatusCreated, nil)
	}
}

// Get orders
//
//	@Summary	Get orders
//	@Schemes
//	@Description	Orders
//	@Tags			orders
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apicontracts.PaginatedResult[apicontracts.Cluster]
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/orders	[get]
//	@Security		ApiKey || AccessToken
func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: global
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		// TODO: need to be filtered
		orders, err := resourcesservice.GetClusterorders(ctx, apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeRor,
			Subject: string(aclmodels.Acl2RorSubjectGlobal),
		})
		if err != nil {
			rlog.Errorc(ctx, "error getting orders", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

// Get order by uid
//
//	@Summary	Get order by uid
//	@Schemes
//	@Description	Orders
//	@Tags			orders
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{object}	apiresourcecontracts.ResourceClusterorders
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/orders/{uid}	[get]
//	@Security		ApiKey || AccessToken
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()
		// Access check
		// Scope: ror
		// Subject: global
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		uid := c.Param("uid")
		if uid == "" || len(uid) == 0 {
			rlog.Errorc(ctx, "invalid id", nil)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		universalId, err := uuid.Parse(uid)
		if err != nil {
			rlog.Errorc(ctx, "invalid id", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		order, err := resourcesservice.GetClusterOrderByUid(ctx, apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeRor,
			Subject: string(aclmodels.Acl2RorSubjectGlobal),
		}, universalId.String())
		if err != nil {
			rlog.Errorc(ctx, "error getting orders", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

// Delete order
//
//	@Summary	Delete a order by uid
//	@Schemes
//	@Description	Orders
//	@Tags			orders
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200					{bool}		bool
//	@Failure		403					{object}	rorerror.RorError
//	@Failure		401					{object}	rorerror.RorError
//	@Failure		500					{object}	rorerror.RorError
//	@Router			/v1/orders/{uid}	[delete]
//	@Security		ApiKey || AccessToken
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		uid := c.Param("uid")
		if uid == "" || len(uid) == 0 {
			rlog.Errorc(ctx, "invalid id", nil)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}

		universalId, err := uuid.Parse(uid)
		if err != nil {
			rlog.Errorc(ctx, "invalid id", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: "Invalid id",
			})
			return
		}
		// Access check
		// Scope: ror
		// Subject: global
		// Access: delete
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectGlobal)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Delete {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		resource := apiresourcecontracts.ResourceUpdateModel{
			Uid: universalId.String(),
		}
		err = resourcesservice.ResourceDeleteService(ctx, resource)
		if err != nil {
			rlog.Errorc(ctx, "error getting orders", err)
			c.JSON(http.StatusBadRequest, rorerror.RorError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
