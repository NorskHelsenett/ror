package resourceupdatev2

import (
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rorkubernetes"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func SendResource(action rortypes.ResourceAction, input *unstructured.Unstructured) {
	rorres := rorkubernetes.NewResourceFromDynamicClient(input)
	err := rorres.SetRorMeta(rortypes.ResourceRorMeta{
		Version:  "v2",
		Ownerref: clients.RorConfig.CreateOwnerref(),
		Action:   action,
	})
	if err != nil {
		rlog.Error("error setting rormeta", err)
		return
	}

	rorres.GenRorHash()

	if action != rortypes.K8sActionDelete && ResourceCache.CleanupRunning() {
		ResourceCache.MarkActive(rorres.GetUID())
	}

	needUpdate := ResourceCache.HashList.CheckUpdateNeeded(rorres.GetUID(), rorres.GetRorHash())
	if needUpdate {

		ResourceCache.WorkQueue.Add(rorres)
		// if err != nil {
		// 	rlog.Error("error sending resource update to ror, added to retryQeue", err)
		// 	ResourceCache.WorkQeueue.Add(rorres)
		// 	return
		// }

	}

}
