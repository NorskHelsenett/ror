package messagebuscontracts

const (
	Workqueue_Slack_Message_Create = "ror.slack.message.create"
	Workqueue_Switchboard_Post     = "ror.switchboard.post"

	Route_Acl_Update      = "acl.update"
	Route_Auth            = "auth"
	Route_Cluster_Created = "cluster.created"

	Route_ResourceCreated = "resource.created"
	Route_ResourceUpdated = "resource.updated"
	Route_ResourceDeleted = "resource.deleted"

	Route_ProviderTanzuClusterCreate = "provider.tanzu.cluster.create"
	Route_ProviderTanzuClusterModify = "provider.tanzu.cluster.modify"
	Route_ProviderTanzuClusterDelete = "provider.tanzu.cluster.delete"
	Route_ProviderTanzuOperatorOrder = "provider.tanzu.operator.order"

	Event_Broadcast           = "event.broadcast"
	Event_ClusterCreated      = "event.cluster.created"
	Event_ClusterOrderUpdated = "event.clusterorder.updated"

	ExchangeRor          = "ror"
	ExchangeRorResources = "ror.resources"
	ExchangeRorEvents    = "ror.events"
	ExchangeTanzu        = "tanzu"
	ExchangeK3d          = "k3d"
	ExchangeKind         = "kind"
)
