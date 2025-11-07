package rorerror

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/gin-gonic/gin"
)

func TestNewRorError(t *testing.T) {
	t.Run("creates error with status and message", func(t *testing.T) {
		err := NewRorError(400, "bad request")

		if err.Status != 400 {
			t.Errorf("Expected status 400, got %d", err.Status)
		}
		if err.Message != "bad request" {
			t.Errorf("Expected message 'bad request', got %s", err.Message)
		}
		if len(err.errors) != 0 {
			t.Errorf("Expected no additional errors, got %d", len(err.errors))
		}
	})

	t.Run("creates error with additional errors", func(t *testing.T) {
		err1 := errors.New("validation error 1")
		err2 := errors.New("validation error 2")

		err := NewRorError(422, "validation failed", err1, err2)

		if err.Status != 422 {
			t.Errorf("Expected status 422, got %d", err.Status)
		}
		if err.Message != "validation failed" {
			t.Errorf("Expected message 'validation failed', got %s", err.Message)
		}
		if len(err.errors) != 2 {
			t.Errorf("Expected 2 additional errors, got %d", len(err.errors))
		}
	})

	t.Run("creates error with no additional errors", func(t *testing.T) {
		err := NewRorError(404, "not found")

		if len(err.errors) != 0 {
			t.Errorf("Expected no additional errors, got %d", len(err.errors))
		}
	})
}

func TestNewRorErrorFromError(t *testing.T) {
	t.Run("creates error from existing error", func(t *testing.T) {
		sourceErr := errors.New("database connection failed")
		rorErr := NewRorErrorFromError(500, sourceErr)

		if rorErr.GetStatusCode() != 500 {
			t.Errorf("Expected status 500, got %d", rorErr.GetStatusCode())
		}
		if rorErr.GetMessage() != "database connection failed" {
			t.Errorf("Expected message 'database connection failed', got %s", rorErr.GetMessage())
		}
	})

	t.Run("returns RorError interface", func(t *testing.T) {
		sourceErr := errors.New("test error")
		var rorErr RorError = NewRorErrorFromError(400, sourceErr)

		if rorErr == nil {
			t.Error("Expected non-nil RorError interface")
		}
	})
}

func TestErrorData_Error(t *testing.T) {
	t.Run("returns formatted error string", func(t *testing.T) {
		err := NewRorError(404, "resource not found")
		expected := "Status: 404, Message: resource not found"

		if err.Error() != expected {
			t.Errorf("Expected '%s', got '%s'", expected, err.Error())
		}
	})

	t.Run("handles different status codes", func(t *testing.T) {
		testCases := []struct {
			status   int
			message  string
			expected string
		}{
			{200, "success", "Status: 200, Message: success"},
			{400, "bad request", "Status: 400, Message: bad request"},
			{500, "internal error", "Status: 500, Message: internal error"},
		}

		for _, tc := range testCases {
			err := NewRorError(tc.status, tc.message)
			if err.Error() != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, err.Error())
			}
		}
	})
}

func TestErrorData_GetStatusCode(t *testing.T) {
	testCases := []int{200, 400, 401, 403, 404, 422, 500, 503}

	for _, status := range testCases {
		err := NewRorError(status, "test")
		if err.GetStatusCode() != status {
			t.Errorf("Expected status %d, got %d", status, err.GetStatusCode())
		}
	}
}

func TestErrorData_GetMessage(t *testing.T) {
	testCases := []string{
		"simple message",
		"message with special chars: !@#$%",
		"multiline\nmessage",
		"",
	}

	for _, msg := range testCases {
		err := NewRorError(400, msg)
		if err.GetMessage() != msg {
			t.Errorf("Expected message '%s', got '%s'", msg, err.GetMessage())
		}
	}
}

func TestErrorData_GinLogErrorJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("sends JSON response without aborting", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := NewRorError(400, "validation failed")
		err.GinLogErrorJSON(c)

		if w.Code != 400 {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		if response["status"] != float64(400) {
			t.Errorf("Expected status 400 in JSON, got %v", response["status"])
		}
		if response["message"] != "validation failed" {
			t.Errorf("Expected message 'validation failed', got %v", response["message"])
		}

		if c.IsAborted() {
			t.Error("Context should not be aborted with GinLogErrorJSON")
		}
	})

	t.Run("handles fields correctly", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := NewRorError(500, "server error")
		err.GinLogErrorJSON(c, rlog.String("user_id", "123"))

		if w.Code != 500 {
			t.Errorf("Expected status 500, got %d", w.Code)
		}
	})
}

