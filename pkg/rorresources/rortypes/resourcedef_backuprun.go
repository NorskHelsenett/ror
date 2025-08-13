package rortypes

import "time"

type ResourceBackupRun struct {
	Id       string                  `json:"id"`
	Provider string                  `json:"provider"`
	Source   string                  `json:"source"`
	Status   ResourceBackupRunStatus `json:"status"`
	Spec     ResourceBackupJobSpec   `json:"spec"`
}

// Once instance of a run from a backup job
type ResourceBackupRunStatus struct {
	Id                 string                      `json:"id"`
	BackupJobId        string                      `json:"backupJobId"`
	BackupTargets      []ResourceBackupTarget      `json:"backupTargets"`
	BackupDestinations []ResourceBackupDestination `json:"backupDestinations"`

	// When the run was started
	StartTime time.Time `json:"startTime"`

	// When the run was finished
	EndTime time.Time `json:"endTime"`

	// When the run will expire and be deleted
	ExpiryTime    time.Time             `json:"expiryTime"`
	BackupStorage ResourceBackupStorage `json:"backupStorage"`
	LastUpdated   time.Time             `json:"lastUpdated"`
}

// Storage used by an instance of a run or a target
type ResourceBackupStorage struct {
	// What unit are the sizes in
	Unit string `json:"unit"`

	// Total changed data in the run, incremental will have changes since last time, full runs will have the entire VM - not including unusued space
	SourceSize int `json:"sourceSize"`

	// The total logical size of the VM
	LogicalSize int `json:"logicalSize"`

	// Physical data written to the backup system
	PhysicalSize int `json:"physicalSize"`
}
