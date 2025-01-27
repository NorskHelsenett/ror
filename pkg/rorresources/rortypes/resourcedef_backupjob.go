package rortypes

import "time"

type ResourceBackupJob struct {
	Id       string                  `json:"id"`
	Provider string                  `json:"provider"`
	Status   ResourceBackupJobStatus `json:"status"`
	Spec     ResourceBackupJobSpec   `json:"spec"`
}

// The requested parameters about a job
type ResourceBackupJobSpec struct {

	// The name of the job
	Name string `json:"name"`

	// Status of the backup job, active, paused, etc.
	Status string `json:"status"`

	// Defines the name of the system the run originates from
	SourceName string `json:"sourceName"`

	// Defines the id of the system the run originates from
	SourceId string `json:"sourceId"`

	// Defines the policy if applicable at the local system
	// If policies are not used these can be left as blank
	PolicyId   string `json:"policyId"`
	PolicyName string `json:"policyName"`

	// Direct targets for this backup job
	ActiveTargets []ResourceBackupTarget `json:"activeTargets"`

	// Indirect targets for this backup job
	IndirectBackupTargets []ResourceIndirectBackupTarget `json:"indirectBackupTargets"`

	// Any destination defined by this backup job
	BackupDestinations []ResourceBackupDestination `json:"backupDestinations"`

	// Some backup systems allow StartTime to be defined per backupJob, while some use policies
	StartTime time.Time `json:"startTime"`

	// Some backup systems allow EndTime to be defined per backupJob, while some use policies
	EndTime time.Time `json:"endTime"`

	// Some backup systems allow ExpiryTime to be defined per backupJob, while some use policies
	ExpiryTime time.Time `json:"expiryTime"`
}

// The observed parameters about a job
type ResourceBackupJobStatus struct {
	ResourceBackupJobSpec
	Runs []ResourceBackupRun `json:"runs"`
}

// Once instance of a run from a backup job
type ResourceBackupRun struct {
	BackupTargets      []ResourceBackupTarget      `json:"backupTargets"`
	BackupDestinations []ResourceBackupDestination `json:"backupDestinations"`
	StartTime          time.Time                   `json:"startTime"`
	EndTime            time.Time                   `json:"endTime"`
	ExpiryTime         time.Time                   `json:"expiryTime"`
	BackupStorage      ResourceBackupStorage       `json:"backupStorage"`
}

// Defines a singular direct backup target, this could be a VM, a storage object, etc.
type ResourceDirectBackupTarget struct {
	BackupTargets     []ResourceBackupTarget    `json:"backupTargets"`
	BackupDestination ResourceBackupDestination `json:"backupDestination"`
	StartTime         time.Time                 `json:"startTime"`
	EndTime           time.Time                 `json:"endTime"`
	ExpiryTime        time.Time                 `json:"expiryTime"`
	BackupStorage     ResourceBackupStorage     `json:"backupStorage"`
}

// Defines a singular backup target, this could be a VM, a storage object, etc.
type ResourceBackupTarget struct {
	Name        string            `json:"name"`        // Defines the object's name
	Id          string            `json:"id"`          // Defines the object's id
	ExternalIds map[string]string `json:"externalIds"` // Defines any external ids by the backup system(s)
}

// Defines a indrect backup target, which can result into multiple objects (For example a tag or multiple tags should result into being in a backup job)
// One instance indicates all should match
// Multiple instances would indicate different matching groups
type ResourceIndirectBackupTarget struct {
	Type      string              `json:"type"`      // The type of indrect target, some allow tags and/or other types of criteria
	Ids       []string            `json:"ids"`       // For where the ids are referenced
	KeyValues map[string][]string `json:"keyValues"` // For key values pairs, some allow the same key with different values
}

type ResourceBackupDestination struct {
	Name string `json:"name"`
	Id   string `json:"id"`

	// Local, remote, archive, etc.
	Type string `json:"type"`

	// Status spesific to the destination - remote being unavailable
	Status     string    `json:"status"`
	ExpiryTime time.Time `json:"expiryTime"` // ExpiryTime is defined per destination
}

type ResourceBackupStorage struct {
	Unit         string `json:"unit"`
	SourceSize   int    `json:"sourceSize"`   // Total changed data in the run, incremental will have changes since last time, full runs will have the entire VM - not including unusued space
	LogicalSize  int    `json:"logicalSize"`  // The total logical size of the VM
	PhysicalSize int    `json:"physicalSize"` // Physical data written to the backup system
}
