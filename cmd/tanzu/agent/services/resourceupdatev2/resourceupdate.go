package resourceupdatev2

import (
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/services/authservice"
	"github.com/NorskHelsenett/ror/internal/models/rorResources"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func SendResource(action apiresourcecontracts.ResourceAction, input *unstructured.Unstructured) {
	resource, err := rorResources.NewFromUnstructured(input)
	resourceReturn := resource.NewResourceUpdateModel(authservice.CreateOwnerref(), action)
	if err != nil {
		return
	}

	if action != apiresourcecontracts.K8sActionDelete {
		if ResourceCache.CleanupRunning() {
			ResourceCache.MarkActive(resourceReturn.Uid)
		}
		needUpdate := ResourceCache.HashList.checkUpdateNeeded(resourceReturn.Uid, resourceReturn.Hash)
		if needUpdate {
			err = sendResourceUpdateToRor(resourceReturn)
			if err != nil {
				rlog.Error("error sending resource update to ror, added to retryque", err)
				ResourceCache.Workqueue.Add(resourceReturn)
				return
			}
			ResourceCache.HashList.updateHash(resourceReturn.Uid, resourceReturn.Hash)
		}
	} else if action == apiresourcecontracts.K8sActionDelete {
		err := sendResourceUpdateToRor(resourceReturn)
		if err != nil {
			rlog.Error("error sending resource update to ror, added to retryque", err)
			ResourceCache.Workqueue.Add(resourceReturn)
			return
		}
	}

}
