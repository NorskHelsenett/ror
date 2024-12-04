package clusters

import (
	"net/http"

	"github.com/NorskHelsenett/ror/cmd/api/responses"
	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	aclservice "github.com/NorskHelsenett/ror/internal/acl/services"
	"github.com/NorskHelsenett/ror/internal/models/viewsmodels"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
)

// Placeholder for views not yet implemented
//
//	@Summary	This is a dummy view
//	@Schemes
//	@Description	Just a dummy view
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{string}	This	is		not	the	view	you	are	looking	for
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/views/ingresses [get]
//	@Router			/v1/clusters/{clusterid}/views/nodes [get]
//	@Router			/v1/clusters/{clusterid}/views/applications [get]
//	@Router			/v1/clusters/views/errorlist [get]
//	@Router			/v1/clusters/views/clusterlist [get]
//	@Security		ApiKey || AccessToken
func DummyView() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		c.JSON(http.StatusOK, "{\"message\":\"this is not the view you are looking for\"}")
	}
}

// View for policy reports
//
//	@Summary	Policy report view
//	@Schemes
//	@Description	A structured presentation of policyreports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{string}	This	is		not	the	view	you	are	looking	for
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/views/policyreports [get]
//	@Security		ApiKey || AccessToken
func PolicyreportsView() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		clusterid := c.Param("clusterid")
		if clusterid == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		ownerref := apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeCluster,
			Subject: clusterid,
		}

		accessObject := aclservice.CheckAccessByOwnerref(ctx, ownerref)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		policyreport, err := clustersservice.GetViewPolicyreport(ctx, ownerref)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, policyreport)
	}
}

// View for policy reports
//
//	@Summary	Policy report view
//	@Schemes
//	@Description	A structured presentation of policyreports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			type	query		viewsmodels.PolicyreportGlobalQueryType	true	"type"
//	@Success		200		{string}	This									is		not	the	view	you	are	looking	for
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/views/policyreports [get]
//	@Security		ApiKey || AccessToken
func PolicyreportSummaryView() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		query := viewsmodels.PolicyreportGlobalQueryType(c.Query("type"))

		if query == "" {
			c.JSON(http.StatusBadRequest, "No type defined")
			return
		}
		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		//TODO: investegate why this worked (cluster with upper C)
		//accessObject := aclservice.CheckAccessByContextScopeSubject(ctx, aclmodels.Acl2ScopeRor, "Cluster")
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		clusterID := c.Query("clusterid")

		policyreport, err := clustersservice.GetViewPolicyReportSummary(ctx, query, clusterID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, policyreport)
	}
}

// View for policy reports
//
//	@Summary	Policy report view
//	@Schemes
//	@Description	A structured presentation of policyreports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			type	query		viewsmodels.PolicyreportGlobalQueryType	true	"type"
//	@Success		200		{string}	This									is		not	the	view	you	are	looking	for
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/views/policyreports [get]
//	@Security		ApiKey || AccessToken
func VulnerabilityreportSummaryView() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		query := viewsmodels.PolicyreportGlobalQueryType(c.Query("type"))

		if query == "" {
			c.JSON(http.StatusBadRequest, "No type defined")
			return
		}

		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		clusterID := c.Query("clusterid")

		policyreport, err := clustersservice.GetViewPolicyReportSummary(ctx, query, clusterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, policyreport)
	}
}

// View for vulnerability reports
//
//	@Summary	Vulnerability reports view
//	@Schemes
//	@Description	A structured presentation of vulnerability reports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id	path		string	true	"id"
//	@Success		200	{string}	This	is		not	the	view	you	are	looking	for
//	@Failure		403	{string}	Forbidden
//	@Failure		401	{string}	Unauthorized
//	@Failure		500	{string}	Failure	message
//	@Router			/v1/clusters/{clusterid}/views/vulnerabilityreports [get]
//	@Security		ApiKey || AccessToken
func VulnerabilityReportsView() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		clusterid := c.Param("clusterid")
		if clusterid == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		ownerref := apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeCluster,
			Subject: clusterid,
		}

		accessObject := aclservice.CheckAccessByOwnerref(ctx, ownerref)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		vulnerabilityreports, err := clustersservice.GetViewVulnerabilityReports(ctx, ownerref)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, vulnerabilityreports)
	}
}

// View for vulnerability reports
//
//	@Summary	Vulnerability reports view
//	@Schemes
//	@Description	A structured presentation of vulnerability reports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			cveid	path		string	true	"cveid"
//	@Success		200		{string}	This	is		not	the	view	you	are	looking	for
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/views/vulnerabilityreports/byid/:cveid [get]
//	@Security		ApiKey || AccessToken
func VulnerabilityReportsViewById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		cveId := c.Param("cveid")
		if cveId == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		vulnerabilityreports, err := clustersservice.GetViewVulnerabilityReportsById(ctx, cveId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, vulnerabilityreports)
	}
}

// VulnerabilityReportsGlobal godoc
//
//	@Summary		Get vulnerability reports summary per cluster
//	@Schemes
//	@Description	Shows a summary of trivy vulnerability reports per cluster categorized by amount of critical/high/medium/low vulnerabilities.
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200		{object}	[]viewsmodels.VulnerabilityReportsView
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/views/vulnerabilityreports [get]
//	@Security		ApiKey || AccessToken
func VulnerabilityReportsGlobal() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		vulnerabilityreports, err := clustersservice.GetViewVulnerabilityReportsGlobal(ctx)
		if err != nil {
			rlog.Error("error while fetching global vulnerability reports: %w", err)
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, vulnerabilityreports)
	}
}

