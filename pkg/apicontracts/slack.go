package messagebuscontracts

type SlackMessageCreate struct {
	EventBase
	Message   string `json:"message"`
	ChannelId string `json:"channelId"`
}
