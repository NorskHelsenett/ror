// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package resourcescontroller

import (
	"github.com/NorskHelsenett/ror/cmd/api/responses"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/gin-gonic/gin"
)

// Get a list of cluster resources og given group/version/kind.
//
// @Summary	Get resources
// @Schemes
// @Description	Get a list of resources
// @Tags			resources
// @Accept			application/json
// @Produce		application/json
// @Param			ownerScope	query		aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
// @Param			ownerSubject	query		string	true	"The name og the owner"
// @Param			apiversion	query		string	true	"ApiVersion"
// @Param			kind	query		string	true	"Kind"
// @Success		200		{array}		apiresourcecontracts.ResourceNode
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v1/resources [get]
// @Security		ApiKey || AccessToken
func GetResources() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		query := apiresourcecontracts.NewResourceQueryFromClient(c)

		accessObject := aclservice.CheckAccessByOwnerref(ctx, query.Owner)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "")
			return
		}

		if query.ApiVersion == "v1" && query.Kind == "Namespace" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceNamespace](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Node" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceNode](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "PersistentVolumeClaim" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourcePersistentVolumeClaim](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "Deployment" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceDeployment](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "storage.k8s.io/v1" && query.Kind == "StorageClass" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceStorageClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "wgpolicyk8s.io/v1alpha2" && query.Kind == "PolicyReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourcePolicyReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "argoproj.io/v1alpha1" && query.Kind == "Application" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceApplication](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "argoproj.io/v1alpha1" && query.Kind == "AppProject" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceAppProject](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "cert-manager.io/v1" && query.Kind == "Certificate" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceCertificate](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Service" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceService](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Pod" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourcePod](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "ReplicaSet" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceReplicaSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "StatefulSet" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceStatefulSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "DaemonSet" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceDaemonSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "networking.k8s.io/v1" && query.Kind == "Ingress" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceIngress](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "networking.k8s.io/v1" && query.Kind == "IngressClass" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceIngressClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "VulnerabilityReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceVulnerabilityReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ExposedSecretReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceExposedSecretReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ConfigAuditReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceConfigAuditReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "RbacAssessmentReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceRbacAssessmentReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && query.Kind == "TanzuKubernetesCluster" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceTanzuKubernetesCluster](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && query.Kind == "TanzuKubernetesRelease" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceTanzuKubernetesRelease](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "vmoperator.vmware.com/v1alpha1" && query.Kind == "VirtualMachineClass" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceVirtualMachineClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "vmoperator.vmware.com/v1alpha1" && query.Kind == "VirtualMachineClassBinding" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceVirtualMachineClassBinding](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "KubernetesCluster" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceKubernetesCluster](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "ClusterOrder" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceClusterOrder](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Project" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceProject](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Configuration" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceConfiguration](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ClusterComplianceReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceClusterComplianceReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "ClusterVulnerabilityReport" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceClusterVulnerabilityReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Route" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceRoute](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "SlackMessage" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceSlackMessage](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "VulnerabilityEvent" {
			resources, err := resourcesservice.GetResources[apiresourcecontracts.ResourceVulnerabilityEvent](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
	}
}

// Get a cluster resources og given group/version/kind/uid.
//
// @Summary	Get resource
// @Schemes
// @Description	Get a resource by uid
// @Tags			resources
// @Accept			application/json
// @Produce		application/json
// @Param			uid	path		string	true	"The uid of the resource"
// @Param			ownerScope	query		aclmodels.Acl2Scope	true	"The kind of the owner, currently only support 'Cluster'"
// @Param			ownerSubject	query		string	true	"The name og the owner"
// @Param			apiversion	query		string	true	"ApiVersion"
// @Param			kind	query		string	true	"Kind"
// @Success		200		{array}		apiresourcecontracts.ResourceNode
// @Failure		403		{string}	Forbidden
// @Failure		401		{string}	Unauthorized
// @Failure		500		{string}	Failure	message
// @Router			/v1/resource/uid/{uid} [get]
// @Security		ApiKey || AccessToken
func GetResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		if c.Param("uid") == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		query := apiresourcecontracts.ResourceQuery{
			Owner: apiresourcecontracts.ResourceOwnerReference{
				Scope:   aclmodels.Acl2Scope(c.Query("ownerScope")),
				Subject: c.Query("ownerSubject"),
			},
			Kind:       c.Query("kind"),
			ApiVersion: c.Query("apiversion"),
			Uid:        c.Param("uid"),
		}

		accessObject := aclservice.CheckAccessByOwnerref(ctx, query.Owner)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "")
			return
		}

		if query.ApiVersion == "v1" && query.Kind == "Namespace" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceNamespace](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Node" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceNode](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "PersistentVolumeClaim" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourcePersistentVolumeClaim](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "Deployment" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceDeployment](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "storage.k8s.io/v1" && query.Kind == "StorageClass" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceStorageClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "wgpolicyk8s.io/v1alpha2" && query.Kind == "PolicyReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourcePolicyReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "argoproj.io/v1alpha1" && query.Kind == "Application" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceApplication](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "argoproj.io/v1alpha1" && query.Kind == "AppProject" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceAppProject](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "cert-manager.io/v1" && query.Kind == "Certificate" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceCertificate](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Service" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceService](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "v1" && query.Kind == "Pod" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourcePod](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "ReplicaSet" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceReplicaSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "StatefulSet" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceStatefulSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "apps/v1" && query.Kind == "DaemonSet" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceDaemonSet](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "networking.k8s.io/v1" && query.Kind == "Ingress" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceIngress](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "networking.k8s.io/v1" && query.Kind == "IngressClass" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceIngressClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "VulnerabilityReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceVulnerabilityReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ExposedSecretReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceExposedSecretReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ConfigAuditReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceConfigAuditReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "RbacAssessmentReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceRbacAssessmentReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && query.Kind == "TanzuKubernetesCluster" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceTanzuKubernetesCluster](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && query.Kind == "TanzuKubernetesRelease" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceTanzuKubernetesRelease](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "vmoperator.vmware.com/v1alpha1" && query.Kind == "VirtualMachineClass" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceVirtualMachineClass](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "vmoperator.vmware.com/v1alpha1" && query.Kind == "VirtualMachineClassBinding" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceVirtualMachineClassBinding](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "KubernetesCluster" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceKubernetesCluster](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "ClusterOrder" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceClusterOrder](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Project" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceProject](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Configuration" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceConfiguration](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "aquasecurity.github.io/v1alpha1" && query.Kind == "ClusterComplianceReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceClusterComplianceReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "ClusterVulnerabilityReport" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceClusterVulnerabilityReport](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "Route" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceRoute](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "SlackMessage" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceSlackMessage](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
		if query.ApiVersion == "general.ror.internal/v1alpha1" && query.Kind == "VulnerabilityEvent" {
			resources, err := resourcesservice.GetResource[apiresourcecontracts.ResourceVulnerabilityEvent](ctx, query)
			if err != nil {
				c.JSON(http.StatusNotFound, responses.Cluster{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusOK, resources)
		}
	}
}
