package models

import "time"

type InstalledState struct {
	AppLog []App `json:"appLog"`
}

type App struct {
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	ConfigMd5 string    `json:"configMd5"`
}

type ApplicationFailLog struct {
	Failures []Failure `json:"failures"`
}

type Failure struct {
	App   App
	Count int32 `json:"count"`
}

type TaskStatus struct {
	ConfigChanged bool `json:"configChanged"`
	Installed     bool `json:"installed"`
}
