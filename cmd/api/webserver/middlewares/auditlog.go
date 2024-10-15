package middlewares

import (
	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"

	"github.com/NorskHelsenett/ror/pkg/context/gincontext"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
)

func AuditLogMiddleware(msg string, category models.AuditCategory, action models.AuditAction) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := gincontext.GetUserFromGinContext(c)
		ctx := c.Request.Context()
		if err != nil {
			rlog.Errorc(ctx, "unable to get user from auditlog middleware", err)
		}
		c.Next()
		if c.Writer.Status() != 200 {
			return
		}
		newObject, _ := c.Get("newObject")
		oldObject, _ := c.Get("oldObject")
		_, err = auditlog.Create(ctx, msg, category, action, user, newObject, oldObject)
		if err != nil {
			rlog.Errorc(ctx, "could not create auditlog", err)
		}
	}
}
