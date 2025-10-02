package rortypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceBackupJob struct {
	Id       string                  `json:"id"`
	Provider string                  `json:"provider"`
	Source   string                  `json:"source"`
	Status   ResourceBackupJobStatus `json:"status"`
	Spec     ResourceBackupJobSpec   `json:"spec"`
}

// The requested parameters about a job
type ResourceBackupJobSpec struct {

	// The name of the job
	Name string `json:"name"`

	// Status of the backup job, active, paused, etc.
	Status string `json:"status"`

	// Defines the policy if applicable at the local system
	// If policies are not used these can be left as blank
	PolicyId string `json:"policyId"`

	Schedules []ResourceBackupSchedule `json:"schedules"`

	// Direct targets for this backup job
	ActiveTargets []ResourceBackupTarget `json:"activeTargets"`

	// Indirect targets for this backup job
	IndirectBackupTargets []ResourceIndirectBackupTarget `json:"indirectBackupTargets"`

	// Any destination defined by this backup job
	BackupDestinations []ResourceBackupDestination `json:"backupDestinations"`
}

// The observed parameters about a job
type ResourceBackupJobStatus struct {
	ResourceBackupJobSpec `json:"resourceBackupJobSpec"`
	Location              string      `json:"location"`
	LastUpdated           metav1.Time `json:"lastUpdated"`
	PolicyName            string      `json:"policyName"`

	// Any runs connected to this backup job
	BackupRunIds []string `json:"backupRunIds"`
}

// Defines a singular backup target, this could be a VM, a storage object, etc.
type ResourceBackupTarget struct {
	// Defines the object's name
	Name string `json:"name"`

	// Defines the object's id
	Id string `json:"id"`

	// Defines any external id
	ExternalId string `json:"externalId"`

	// Defines the source of this object to the backup system
	Source *ResourceBackupSource `json:"source,omitempty"`

	// Defines the size of the snapshots of the target
	Size *ResourceBackupStorage `json:"size,omitempty"`
}

// The backup target's source, a vCenter, a NetApp system, etc.
type ResourceBackupSource struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Uuid string `json:"uuid"`
	Type string `json:"type"`
}

// Defines a indirect backup target, which can result into multiple objects (For example a tag or multiple tags should result into being in a backup job)
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
	Status string `json:"status"`
}

type BackupScheduleType string

const (
	BackupScheduleTypeLocal   = "local"
	BackupScheduleTypeReplica = "replica"
	BackupScheduleTypeArchive = "archive"
)

type ResourceBackupSchedule struct {
	Type BackupScheduleType `json:"type"`

	// When will the job start
	StartTime string `json:"startTime"`

	// When will the job be forcibly stopped, if empty it will continue indefinitely
	EndTime string `json:"endTime"`

	//  How many time per unit will this backup run
	Frequency int `json:"frequency"`

	// What unit of time is this schedule going to run in
	Unit      string                          `json:"unit"`
	Retention ResourceBackupScheduleRetention `json:"retention"`
}

type ResourceBackupScheduleRetention struct {
	Duration int    `json:"duration"`
	Unit     string `json:"unit"`
}
