package rorerror_test

import (
	"errors"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror/v2"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

// Example_basicError demonstrates creating a simple error with a custom message.
func Example_basicError() {
	err := rorerror.NewRorError(400, "invalid email format")

	// Use in logging or responses
	_ = err.GetStatusCode() // 400
	_ = err.GetMessage()    // "invalid email format"
}

// Example_errorFromExisting demonstrates creating an error from an existing error.
func Example_errorFromExisting() {
	// Simulate a database error
	dbErr := errors.New("connection timeout")

	// Wrap it with HTTP status
	httpErr := rorerror.NewRorErrorFromError(500, dbErr)

	_ = httpErr.GetMessage() // "connection timeout"
}

// Example_ginHandler demonstrates typical usage in a Gin HTTP handler.
func Example_ginHandler() {
	router := gin.New()

	router.GET("/users/:id", func(c *gin.Context) {
		// Validate input
		userID := c.Param("id")
		if userID == "" {
			err := rorerror.NewRorError(400, "user ID is required")
			err.GinLogErrorAbort(c, rlog.String("endpoint", "/users/:id"))
			return
		}

		// Simulate database lookup
		user, err := findUser(userID)
		if rorerror.GinHandleErrorAndAbort(c, 500, err, rlog.String("user_id", userID)) {
			return // Error was handled automatically
		}

		// Success response
		c.JSON(200, user)
	})
}

// Example_withStructuredFields demonstrates adding structured logging fields.
func Example_withStructuredFields() {
	router := gin.New()

	router.POST("/orders", func(c *gin.Context) {
		orderID := "ORD-12345"
		userID := "user-abc"

		err := validateOrder(orderID)
		if err != nil {
			httpErr := rorerror.NewRorErrorFromError(400, err)
			httpErr.GinLogErrorAbort(c,
				rlog.String("order_id", orderID),
				rlog.String("user_id", userID),
				rlog.Int("validation_step", 1))
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})
}

// Example_securityMasking demonstrates automatic masking of sensitive fields.
func Example_securityMasking() {
	router := gin.New()

	router.POST("/auth", func(c *gin.Context) {
		apiKey := "secret123key456"

		if !isValidKey(apiKey) {
			err := rorerror.NewRorError(401, "invalid API key")
			// The apikey field will be masked in logs as "se**********56"
			err.GinLogErrorAbort(c, rlog.String("apikey", apiKey))
			return
		}

		c.JSON(200, gin.H{"status": "authenticated"})
	})
}

// Example_multipleErrors demonstrates tracking multiple errors internally.
func Example_multipleErrors() {
	// Collect multiple validation errors
	var validationErrors []error
	validationErrors = append(validationErrors, errors.New("email is invalid"))
	validationErrors = append(validationErrors, errors.New("age must be positive"))

	// Create error with all validation failures
	// The additional errors are logged but not included in JSON response
	httpErr := rorerror.NewRorError(
		400,
		"validation failed",
		validationErrors...,
	)

	_ = httpErr.GetMessage() // "validation failed"
	// Additional errors are logged automatically when calling GinLogErrorAbort
}

// Example_earlyReturn demonstrates using GinHandleErrorAndAbort for clean early returns.
func Example_earlyReturn() {
	router := gin.New()

	router.GET("/reports/:id", func(c *gin.Context) {
		reportID := c.Param("id")

		// Step 1: Check permissions
		err := checkPermissions(c, reportID)
		if rorerror.GinHandleErrorAndAbort(c, http.StatusForbidden, err,
			rlog.String("report_id", reportID)) {
			return
		}

		// Step 2: Fetch from database
		report, err := fetchReport(reportID)
		if rorerror.GinHandleErrorAndAbort(c, http.StatusInternalServerError, err,
			rlog.String("report_id", reportID),
			rlog.String("operation", "fetch")) {
			return
		}

		// Step 3: Generate PDF
		pdf, err := generatePDF(report)
		if rorerror.GinHandleErrorAndAbort(c, http.StatusInternalServerError, err,
			rlog.String("report_id", reportID),
			rlog.String("operation", "generate_pdf")) {
			return
		}

		c.Data(200, "application/pdf", pdf)
	})
}

// Example_jsonVsAbort demonstrates the difference between JSON and Abort methods.
func Example_jsonVsAbort() {
	router := gin.New()

	// Using GinLogErrorJSON - allows further processing
	router.GET("/partial", func(c *gin.Context) {
		if someCondition() {
			err := rorerror.NewRorError(400, "warning: partial data")
			err.GinLogErrorJSON(c) // Logs and sends JSON, but continues
			// Can still add more to the response or run cleanup
		}
		// Handler continues...
	})

	// Using GinLogErrorAbort - stops processing
	router.GET("/strict", func(c *gin.Context) {
		if someCondition() {
			err := rorerror.NewRorError(400, "invalid request")
			err.GinLogErrorAbort(c) // Logs, sends JSON, and aborts
			return                  // Nothing after this executes
		}
	})
}

// Mock functions for examples
func findUser(id string) (interface{}, error)          { return nil, nil }
func validateOrder(id string) error                    { return nil }
func isValidKey(key string) bool                       { return true }
func checkPermissions(c *gin.Context, id string) error { return nil }
func fetchReport(id string) (interface{}, error)       { return nil, nil }
func generatePDF(report interface{}) ([]byte, error)   { return nil, nil }
func someCondition() bool                              { return false }