// View for vulnerability reports by CVE ID
//
//	@Summary	Vulnerability reports view by CVE ID
//	@Schemes
//	@Description	A structured presentation of vulnerability reports by CVE ID
//	@Accept			application/json
//	@Produce		application/json
//	@Query			cveid	{string}
//	@Success		200		{string}	This	is		not	the	view	you	are	looking	for
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/views/vulnerabilityreports/byid [get]
//	@Security		ApiKey || AccessToken
func GlobalVulnerabilityReportsViewById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		cveId := c.Query("cveid")
		if cveId == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		vulnerabilityreports, err := clustersservice.GetGlobalViewVulnerabilityReportsById(ctx, cveId)
		if err != nil {
			rlog.Errorc(ctx, "Error while getting global vulnerabilityreportview by CVE ID", err)
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, vulnerabilityreports)
	}
}

// View for compliance reports
//
//	@Summary	Compliance reports view
//	@Schemes
//	@Description	A structured presentation of compliance reports
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Param			clusterid	path	string	true	"clusterid"
//	@Success		200		{array}		viewsmodels.ComplianceReport
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/:clusterid/views/compliancereports [get]
//	@Security		ApiKey || AccessToken
func ComplianceReports() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		clusterId := c.Param("clusterid")
		if clusterId == "" {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		ownerref := apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeCluster,
			Subject: clusterId,
		}

		accessObject := aclservice.CheckAccessByOwnerref(ctx, ownerref)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		complianceReports, err := clustersservice.GetClusterComplianceReports(ctx, clusterId)
		if err != nil {
			rlog.Error("could not get compliance reports for cluster "+clusterId, err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, complianceReports)
	}
}

// ComplianceReportsGlobal godoc
//
//	@Summary		Get compliance reports summary per cluster
//	@Schemes
//	@Description	Shows a summary of trivy compliance reports per cluster categorized by amount failed or passed.
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200		{object}	[]viewsmodels.ComplianceReport
//	@Failure		403		{string}	Forbidden
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	Failure	message
//	@Router			/v1/clusters/views/compliancereports [get]
//	@Security		ApiKey || AccessToken
func ComplianceReportsGlobal() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		// Access check
		// Scope: ror
		// Subject: cluster
		// Access: read
		accessQuery := aclmodels.NewAclV2QueryAccessScopeSubject(aclmodels.Acl2ScopeRor, aclmodels.Acl2RorSubjectCluster)
		accessObject := aclservice.CheckAccessByContextAclQuery(ctx, accessQuery)
		if !accessObject.Read {
			c.JSON(http.StatusForbidden, "403: No access")
			return
		}

		complianceReports, err := clustersservice.GetClusterComplianceReportsGlobal(ctx)
		if err != nil {
			rlog.Error("error while fetching global vulnerability reports: %w", err)
			c.JSON(http.StatusInternalServerError, responses.Cluster{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, complianceReports)
	}
}

//	@Summary		Get cluster self data
//	@Schemes		https
//	@Description	Get cluster self data
//	@Tags			clusters
//	@Accept			application/json
//	@Produce		application/json

// @Success	200	{string}	Get	data	for	the	cluster
// @Failure	403	{string}	Forbidden
// @Failure	401	{string}	Unauthorized
// @Failure	500	{string}	Failure	message
// @Router		/v1/clusters/self [get]
// @Security	ApiKey || AccessToken
func GetSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gincontext.GetRorContextFromGinContext(c)
		defer cancel()

		identity := rorcontext.GetIdentityFromRorContext(ctx)
		if !identity.IsCluster() {
			c.JSON(http.StatusForbidden, "wrong identity type")
			return
		}

		// TODO: Should this use aclservice.CheckAccessByAccessQuery?

		clusters, err := clustersservice.GetByFilter(ctx, &apicontracts.Filter{
			Filters: []apicontracts.FilterMetadata{
				{
					Field:     "clusterid",
					MatchMode: apicontracts.MatchModeEquals,
					Value:     identity.ClusterIdentity.Id,
				},
			},
		})
		if err != nil {
			c.JSON(http.StatusNotFound, "")
			return
		}

		foundWithOldId := false
		if clusters.TotalCount == 0 {
			clusters, err = clustersservice.GetByFilter(ctx, &apicontracts.Filter{
				Filters: []apicontracts.FilterMetadata{
					{
						Field:     "clusteridold",
						MatchMode: apicontracts.MatchModeEquals,
						Value:     identity.ClusterIdentity.Id,
					},
				},
			})

			if err != nil {
				c.JSON(http.StatusNotFound, "")
				return
			}
			foundWithOldId = true
		}

		if clusters.TotalCount > 1 {
			c.JSON(http.StatusNotFound, "Multiple clusters found")
			return
		}

		if len(clusters.Data) == 0 {
			c.JSON(http.StatusNotFound, "")
			return
		}

		clusterId := clusters.Data[0].ClusterId
		if foundWithOldId {
			clusterId = clusters.Data[0].Identifier
		}

		result := apicontracts.ClusterSelf{
			ClusterId: clusterId,
		}
		c.JSON(http.StatusOK, result)
	}
}
