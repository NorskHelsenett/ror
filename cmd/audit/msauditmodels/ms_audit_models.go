package msauditmodels

type CommitPostMessage struct {
	Branch        string         `json:"branch"`
	CommitMessage string         `json:"commit_message"`
	Actions       []CommitAction `json:"actions"`
}

type CommitAction struct {
	Action   string `json:"action"`
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}
