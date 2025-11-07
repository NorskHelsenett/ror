// Package rorerror provides structured error handling for HTTP APIs with Gin framework integration.
//
// This package simplifies error handling in HTTP services by providing:
//   - Structured error types with HTTP status codes
//   - Automatic JSON serialization
//   - Integration with Gin web framework
//   - Context-aware logging through rlog
//   - Automatic masking of sensitive fields (e.g., API keys)
//
// # Basic Usage
//
// Creating errors:
//
//	// Create error from string message
//	err := rorerror.NewRorError(400, "invalid request")
//
//	// Create error from existing error
//	err := rorerror.NewRorErrorFromError(500, someError)
//
// Using with Gin handlers:
//
//	func Handler(c *gin.Context) {
//	    if err := doSomething(); err != nil {
//	        rorerror := rorerror.NewRorErrorFromError(400, err)
//	        rorerror.GinLogErrorAbort(c)
//	        return
//	    }
//	}
//
// Adding structured fields:
//
//	rorerror.GinLogErrorAbort(c,
//	    rlog.String("user_id", userId),
//	    rlog.Int("attempt", retryCount))
//
// # Migration from v1
//
// The v2 package uses an interface-based design for better flexibility:
//
// v1 code:
//
//	err := rorerror.NewRorError(400, "bad request")
//	err.GinLogErrorAbort(c, rorerror.String("key", "value"))
//
// v2 code (fields now use rlog directly):
//
//	err := rorerror.NewRorError(400, "bad request")
//	err.GinLogErrorAbort(c, rlog.String("key", "value"))
//
// Key changes:
//   - Field type is now rlog.Field (type alias)
//   - GinHandleErrorAndAbort now takes rlog.Field parameters
//   - ErrorData.errors is now private (use GetMessage() for display)
//   - Added RorError interface for better abstraction
//
// # Security Features
//
// The package automatically masks sensitive field values in logs:
//   - Fields with key "apikey" are masked (shows first 2 and last 2 chars)
//   - Masking applies only to log output, not JSON responses
//
// Example:
//
//	// "secret123key" will be logged as "se********ey"
//	rorerror.GinLogErrorAbort(c, rlog.String("apikey", "secret123key"))
package rorerror

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RorError defines the interface for structured HTTP errors.
// Implementations provide HTTP status codes, error messages, and integration
// with the Gin web framework for logging and JSON response handling.
type RorError interface {
	// GetStatusCode returns the HTTP status code for this error.
	GetStatusCode() int

	// GetMessage returns the human-readable error message.
	GetMessage() string

	// Error implements the error interface, returning a formatted error string.
	Error() string

	// GinLogErrorAbort logs the error with context and aborts the Gin request
	// with a JSON response containing the error details.
	GinLogErrorAbort(c *gin.Context, fields ...Field)

	// GinLogErrorJSON logs the error with context and returns a JSON response
	// containing the error details without aborting the request.
	GinLogErrorJSON(c *gin.Context, fields ...Field)
}

// ErrorData is the concrete implementation of RorError.
// It contains the HTTP status code, error message, and optional additional errors.
type ErrorData struct {
	Status  int     `json:"status" example:"400"`          // HTTP status code
	Message string  `json:"message" example:"Bad Request"` // Error message
	errors  []error // Additional errors for internal tracking (not serialized)
}

// NewRorErrorFromError creates a RorError from an existing error and HTTP status code.
// The error's Error() message becomes the error message in the response.
//
// Parameters:
//   - status: HTTP status code (e.g., 400, 500)
//   - err: The source error whose message will be used
//
// Returns:
//   - RorError interface that can be used with Gin handlers
//
// Example:
//
//	if err := db.Query(); err != nil {
//	    return rorerror.NewRorErrorFromError(500, err)
//	}
func NewRorErrorFromError(status int, err error) RorError {
	rorerror := ErrorData{
		Status:  status,
		Message: err.Error(),
	}
	return rorerror
}

// NewRorError creates a new ErrorData with the given status code and message.
// Additional errors can be passed for internal tracking and logging.
//
// Parameters:
//   - status: HTTP status code (e.g., 400, 500)
//   - err: Human-readable error message
//   - errors: Optional additional errors for logging (not included in JSON response)
//
// Returns:
//   - ErrorData struct implementing RorError interface
//
// Example:
//
//	// Simple error
//	err := rorerror.NewRorError(400, "invalid email format")
//
//	// With additional error context
//	err := rorerror.NewRorError(500, "database error", dbErr, connErr)
func NewRorError(status int, err string, errors ...error) ErrorData {
	rorerror := ErrorData{
		Status:  status,
		Message: err,
		errors:  errors,
	}
	return rorerror
}

// Error implements the error interface, returning a formatted string
// containing both the status code and message.
//
// Returns:
//   - Formatted error string: "Status: <code>, Message: <message>"
//
// Example:
//
//	err := rorerror.NewRorError(404, "not found")
//	fmt.Println(err.Error()) // Output: Status: 404, Message: not found
func (e ErrorData) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", e.Status, e.Message)
}

// GetStatusCode returns the HTTP status code associated with this error.
//
// Returns:
//   - HTTP status code as integer
func (e ErrorData) GetStatusCode() int {
	return e.Status
}

// GetMessage returns the human-readable error message.
//
// Returns:
//   - Error message string
func (e ErrorData) GetMessage() string {
	return e.Message
}

