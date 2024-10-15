package auditlog

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/models"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"
	auditlogrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/auditlogRepo"
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
)

// Create creates a new auditlog entry in the database
func Create(ctx context.Context, msg string, category models.AuditCategory, action models.AuditAction, user *identitymodels.User, newObject any, oldObject any) (*mongoTypes.MongoAuditLog, error) {
	auditLog := mongoTypes.MongoAuditLog{}
	auditLogMetadata := mongoTypes.MongoAuditLogMetadata{}
	auditLogMetadata.Msg = msg
	auditLogMetadata.Timestamp = time.Now()
	auditLogMetadata.Category = category
	auditLogMetadata.Action = action
	auditLogMetadata.User = *user
	auditLog.Metadata = auditLogMetadata
	data := make(map[string]any)
	data["new_object"] = newObject
	data["old_object"] = oldObject
	auditLog.Data = data
	result, err := auditlogrepo.Create(ctx, auditLog)
	if err != nil {
		return nil, fmt.Errorf("could not create auditlog: %v", err)
	}
	return result, nil
}
