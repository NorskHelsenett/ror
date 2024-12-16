package rortypes

import "time"

type ResourceBackupJob struct {
	SourceName            string                `json:"sourceName"` // Defines the name of the system the run originates from
	SourceId              string                `json:"sourceId"`   // Defines the id of the system the run originates from
	JobId                 string                `json:"jobId"`      // Defines the id of the job from the system in originates from
	ResourceBackupJobSpec ResourceBackupJobSpec `json:"resourceBackupJobSpec"`
	ResourceBackupRuns    []ResourceBackupRun   `json:"resourceBackupRuns"`
}

type ResourceBackupJobSpec struct {
	Name              string   `json:"name"` // The name of the job
	TargetObjectsRefs []string `json:"targetObjectsRefs"`
}

// Once instance of a run from a backup job
type ResourceBackupRun struct {
	BackupTarget      []ResourceBackupTarget      `json:"backupTarget"`
	BackupDestination []ResourceBackupDestination `json:"backupDestination"` // A run can have multiple targets defined into a single job, multiple jobs would result in multiple runs
	StartTime         time.Time                   `json:"startTime"`
	EndTime           time.Time                   `json:"endTime"`
	ExpiryTime        time.Time                   `json:"expiryTime"`
	BackupStorage     ResourceBackupStorage       `json:"backupStorage"`
}

// Defines a singular backup target, this could be a VM, a storage object, etc.
type ResourceBackupTarget struct {
	Name        string            `json:"name"`        // Defines the object's name
	Id          string            `json:"id"`          // Defines the object's id
	ExternalIds map[string]string `json:"externalIds"` // Defines any external ids by the backup system(s)
}

// the ExpiryTime is defined per target

type ResourceBackupDestination struct {
	Name       string    `json:"name"`
	Id         string    `json:"id"`
	Type       string    `json:"type"`    // Local, remote, archive, etc.
	Success    bool      `json:"success"` // Whether the run was successful in creating a copy, despite any possible issues
	ExpiryTime time.Time `json:"expiryTime"`
}

type ResourceBackupStorage struct {
	Unit         string `json:"unit"`
	SourceSize   int    `json:"sourceSize"`   // Total changed data in the run, incremental will have changes since last time, full runs will have the entire VM - not including unusued space
	LogicalSize  int    `json:"logicalSize"`  // The total logical size of the VM
	PhysicalSize int    `json:"physicalSize"` // Physical data written to the backup system
}
