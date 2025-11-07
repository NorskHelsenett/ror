package rorerror

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

// Field is a type alias for rlog.Field, used for adding structured logging fields
// to error logs. This allows consistent field formatting across the application.
type Field = rlog.Field

// GinHandleErrorAndAbort is a convenience function that handles error checking,
// logging, and response in a single call. It's useful for early returns in handlers.
//
// If err is not nil:
//   - Creates a RorError from the error
//   - Logs the error with context and any provided fields
//   - Aborts the Gin request with appropriate JSON response
//   - Returns true
//
// If err is nil:
//   - Does nothing
//   - Returns false
//
// Parameters:
//   - c: Gin context for the current request
//   - status: HTTP status code to return if error is not nil
//   - err: Error to check and handle (can be nil)
//   - fields: Optional structured logging fields
//
// Returns:
//   - true if error was handled, false if err was nil
//
// Example:
//
//	func Handler(c *gin.Context) {
//	    data, err := fetchData()
//	    if rorerror.GinHandleErrorAndAbort(c, 500, err, rlog.String("source", "database")) {
//	        return // Error was handled
//	    }
//	    // Continue processing...
//	}
func GinHandleErrorAndAbort(c *gin.Context, status int, err error, fields ...Field) bool {
	if err != nil {
		rorerror := NewRorErrorFromError(status, err)
		fields = append(fields, rlog.Int("statuscode", status))
		rorerror.GinLogErrorAbort(c, fields...)
		return true
	}
	return false
}

// maskValue masks sensitive string values for logging purposes.
// It preserves the first 2 and last 2 characters while replacing the middle
// with asterisks. This helps with debugging while protecting sensitive data.
//
// Parameters:
//   - value: The string to mask (must be at least 4 characters)
//
// Returns:
//   - Masked string showing only first 2 and last 2 characters
//
// Example:
//
//	masked := maskValue("secret123key")
//	// Returns: "se********ey"
func maskValue(value string) string {
	if len(value) < 4 {
		// For short strings, mask all characters
		return strings.Repeat("*", len(value))
	}
	maskedKey := string(value[0:2]) + strings.Repeat("*", len(value)-4) + string(value[len(value)-2:])
	return maskedKey
}

// GinLogErrorJSON logs the error with context and returns the error as a JSON response.
// Unlike GinLogErrorAbort, this does not abort the request, allowing the handler
// to continue processing if needed.
//
// The JSON response contains:
//   - status: HTTP status code
//   - message: Error message
//
// Parameters:
//   - c: Gin context for the current request
//   - fields: Optional structured logging fields (e.g., rlog.String("key", "value"))
//
// Example:
//
//	func Handler(c *gin.Context) {
//	    err := rorerror.NewRorError(400, "validation failed")
//	    err.GinLogErrorJSON(c, rlog.String("field", "email"))
//	    // Handler can continue if needed
//	}
func (e ErrorData) GinLogErrorJSON(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.JSON(e.Status, e)
}

// GinLogErrorAbort logs the error with context and aborts the Gin request
// with a JSON response containing the error details. This is the most common
// way to handle errors in Gin handlers.
//
// The JSON response contains:
//   - status: HTTP status code
//   - message: Error message
//
// After calling this method:
//   - The request is aborted (no further handlers execute)
//   - The error is logged with full context
//   - A JSON error response is sent to the client
//
// Parameters:
//   - c: Gin context for the current request
//   - fields: Optional structured logging fields (e.g., rlog.String("key", "value"))
//
// Example:
//
//	func Handler(c *gin.Context) {
//	    if !isAuthorized {
//	        err := rorerror.NewRorError(403, "unauthorized")
//	        err.GinLogErrorAbort(c, rlog.String("user_id", userId))
//	        return
//	    }
//	}
func (e ErrorData) GinLogErrorAbort(c *gin.Context, fields ...Field) {
	e.logError(c, fields...)
	c.AbortWithStatusJSON(e.Status, e)
}

// logError is an internal method that handles the actual logging of errors.
// It performs the following:
//   - Adds any additional errors from e.errors to the log fields
//   - Masks sensitive fields (e.g., "apikey") to protect credentials
//   - Logs the error with full request context using rlog
//
// Security note: Fields with key "apikey" are automatically masked,
// showing only the first 2 and last 2 characters.
//
// Parameters:
//   - c: Gin context containing the HTTP request context
//   - fields: Structured logging fields to include in the log entry
func (e ErrorData) logError(c *gin.Context, fields ...Field) {
	if len(e.errors) > 0 {
		for _, errs := range e.errors {
			fields = append(fields, rlog.String("error", errs.Error()))
		}
	}
	for i, field := range fields {
		if field.Key == "apikey" {
			// Only mask if the field is a string type (non-empty String value)
			if field.String != "" {
				fields[i].String = maskValue(fields[i].String)
			}
		}
	}
	rlog.Errorc(c.Request.Context(), "error", e, fields...)
}
