package rortypes

import "time"

type ResourceBackupJob struct {
<<<<<<< HEAD
	Status ResourceBackupJobStatus `json:"status"`
	Spec   ResourceBackupJobSpec   `json:"spec"`
}

// The requested parameters about a job
type ResourceBackupJobSpec struct {

	// The name of the job
	Name string `json:"name"`

	// Defines the name of the system the run originates from
	SourceName string `json:"sourceName"`

	// Defines the id of the system the run originates from
	SourceId string `json:"sourceId"`

	// Defines the policy id at the local system that defines the rules for the data, how long it's stored
	// where's it's stored, and other options
	PolicyId             string                         `json:"policyId"`
	TargetObjectsRefs    []string                       `json:"targetObjectsRefs"`
	DirectBackupTarget   []ResourceIndirectBackupTarget `json:"directBackupTarget"`
	IndirectBackupTarget []ResourceDirectBackupTarget   `json:"indirectBackupTarget"`
	BackupDestination    []ResourceBackupDestination    `json:"backupDestination"`

	// Some backup systems allow StartTime to be defined per backupJob, while some use policies
	StartTime time.Time `json:"startTime"`

	// Some backup systems allow EndTime to be defined per backupJob, while some use policies
	EndTime time.Time `json:"endTime"`

	// Some backup systems allow ExpiryTime to be defined per backupJob, while some use policies
	ExpiryTime time.Time `json:"expiryTime"`
}

// The observed parameters about a job
type ResourceBackupJobStatus struct {

	// The name of the job
	Name string `json:"name"`

	// Defines the name of the system the run originates from
	SourceName string `json:"sourceName"`

	// Defines the id of the system the run originates from
	SourceId string `json:"sourceId"`

	// Defines the id of the job from the system in originates from
	JobId string `json:"jobId"`

	// Defines the policy id at the local system that defines the rules for the data, how long it's stored
	// where's it's stored, and other options
	PolicyId             string                         `json:"policyId"`
	DirectBackupTarget   []ResourceIndirectBackupTarget `json:"DirectBackupTarget"`
	IndirectBackupTarget []ResourceDirectBackupTarget   `json:"backupTarget"`
	ResourceBackupRuns   []ResourceBackupRun            `json:"resourceBackupRuns"`

	// Some backup systems allow StartTime to be defined per backupJob, while some use policies
	StartTime time.Time `json:"startTime"`

	// Some backup systems allow EndTime to be defined per backupJob, while some use policies
	EndTime time.Time `json:"endTime"`

	// Some backup systems allow ExpiryTime to be defined per backupJob, while some use policies
	ExpiryTime time.Time `json:"expiryTime"`
=======
	SourceName            string                `json:"sourceName"` // Defines the name of the system the run originates from
	SourceId              string                `json:"sourceId"`   // Defines the id of the system the run originates from
	JobId                 string                `json:"jobId"`      // Defines the id of the job from the system in originates from
	ResourceBackupJobSpec ResourceBackupJobSpec `json:"resourceBackupJobSpec"`
	ResourceBackupRuns    []ResourceBackupRun   `json:"resourceBackupRuns"`
}

type ResourceBackupJobSpec struct {
	Name              string   `json:"name"` // The name of the job
	TargetObjectsRefs []string `json:"targetObjectsRefs"`
>>>>>>> 550cbdd (Added resourceBackupJob types and generator stuff)
}

// Once instance of a run from a backup job
type ResourceBackupRun struct {
<<<<<<< HEAD
	BackupTarget      []ResourceDirectBackupTarget `json:"backupTarget"`
	BackupDestination []ResourceBackupDestination  `json:"backupDestination"`
	StartTime         time.Time                    `json:"startTime"`
	EndTime           time.Time                    `json:"endTime"`
	ExpiryTime        time.Time                    `json:"expiryTime"`
	BackupStorage     ResourceBackupStorage        `json:"backupStorage"`
}

// Defines a singular direct backup target, this could be a VM, a storage object, etc.
type ResourceDirectBackupTarget struct {
=======
	BackupTarget      []ResourceBackupTarget      `json:"backupTarget"`
	BackupDestination []ResourceBackupDestination `json:"backupDestination"` // A run can have multiple targets defined into a single job, multiple jobs would result in multiple runs
	StartTime         time.Time                   `json:"startTime"`
	EndTime           time.Time                   `json:"endTime"`
	ExpiryTime        time.Time                   `json:"expiryTime"`
	BackupStorage     ResourceBackupStorage       `json:"backupStorage"`
}

// Defines a singular backup target, this could be a VM, a storage object, etc.
type ResourceBackupTarget struct {
>>>>>>> 550cbdd (Added resourceBackupJob types and generator stuff)
	Name        string            `json:"name"`        // Defines the object's name
	Id          string            `json:"id"`          // Defines the object's id
	ExternalIds map[string]string `json:"externalIds"` // Defines any external ids by the backup system(s)
}

<<<<<<< HEAD
// Defines a indrect backup target, which can result into multiple objects (For example a tag or multiple tags should result into being in a backup job)
// One instance indicates all should match
// Multiple instances would indicate different matching groups
type ResourceIndirectBackupTarget struct {
	Type      string              `json:"type"`      // The type of indrect target, some allow tags and/or other types of criteria
	Ids       []string            `json:"ids"`       // For where the ids are referenced
	KeyValues map[string][]string `json:"keyValues"` // For key values pairs, some allow the same key with different values
}
=======
// the ExpiryTime is defined per target
>>>>>>> 550cbdd (Added resourceBackupJob types and generator stuff)

type ResourceBackupDestination struct {
	Name       string    `json:"name"`
	Id         string    `json:"id"`
<<<<<<< HEAD
	Type       string    `json:"type"`       // Local, remote, archive, etc.
	Success    bool      `json:"success"`    // Whether the run was successful in creating a copy, despite any possible issues
	ExpiryTime time.Time `json:"expiryTime"` // ExpiryTime is defined per destination
=======
	Type       string    `json:"type"`    // Local, remote, archive, etc.
	Success    bool      `json:"success"` // Whether the run was successful in creating a copy, despite any possible issues
	ExpiryTime time.Time `json:"expiryTime"`
>>>>>>> 550cbdd (Added resourceBackupJob types and generator stuff)
}

type ResourceBackupStorage struct {
	Unit         string `json:"unit"`
	SourceSize   int    `json:"sourceSize"`   // Total changed data in the run, incremental will have changes since last time, full runs will have the entire VM - not including unusued space
	LogicalSize  int    `json:"logicalSize"`  // The total logical size of the VM
	PhysicalSize int    `json:"physicalSize"` // Physical data written to the backup system
}
