package routes

import (
	"github.com/NorskHelsenett/ror/cmd/api/auth"
	ctrlAcl "github.com/NorskHelsenett/ror/cmd/api/controllers/acl"
	ctrlApikeys "github.com/NorskHelsenett/ror/cmd/api/controllers/apikeys"
	ctrlAuditlogs "github.com/NorskHelsenett/ror/cmd/api/controllers/auditlogs"
	ctrlClusters "github.com/NorskHelsenett/ror/cmd/api/controllers/clusters"
	ctrlDatacenters "github.com/NorskHelsenett/ror/cmd/api/controllers/datacenters"
	ctrlDesiredVersion "github.com/NorskHelsenett/ror/cmd/api/controllers/desired_version"
	ctrlHealth "github.com/NorskHelsenett/ror/cmd/api/controllers/health"
	"github.com/NorskHelsenett/ror/cmd/api/controllers/infocontroller"
	ctrlM2mConfiguration "github.com/NorskHelsenett/ror/cmd/api/controllers/m2m/configurationcontroller"
	ctrlM2mEaster "github.com/NorskHelsenett/ror/cmd/api/controllers/m2m/easter"
	ctrlMetrics "github.com/NorskHelsenett/ror/cmd/api/controllers/metrics"
	"github.com/NorskHelsenett/ror/cmd/api/controllers/notinuse"
	ctrlOperatorConfigs "github.com/NorskHelsenett/ror/cmd/api/controllers/operatorconfigs"
	ctrlOrder "github.com/NorskHelsenett/ror/cmd/api/controllers/order"
	ctrlPrices "github.com/NorskHelsenett/ror/cmd/api/controllers/prices"
	ctrlProjects "github.com/NorskHelsenett/ror/cmd/api/controllers/projects"
	ctrlProviders "github.com/NorskHelsenett/ror/cmd/api/controllers/providers"
	"github.com/NorskHelsenett/ror/cmd/api/controllers/resourcescontroller"
	ctrlRulesets "github.com/NorskHelsenett/ror/cmd/api/controllers/rulesetsController"
	ctrlTasks "github.com/NorskHelsenett/ror/cmd/api/controllers/tasks"
	ctrlUsers "github.com/NorskHelsenett/ror/cmd/api/controllers/users"
	v2resourcescontroller "github.com/NorskHelsenett/ror/cmd/api/controllers/v2/resourcescontroller"
	ctrlWorkspaces "github.com/NorskHelsenett/ror/cmd/api/controllers/workspaces"
	"time"

	"github.com/NorskHelsenett/ror/cmd/api/controllers/v2/handlerv2self"
	"github.com/NorskHelsenett/ror/cmd/api/webserver/middlewares"
	"github.com/NorskHelsenett/ror/cmd/api/webserver/sse"
	"github.com/NorskHelsenett/ror/internal/models"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/cmd/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {

	timeoutduration, err := time.ParseDuration(viper.GetString(configconsts.HTTP_TIMEOUT))
	if err != nil {
		rlog.Error("Could not parse timeout duration", err)
		timeoutduration = 15 * time.Second
	}

	v1 := router.Group("/v1")
	{
		eventsRoute := v1.Group("events", auth.AuthenticationMiddleware)
		{
			eventsRoute.GET("listen", middlewares.HeadersMiddleware(), sse.Server.HandleSSE())
			eventsRoute.POST("send", sse.Server.Send())
		}
		clusterloginRoute := v1.Group("clusters")
		{
			logintimeoutduration := 60 * time.Second
			clusterloginRoute.Use(middlewares.TimeoutMiddleware(logintimeoutduration))
			clusterloginRoute.Use(auth.AuthenticationMiddleware)
			clusterloginRoute.POST("/:clusterid/login", ctrlClusters.GetKubeconfig())
		}

		v1.Use(middlewares.TimeoutMiddleware(timeoutduration))
		// allow anonymous, for self registrering of agents
		v1.POST("/clusters/register", ctrlApikeys.CreateForAgent())
		infoRoute := v1.Group("/info")
		{
			infoRoute.GET("/version", infocontroller.GetVersion())
		}
		m2mRoute := v1.Group("/m2m")
		{
			// Move along Nothing to see here
			m2mRoute.GET("/", ctrlM2mEaster.RegisterM2m())
			m2mRoute.POST("/heartbeat", notinuse.NotInUse())
			// Allow anonymous POST requests
		}

		v1.Use(auth.AuthenticationMiddleware)
		aclRoute := v1.Group("/acl")
		{
			aclRoute.POST("", ctrlAcl.Create())
			aclRoute.PUT("/:id", ctrlAcl.Update())
			aclRoute.DELETE("/:id", ctrlAcl.Delete())
			aclRoute.GET("/:id", ctrlAcl.GetById())

			aclRoute.HEAD("/:scope/:subject/:access", ctrlAcl.CheckAcl())

			aclRoute.HEAD("/access/:scope/:subject/:access", ctrlAcl.CheckAcl())
			//			aclRoute.GET("/access/:scope/:subject/", ctrlAcl.CheckAcl()) // /api/acl/cluster/sdi-ror-dev-32342
			//			aclRoute.GET("/access/:scope/", ctrlAcl.CheckAcl())          // /api/acl/cluster
			aclRoute.POST("/filter", ctrlAcl.GetByFilter())
			aclRoute.GET("/migrate", ctrlAcl.MigrateAcls())
			aclRoute.GET("/scopes", ctrlAcl.GetScopes())
		}

		apikeysRoute := v1.Group("apikeys")
		{
			apikeysRoute.POST("/filter", ctrlApikeys.GetByFilter())
			apikeysRoute.DELETE("/:id", ctrlApikeys.Delete())
			apikeysRoute.POST("", ctrlApikeys.CreateApikey())
		}

		auditlogsRoute := v1.Group("auditlogs")
		{
			auditlogsRoute.GET("/:id", ctrlAuditlogs.GetById())
			auditlogsRoute.POST("/filter", ctrlAuditlogs.GetByFilter())
			auditlogsRoute.GET("/metadata", ctrlAuditlogs.GetMetadata())
		}

		clusterRoute := v1.Group("cluster")
		{
			clusterRoute.GET("/:clusterid", ctrlClusters.ClusterGetById())
			clusterRoute.GET("/:clusterid/exists", ctrlClusters.ClusterExistsById())
			clusterRoute.POST("/:clusterid/heartbeat", ctrlClusters.RegisterHeartbeat())
			clusterRoute.PATCH("/:clusterid/metadata", ctrlClusters.UpdateMetadata())
			clusterRoute.POST("/heartbeat", ctrlClusters.RegisterHeartbeat())
		}

		clustersRoute := v1.Group("clusters")
		{
			clustersRoute.GET("/:clusterid", ctrlClusters.ClusterGetById())
			clustersRoute.GET("/:clusterid/exists", ctrlClusters.ClusterExistsById())
			clustersRoute.PATCH("/:clusterid/metadata", ctrlClusters.UpdateMetadata())

			clustersRoute.GET("/:clusterid/views/policyreports", ctrlClusters.PolicyreportsView())
			clustersRoute.GET("/:clusterid/views/vulnerabilityreports", ctrlClusters.VulnerabilityReportsView())
			clustersRoute.GET("/:clusterid/views/compliancereports", ctrlClusters.ComplianceReports())
			clustersRoute.GET("/:clusterid/views/ingresses", ctrlClusters.DummyView())
			clustersRoute.GET("/:clusterid/views/nodes", ctrlClusters.DummyView())
			clustersRoute.GET("/:clusterid/views/applications", ctrlClusters.DummyView())
			clustersRoute.GET("/:clusterid/configs/:name", ctrlM2mConfiguration.GetTaskConfiguration())

			clustersRoute.GET("/views/policyreports", ctrlClusters.PolicyreportSummaryView())
			clustersRoute.GET("/views/vulnerabilityreports/byid/:cveid", ctrlClusters.VulnerabilityReportsViewById())
			clustersRoute.GET("/views/vulnerabilityreports/byid", ctrlClusters.GlobalVulnerabilityReportsViewById())
			clustersRoute.GET("/views/vulnerabilityreports", ctrlClusters.VulnerabilityReportsGlobal())
			clustersRoute.GET("/views/compliancereports", ctrlClusters.ComplianceReportsGlobal())
			clustersRoute.POST("/filter", ctrlClusters.ClusterByFilter())
			clustersRoute.POST("/heartbeat", ctrlClusters.RegisterHeartbeat())
			clustersRoute.GET("/metadata", ctrlClusters.GetMetadata())

			clustersRoute.GET("/views/errorlist", ctrlClusters.DummyView())
			clustersRoute.GET("/views/clusterlist", ctrlClusters.DummyView())

			clustersRoute.GET("/self", ctrlClusters.GetSelf())

			clustersRoute.POST("/workspace/:workspaceName/filter", ctrlClusters.ClusterGetByWorkspace())
			clustersRoute.GET("/controlplanesMetadata", ctrlClusters.GetControlPlanesMetadata())

			clustersRoute.POST("", ctrlClusters.CreateCluster())
		}

		configsRoute := v1.Group("configs")
		{
			configsRoute.GET("operator", ctrlM2mConfiguration.GetOperatorConfiguration())
		}

		datacentersRoute := v1.Group("datacenters")
		{
			datacentersRoute.GET("", ctrlDatacenters.GetAll())
			datacentersRoute.GET("/:datacenterName", ctrlDatacenters.GetByName())
			datacentersRoute.GET("/id/:id", ctrlDatacenters.GetById())
			datacentersRoute.POST("", ctrlDatacenters.Create())
			datacentersRoute.PUT("/:datacenterId", ctrlDatacenters.Update())
		}

		desiredVersionsRoute := v1.Group("/desired_versions")
		{
			desiredVersionsRoute.GET("", ctrlDesiredVersion.GetAll())
			desiredVersionsRoute.GET("/:key", ctrlDesiredVersion.GetByKey())
			desiredVersionsRoute.POST("", ctrlDesiredVersion.Create())
			desiredVersionsRoute.PUT("/:key", ctrlDesiredVersion.Update())
			desiredVersionsRoute.DELETE("/:key", ctrlDesiredVersion.Delete())
		}

		ordersRoute := v1.Group("orders")
		{
			ordersRoute.POST("/cluster", ctrlOrder.OrderCluster())
			ordersRoute.DELETE("/cluster", ctrlOrder.DeleteCluster())
			ordersRoute.GET("", ctrlOrder.GetOrders())
			ordersRoute.GET("/:uid", ctrlOrder.GetOrder())
			ordersRoute.DELETE("/:uid", ctrlOrder.DeleteOrder())
		}

		metricsRoute := v1.Group("metrics")
		{
			metricsRoute.GET("", ctrlMetrics.GetTotalByUser())
			metricsRoute.POST("", ctrlMetrics.RegisterResourceMetricsReport())

			metricsRoute.GET("/datacenters", ctrlMetrics.GetForDatacenters())
			metricsRoute.GET("/datacenter/:datacenterName", ctrlMetrics.GetByDatacenterName())

			metricsRoute.GET("/clusters", ctrlMetrics.GetForClusters())
			metricsRoute.GET("/clusters/workspace/:workspaceName", ctrlMetrics.GetForClustersByWorkspace())
			metricsRoute.GET("/cluster/:clusterId", ctrlMetrics.GetByClusterId())

			metricsRoute.GET("/custom/cluster/:property", ctrlMetrics.MetricsForClustersByProperty())

			metricsRoute.GET("/total", ctrlMetrics.GetTotal())

			metricsRoute.GET("/workspace/:workspaceName", ctrlMetrics.GetByWorkspaceName())
			metricsRoute.POST("/workspaces/filter", ctrlMetrics.GetForWorkspaces())
			metricsRoute.POST("/workspaces/datacenter/:datacenterName/filter", ctrlMetrics.GetForWorkspacesByDatacenter())
		}

		operatorconfigRoute := v1.Group("operatorconfigs")
		{
			operatorconfigRoute.GET("", ctrlOperatorConfigs.GetAll())
			operatorconfigRoute.GET("/:id", ctrlOperatorConfigs.GetById())
			operatorconfigRoute.POST("", ctrlOperatorConfigs.Create())
			operatorconfigRoute.PUT("/:id", ctrlOperatorConfigs.Update())
			operatorconfigRoute.DELETE("/:id", ctrlOperatorConfigs.Delete())
		}

		providerRouter := v1.Group("providers")
		{
			providerRouter.GET("", ctrlProviders.GetAll())
			providerRouter.GET("/:providerType/kubernetes/versions", ctrlProviders.GetKubernetesVersionByProvider())
		}

		pricesRoute := v1.Group("prices")
		{
			pricesRoute.GET("", ctrlPrices.GetAll())

			pricesRoute.GET("/:priceId", ctrlPrices.GetById())
			pricesRoute.POST("", ctrlPrices.Create(), middlewares.AuditLogMiddleware("Price created", models.AuditCategoryPrice, models.AuditActionCreate))
			pricesRoute.PUT("/:priceId", ctrlPrices.Update(), middlewares.AuditLogMiddleware("Price updated", models.AuditCategoryPrice, models.AuditActionUpdate))
			// pricesRoute.DELETE(":id", ctrlPrices.Delete(), middlewares.AuditLogMiddleware("Price deleted", models.Price.String(), models.DELETE.String()))

			pricesRoute.GET("/provider/:providerName", ctrlPrices.GetByProvider())
		}

		projectsRoute := v1.Group("projects")
		{
			projectsRoute.GET("/:id", ctrlProjects.GetById())
			projectsRoute.GET("/:id/clusters", ctrlProjects.GetClustersByProjectId())

			projectsRoute.POST("/filter", ctrlProjects.GetByFilter())

			projectsRoute.POST("", ctrlProjects.Create())
			projectsRoute.PUT("/:id", ctrlProjects.Update())
			projectsRoute.DELETE(":id", ctrlProjects.Delete())
		}

		resourceRoute := v1.Group("resources")
		{
			resourceRoute.GET("", resourcescontroller.GetResources())
			resourceRoute.POST("", resourcescontroller.NewResource())
			resourceRoute.GET("/uid/:uid", resourcescontroller.GetResource())
			resourceRoute.PUT("/uid/:uid", resourcescontroller.UpdateResource())
			resourceRoute.DELETE("/uid/:uid", resourcescontroller.DeleteResource())
			resourceRoute.HEAD("/uid/:uid", resourcescontroller.ExistsResources())

			resourceRoute.GET("/hashes", resourcescontroller.GetResourceHashList())
		}

		usersRoute := v1.Group("users")
		{
			selfRoute := usersRoute.Group("self")
			selfRoute.GET("", ctrlUsers.GetUser())
			selfRoute.POST("/apikeys", ctrlUsers.CreateApikey())
			selfRoute.POST("/apikeys/filter", ctrlUsers.GetApiKeysByFilter())
			selfRoute.DELETE("/apikeys/:id", ctrlUsers.DeleteApiKey())
		}

		tasksRoute := v1.Group("tasks")
		{
			tasksRoute.GET("", ctrlTasks.GetAll())
			tasksRoute.GET("/:id", ctrlTasks.GetById())
			tasksRoute.POST("", ctrlTasks.Create())
			tasksRoute.PUT("/:id", ctrlTasks.Update())
			tasksRoute.DELETE("", ctrlTasks.Delete())
		}

		workspacesRoute := v1.Group("workspaces")
		{
			workspacesRoute.GET("", ctrlWorkspaces.GetAll())
			workspacesRoute.GET("/:workspaceName", ctrlWorkspaces.GetByName())
			workspacesRoute.GET("/id/:id", ctrlWorkspaces.GetById())
			workspacesRoute.PUT("/:id", ctrlWorkspaces.Update())
			workspacesRoute.POST("/:workspaceName/login", ctrlWorkspaces.GetKubeconfig())
		}

		rulesetsRoute := v1.Group("rulesetsController")
		{
			if viper.GetBool(configconsts.DEVELOPMENT) {
				rulesetsRoute.GET("", ctrlRulesets.GetAll())
			}

			rulesetsRoute.GET("/cluster/:clusterId", ctrlRulesets.GetByCluster())
			rulesetsRoute.GET("/internal", ctrlRulesets.GetInternal())

			rulesetsRoute.PUT("/:rulesetId/resources", ctrlRulesets.AddResource())

			rulesetsRoute.DELETE("/:rulesetId/resources/:resourceId", ctrlRulesets.DeleteResource())

			rulesetsRoute.POST("/:rulesetId/resources/:resourceId/rules", ctrlRulesets.AddResourceRule())
			rulesetsRoute.DELETE("/:rulesetId/resources/:resourceId/rules/:ruleId", ctrlRulesets.DeleteResourceRule())

		}
	}

	v2 := router.Group("/v2")

	v2.Use(middlewares.TimeoutMiddleware(timeoutduration))
	v2.Use(auth.AuthenticationMiddleware)
	selfv2Route := v2.Group("self")
	selfv2Route.GET("", handlerv2self.GetSelf())
	selfv2Route.POST("/apikeys", handlerv2self.CreateOrRenewApikey())
	selfv2Route.DELETE("/apikeys/:id", handlerv2self.DeleteApiKey())

	router.GET("/health", ctrlHealth.GetHealthStatus())
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true})))

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v2.Use(auth.AuthenticationMiddleware)

	resourceRoute := v2.Group("resources")

	resourceRoute.GET("", v2resourcescontroller.GetResources())
	resourceRoute.POST("", v2resourcescontroller.NewResource())
	resourceRoute.GET("/uid/:uid", v2resourcescontroller.GetResource())
	resourceRoute.PUT("/uid/:uid", v2resourcescontroller.UpdateResource())
	resourceRoute.DELETE("/uid/:uid", v2resourcescontroller.DeleteResource())
	resourceRoute.HEAD("/uid/:uid", v2resourcescontroller.ExistsResources())
	resourceRoute.GET("/hashes", v2resourcescontroller.GetResourceHashList())
}
