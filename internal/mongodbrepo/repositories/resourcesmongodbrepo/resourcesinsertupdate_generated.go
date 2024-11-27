// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package resourcesmongodbrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
)

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]
func CreateResourceNamespace(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode]
func CreateResourceNode(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim]
func CreateResourcePersistentVolumeClaim(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment]
func CreateResourceDeployment(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass]
func CreateResourceStorageClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport]
func CreateResourcePolicyReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]
func CreateResourceApplication(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject]
func CreateResourceAppProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate]
func CreateResourceCertificate(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService]
func CreateResourceService(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod]
func CreateResourcePod(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet]
func CreateResourceReplicaSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet]
func CreateResourceStatefulSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet]
func CreateResourceDaemonSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress]
func CreateResourceIngress(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass]
func CreateResourceIngressClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]
func CreateResourceVulnerabilityReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport]
func CreateResourceExposedSecretReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport]
func CreateResourceConfigAuditReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport]
func CreateResourceRbacAssessmentReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster]
func CreateResourceTanzuKubernetesCluster(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease]
func CreateResourceTanzuKubernetesRelease(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass]
func CreateResourceVirtualMachineClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding]
func CreateResourceVirtualMachineClassBinding(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster]
func CreateResourceKubernetesCluster(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder]
func CreateResourceClusterOrder(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject]
func CreateResourceProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration]
func CreateResourceConfiguration(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport]
func CreateResourceClusterComplianceReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport]
func CreateResourceClusterVulnerabilityReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute]
func CreateResourceRoute(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage]
func CreateResourceSlackMessage(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent]
func CreateResourceVulnerabilityEvent(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine]
func CreateResourceVirtualMachine(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceEndpoints]
func CreateResourceEndpoints(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceEndpoints], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Creates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNetworkPolicy]
func CreateResourceNetworkPolicy(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNetworkPolicy], ctx context.Context) error {
	rlog.Debug("inserting resource",
		rlog.String("action", "insert"),
		rlog.String("apiverson", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)
	_, err := mongodb.InsertOne(ctx, ResourceCollectionName, input)
	if err != nil {
		msg := fmt.Sprintf("could not create resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace] by uid
func UpdateResourceNamespace(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode] by uid
func UpdateResourceNode(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim] by uid
func UpdateResourcePersistentVolumeClaim(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment] by uid
func UpdateResourceDeployment(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass] by uid
func UpdateResourceStorageClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport] by uid
func UpdateResourcePolicyReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication] by uid
func UpdateResourceApplication(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject] by uid
func UpdateResourceAppProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate] by uid
func UpdateResourceCertificate(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService] by uid
func UpdateResourceService(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod] by uid
func UpdateResourcePod(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet] by uid
func UpdateResourceReplicaSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet] by uid
func UpdateResourceStatefulSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet] by uid
func UpdateResourceDaemonSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress] by uid
func UpdateResourceIngress(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass] by uid
func UpdateResourceIngressClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport] by uid
func UpdateResourceVulnerabilityReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport] by uid
func UpdateResourceExposedSecretReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport] by uid
func UpdateResourceConfigAuditReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport] by uid
func UpdateResourceRbacAssessmentReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster] by uid
func UpdateResourceTanzuKubernetesCluster(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease] by uid
func UpdateResourceTanzuKubernetesRelease(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass] by uid
func UpdateResourceVirtualMachineClass(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding] by uid
func UpdateResourceVirtualMachineClassBinding(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster] by uid
func UpdateResourceKubernetesCluster(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder] by uid
func UpdateResourceClusterOrder(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject] by uid
func UpdateResourceProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration] by uid
func UpdateResourceConfiguration(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport] by uid
func UpdateResourceClusterComplianceReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport] by uid
func UpdateResourceClusterVulnerabilityReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute] by uid
func UpdateResourceRoute(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage] by uid
func UpdateResourceSlackMessage(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent] by uid
func UpdateResourceVulnerabilityEvent(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine] by uid
func UpdateResourceVirtualMachine(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceEndpoints] by uid
func UpdateResourceEndpoints(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceEndpoints], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}

// Updates resource entry of type apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNetworkPolicy] by uid
func UpdateResourceNetworkPolicy(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNetworkPolicy], ctx context.Context) error {
	rlog.Debug("updating resource",
		rlog.String("action", "update"),
		rlog.String("api version", input.ApiVersion),
		rlog.String("kind", input.Kind),
		rlog.String("uid", input.Uid),
	)

	filter := bson.M{"uid": input.Uid}
	update := bson.M{"$set": input}
	_, err := mongodb.UpdateOne(ctx, ResourceCollectionName, filter, update)
	if err != nil {
		msg := fmt.Sprintf("could not update resource %s/%s with uid %s", input.ApiVersion, input.Kind, input.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	return nil
}