func TestErrorData_GinLogErrorAbort(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("sends JSON response and aborts", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := NewRorError(403, "forbidden")
		err.GinLogErrorAbort(c)

		if w.Code != 403 {
			t.Errorf("Expected status 403, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		if response["status"] != float64(403) {
			t.Errorf("Expected status 403 in JSON, got %v", response["status"])
		}
		if response["message"] != "forbidden" {
			t.Errorf("Expected message 'forbidden', got %v", response["message"])
		}

		if !c.IsAborted() {
			t.Error("Context should be aborted with GinLogErrorAbort")
		}
	})

	t.Run("logs additional errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		err := NewRorError(500, "multiple errors", err1, err2)
		err.GinLogErrorAbort(c)

		// The additional errors should be logged but not in JSON response
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		if response["message"] != "multiple errors" {
			t.Errorf("Expected message 'multiple errors', got %v", response["message"])
		}
	})
}

func TestGinHandleErrorAndAbort(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns false and does nothing when error is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		result := GinHandleErrorAndAbort(c, 500, nil)

		if result {
			t.Error("Expected false when error is nil")
		}
		if c.IsAborted() {
			t.Error("Context should not be aborted when error is nil")
		}
		if w.Code != 200 {
			t.Errorf("Expected status 200 (default), got %d", w.Code)
		}
	})

	t.Run("returns true and handles error when error is not nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := errors.New("database error")
		result := GinHandleErrorAndAbort(c, 500, err)

		if !result {
			t.Error("Expected true when error is not nil")
		}
		if !c.IsAborted() {
			t.Error("Context should be aborted when error is not nil")
		}
		if w.Code != 500 {
			t.Errorf("Expected status 500, got %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		if response["message"] != "database error" {
			t.Errorf("Expected message 'database error', got %v", response["message"])
		}
	})

	t.Run("includes fields in log", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := errors.New("query failed")
		result := GinHandleErrorAndAbort(c, 500, err,
			rlog.String("operation", "fetch"),
			rlog.String("table", "users"))

		if !result {
			t.Error("Expected true when error is not nil")
		}
	})

	t.Run("handles different status codes", func(t *testing.T) {
		testCases := []int{400, 401, 403, 404, 500, 503}

		for _, status := range testCases {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/test", nil)

			err := errors.New("test error")
			GinHandleErrorAndAbort(c, status, err)

			if w.Code != status {
				t.Errorf("Expected status %d, got %d", status, w.Code)
			}
		}
	})
}

