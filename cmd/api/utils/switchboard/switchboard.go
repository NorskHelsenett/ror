package switchboard

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/messages"
)

func PublishStarted(ctx context.Context) {
	// hostname, _ := os.Hostname()

	// message := messagebuscontracts.SwitchboardPost{
	// 	Source: messagebuscontracts.SwitchboardSource{
	// 		Type:      messagebuscontracts.SwitchboardSourceTypeRor,
	// 		ClusterId: "ror",
	// 		Uid:       "ror-api",
	// 	},
	// 	Event:    messages.RulesetRuleTypeStarted,
	// 	Severity: "info",
	// 	Attributes: map[string]string{
	// 		"hostname": hostname,
	// 	},
	// }

	//err := apiconfig.RabbitMQConnection.SendMessage(ctx, message, messagebuscontracts.Workqueue_Switchboard_Post)
	// if err != nil {
	// 	rlog.Errorc(ctx, "could not publish started message", err)
	// }
}

func PublishResourceToSwitchboard(ctx context.Context, rule messages.RulesetRuleType, input apiresourcecontracts.ResourceUpdateModel) error {
	// post := messagebuscontracts.SwitchboardPost{
	// 	Source: messagebuscontracts.SwitchboardSource{
	// 		Type:      messagebuscontracts.SwitchboardSourceTypeCluster,
	// 		ClusterId: input.Owner.Subject,

	// 		Uid:        input.Uid,
	// 		ApiVersion: input.ApiVersion,
	// 		Kind:       input.Kind,
	// 		Namespace:  "default",
	// 	},
	// 	Severity: messagebuscontracts.SwitchboardPostSeverityInfo,
	// 	Event:    rule,
	// }

	//err := apiconfig.RabbitMQConnection.SendMessage(ctx, post, messagebuscontracts.Workqueue_Switchboard_Post)
	return nil //err
}
