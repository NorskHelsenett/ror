package models

type AuditAction string

const (
	AuditActionUnknown AuditAction = "Unknown"
	AuditActionCreate  AuditAction = "Create"
	AuditActionUpdate  AuditAction = "Update"
	AuditActionDelete  AuditAction = "Delete"
	AuditActionRead    AuditAction = "Read"
)