func TestMaskValue(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard API key",
			input:    "secret123key",
			expected: "se********ey",
		},
		{
			name:     "long value",
			input:    "verylongsecretkey123456",
			expected: "ve*******************56",
		},
		{
			name:     "exactly 4 characters",
			input:    "abcd",
			expected: "abcd",
		},
		{
			name:     "short value (3 chars)",
			input:    "abc",
			expected: "***",
		},
		{
			name:     "short value (2 chars)",
			input:    "ab",
			expected: "**",
		},
		{
			name:     "short value (1 char)",
			input:    "a",
			expected: "*",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "5 characters",
			input:    "12345",
			expected: "12*45",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maskValue(tc.input)
			if result != tc.expected {
				t.Errorf("maskValue(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestErrorData_JSONSerialization(t *testing.T) {
	t.Run("marshals to JSON correctly", func(t *testing.T) {
		err := NewRorError(404, "not found")
		data, jsonErr := json.Marshal(err)

		if jsonErr != nil {
			t.Fatalf("Failed to marshal: %v", jsonErr)
		}

		var result map[string]interface{}
		json.Unmarshal(data, &result)

		if result["status"] != float64(404) {
			t.Errorf("Expected status 404, got %v", result["status"])
		}
		if result["message"] != "not found" {
			t.Errorf("Expected message 'not found', got %v", result["message"])
		}
		// errors field should not be serialized
		if _, exists := result["errors"]; exists {
			t.Error("errors field should not be serialized")
		}
	})

	t.Run("additional errors are not serialized", func(t *testing.T) {
		err1 := errors.New("internal error 1")
		err2 := errors.New("internal error 2")
		err := NewRorError(500, "server error", err1, err2)

		data, _ := json.Marshal(err)
		var result map[string]interface{}
		json.Unmarshal(data, &result)

		if _, exists := result["errors"]; exists {
			t.Error("errors field should not be in JSON output")
		}
	})
}

func TestRorErrorInterface(t *testing.T) {
	t.Run("ErrorData implements RorError interface", func(t *testing.T) {
		var _ RorError = NewRorError(400, "test")
	})

	t.Run("NewRorErrorFromError returns RorError interface", func(t *testing.T) {
		var _ RorError = NewRorErrorFromError(500, errors.New("test"))
	})

	t.Run("interface methods are accessible", func(t *testing.T) {
		var rorErr RorError = NewRorError(400, "test message")

		if rorErr.GetStatusCode() != 400 {
			t.Errorf("Expected status 400, got %d", rorErr.GetStatusCode())
		}
		if rorErr.GetMessage() != "test message" {
			t.Errorf("Expected message 'test message', got %s", rorErr.GetMessage())
		}
		if rorErr.Error() != "Status: 400, Message: test message" {
			t.Errorf("Unexpected Error() output: %s", rorErr.Error())
		}
	})
}

func TestLogError_WithAPIKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("masks apikey field in logs", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := NewRorError(401, "unauthorized")
		err.GinLogErrorAbort(c, rlog.String("apikey", "secret123key"))

		// The apikey should be masked in logs (we can't easily verify log output in tests,
		// but we can verify the function doesn't panic)
		if !c.IsAborted() {
			t.Error("Context should be aborted")
		}
	})

	t.Run("handles non-string apikey field", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := NewRorError(401, "unauthorized")
		// This tests that the masking logic handles non-string fields gracefully
		err.GinLogErrorAbort(c, rlog.Int("apikey", 12345))

		if !c.IsAborted() {
			t.Error("Context should be aborted")
		}
	})
}

func TestErrorData_WithMultipleErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("logs all additional errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err1 := errors.New("validation error: email invalid")
		err2 := errors.New("validation error: age must be positive")
		err3 := errors.New("validation error: name required")

		rorErr := NewRorError(422, "validation failed", err1, err2, err3)
		rorErr.GinLogErrorAbort(c)

		// Verify the main message is in the response
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		if response["message"] != "validation failed" {
			t.Errorf("Expected message 'validation failed', got %v", response["message"])
		}
	})

	t.Run("handles empty additional errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		rorErr := NewRorError(400, "simple error")
		rorErr.GinLogErrorAbort(c)

		if !c.IsAborted() {
			t.Error("Context should be aborted")
		}
	})
}

func TestEdgeCases(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("handles zero status code", func(t *testing.T) {
		err := NewRorError(0, "no status")
		if err.GetStatusCode() != 0 {
			t.Errorf("Expected status 0, got %d", err.GetStatusCode())
		}
	})

	t.Run("handles negative status code", func(t *testing.T) {
		err := NewRorError(-1, "negative status")
		if err.GetStatusCode() != -1 {
			t.Errorf("Expected status -1, got %d", err.GetStatusCode())
		}
	})

	t.Run("handles empty message", func(t *testing.T) {
		err := NewRorError(400, "")
		if err.GetMessage() != "" {
			t.Errorf("Expected empty message, got %s", err.GetMessage())
		}
	})

	t.Run("handles unicode in message", func(t *testing.T) {
		message := "错误: データエラー, 🚨 alert"
		err := NewRorError(400, message)
		if err.GetMessage() != message {
			t.Errorf("Expected unicode message preserved, got %s", err.GetMessage())
		}
	})

	t.Run("handles very long message", func(t *testing.T) {
		longMessage := string(make([]byte, 10000))
		err := NewRorError(400, longMessage)
		if len(err.GetMessage()) != 10000 {
			t.Errorf("Expected message length 10000, got %d", len(err.GetMessage()))
		}
	})
}

func BenchmarkNewRorError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRorError(400, "test error")
	}
}

func BenchmarkNewRorErrorFromError(b *testing.B) {
	err := errors.New("test error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewRorErrorFromError(500, err)
	}
}

func BenchmarkGinHandleErrorAndAbort(b *testing.B) {
	gin.SetMode(gin.TestMode)
	err := errors.New("test error")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)
		GinHandleErrorAndAbort(c, 500, err)
	}
}

func BenchmarkMaskValue(b *testing.B) {
	value := "secret123key456789"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maskValue(value)
	}
}
