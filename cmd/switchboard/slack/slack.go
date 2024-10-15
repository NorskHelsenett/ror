package slack

import (
	"github.com/NorskHelsenett/ror/cmd/switchboard/ror"

	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/types"
)

const (
	apiVersion = "general.ror.internal/v1alpha1"
	kind       = "SlackMessage"
)

func CreateSlackMessage(channel, message string, owner rortypes.RorResourceOwnerReference) error {
	sm := &rortypes.ResourceSlackMessage{}
	sm.Spec.ChannelId = channel
	sm.Spec.Message = message

	r := rorresources.NewRorResource(kind, apiVersion)
	r.Metadata.UID = types.UID(uuid.NewString())
	r.RorMeta.Action = rortypes.K8sActionUpdate
	r.RorMeta.Ownerref = owner
	r.SetSlackMessage(sm)

	rs := rorresources.NewResourceSet()
	rs.Add(r)

	_, err := ror.Client.ResourceV2().Update(rs)
	if err != nil {
		return err
	}
	return nil
}
