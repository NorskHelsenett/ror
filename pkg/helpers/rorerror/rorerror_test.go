package rorerror

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	rorerrorv2 "github.com/NorskHelsenett/ror/pkg/helpers/rorerror/v2"
	"github.com/gin-gonic/gin"
)

func TestRorErrorBasicBehavior(t *testing.T) {
	t.Run("NewRorError creates correct structure", func(t *testing.T) {
		v1err := NewRorError(400, "bad request")
		v2err := rorerrorv2.NewRorError(400, "bad request")

		if v1err.Status != v2err.GetStatusCode() {
			t.Errorf("Status mismatch: v1=%d, v2=%d", v1err.Status, v2err.GetStatusCode())
		}
		if v1err.Message != v2err.GetMessage() {
			t.Errorf("Message mismatch: v1=%s, v2=%s", v1err.Message, v2err.GetMessage())
		}
	})

	t.Run("NewRorErrorFromError creates correct structure", func(t *testing.T) {
		err := errors.New("test error")
		v1err := NewRorErrorFromError(500, err)
		v2err := rorerrorv2.NewRorErrorFromError(500, err)

		if v1err.Status != v2err.GetStatusCode() {
			t.Errorf("Status mismatch: v1=%d, v2=%d", v1err.Status, v2err.GetStatusCode())
		}
		if v1err.Message != v2err.GetMessage() {
			t.Errorf("Message mismatch: v1=%s, v2=%s", v1err.Message, v2err.GetMessage())
		}
	})

	t.Run("Error() method returns same format", func(t *testing.T) {
		v1err := NewRorError(404, "not found")
		v2err := rorerrorv2.NewRorError(404, "not found")

		if v1err.Error() != v2err.Error() {
			t.Errorf("Error() mismatch: v1=%s, v2=%s", v1err.Error(), v2err.Error())
		}
	})
}

func TestRorErrorJSON(t *testing.T) {
	t.Run("JSON serialization matches", func(t *testing.T) {
		v1err := NewRorError(422, "validation error")
		v2err := rorerrorv2.NewRorError(422, "validation error")

		v1json := v1err.AsJson()
		v2json, _ := json.Marshal(v2err)

		var v1data, v2data map[string]interface{}
		json.Unmarshal(v1json, &v1data)
		json.Unmarshal(v2json, &v2data)

		if v1data["status"] != v2data["status"] {
			t.Errorf("JSON status mismatch: v1=%v, v2=%v", v1data["status"], v2data["status"])
		}
		if v1data["message"] != v2data["message"] {
			t.Errorf("JSON message mismatch: v1=%v, v2=%v", v1data["message"], v2data["message"])
		}
	})

	t.Run("AsString returns correct format", func(t *testing.T) {
		v1err := NewRorError(400, "test")
		expected := `{"status":400,"message":"test"}`

		if v1err.AsString() != expected {
			t.Errorf("AsString() mismatch: got=%s, want=%s", v1err.AsString(), expected)
		}
	})
}

func TestRorErrorUtilities(t *testing.T) {
	t.Run("IsError returns true for non-zero status", func(t *testing.T) {
		v1err := NewRorError(500, "error")
		if !v1err.IsError() {
			t.Error("IsError() should return true for non-zero status")
		}
	})

	t.Run("IsError returns false for zero status", func(t *testing.T) {
		v1err := RorError{}
		if v1err.IsError() {
			t.Error("IsError() should return false for zero status")
		}
	})

	t.Run("NoRorError is empty", func(t *testing.T) {
		if NoRorError.Status != 0 {
			t.Errorf("NoRorError.Status should be 0, got %d", NoRorError.Status)
		}
		if NoRorError.Message != "" {
			t.Errorf("NoRorError.Message should be empty, got %s", NoRorError.Message)
		}
		if NoRorError.IsError() {
			t.Error("NoRorError.IsError() should be false")
		}
	})
}

func TestGinIntegration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GinLogErrorJSON behaves similarly", func(t *testing.T) {
		// Test v1
		w1 := httptest.NewRecorder()
		c1, _ := gin.CreateTestContext(w1)
		c1.Request = httptest.NewRequest("GET", "/test", nil)
		v1err := NewRorError(400, "bad request")
		v1err.GinLogErrorJSON(c1)

		// Test v2
		w2 := httptest.NewRecorder()
		c2, _ := gin.CreateTestContext(w2)
		c2.Request = httptest.NewRequest("GET", "/test", nil)
		v2err := rorerrorv2.NewRorError(400, "bad request")
		v2err.GinLogErrorJSON(c2)

		if w1.Code != w2.Code {
			t.Errorf("Status code mismatch: v1=%d, v2=%d", w1.Code, w2.Code)
		}

		var v1resp, v2resp map[string]interface{}
		json.Unmarshal(w1.Body.Bytes(), &v1resp)
		json.Unmarshal(w2.Body.Bytes(), &v2resp)

		if v1resp["status"] != v2resp["status"] {
			t.Errorf("Response status mismatch: v1=%v, v2=%v", v1resp["status"], v2resp["status"])
		}
		if v1resp["message"] != v2resp["message"] {
			t.Errorf("Response message mismatch: v1=%v, v2=%v", v1resp["message"], v2resp["message"])
		}
	})

	t.Run("GinLogErrorAbort behaves similarly", func(t *testing.T) {
		// Test v1
		w1 := httptest.NewRecorder()
		c1, _ := gin.CreateTestContext(w1)
		c1.Request = httptest.NewRequest("GET", "/test", nil)
		v1err := NewRorError(500, "server error")
		v1err.GinLogErrorAbort(c1)

		// Test v2
		w2 := httptest.NewRecorder()
		c2, _ := gin.CreateTestContext(w2)
		c2.Request = httptest.NewRequest("GET", "/test", nil)
		v2err := rorerrorv2.NewRorError(500, "server error")
		v2err.GinLogErrorAbort(c2)

		if w1.Code != w2.Code {
			t.Errorf("Status code mismatch: v1=%d, v2=%d", w1.Code, w2.Code)
		}

		if c1.IsAborted() != c2.IsAborted() {
			t.Errorf("Abort status mismatch: v1=%t, v2=%t", c1.IsAborted(), c2.IsAborted())
		}
	})
}

func TestGinHandleErrorAndAbort(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("handles error correctly", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := errors.New("test error")
		result := GinHandleErrorAndAbort(c, http.StatusBadRequest, err)

		if !result {
			t.Error("GinHandleErrorAndAbort should return true when error is provided")
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
		if !c.IsAborted() {
			t.Error("Context should be aborted")
		}

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		if resp["status"] != float64(400) {
			t.Errorf("Response status should be 400, got %v", resp["status"])
		}
		if resp["message"] != "test error" {
			t.Errorf("Response message should be 'test error', got %v", resp["message"])
		}
	})

	t.Run("returns false for nil error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		result := GinHandleErrorAndAbort(c, http.StatusOK, nil)

		if result {
			t.Error("GinHandleErrorAndAbort should return false when error is nil")
		}
		if c.IsAborted() {
			t.Error("Context should not be aborted for nil error")
		}
	})
}

func TestFieldConversion(t *testing.T) {
	t.Run("Field types work correctly", func(t *testing.T) {
		strField := String("key", "value")
		intField := Int("count", 42)
		int64Field := Int64("bignum", 9223372036854775807)
		uintField := Uint("unsigned", 123)
		float64Field := Float64("pi", 3.14159)
		float32Field := Float32("small", 2.5)

		if strField.Key != "key" || strField.Value != "value" {
			t.Errorf("String field incorrect: %+v", strField)
		}
		if intField.Key != "count" || intField.Value != "42" {
			t.Errorf("Int field incorrect: %+v", intField)
		}
		if int64Field.Key != "bignum" || int64Field.Value != "9223372036854775807" {
			t.Errorf("Int64 field incorrect: %+v", int64Field)
		}
		if uintField.Key != "unsigned" || uintField.Value != "123" {
			t.Errorf("Uint field incorrect: %+v", uintField)
		}
		if float64Field.Key != "pi" || float64Field.Value != "3.141590" {
			t.Errorf("Float64 field incorrect: %+v", float64Field)
		}
		if float32Field.Key != "small" || float32Field.Value != "2.500000" {
			t.Errorf("Float32 field incorrect: %+v", float32Field)
		}
	})

	t.Run("Stringp handles nil pointer", func(t *testing.T) {
		var ptr *string
		field := Stringp("nullable", ptr)
		if field.Value != "nil" {
			t.Errorf("Stringp should return 'nil' for nil pointer, got %s", field.Value)
		}

		value := "test"
		ptr = &value
		field = Stringp("nullable", ptr)
		if field.Value != "test" {
			t.Errorf("Stringp should return 'test' for valid pointer, got %s", field.Value)
		}
	})
}

func TestApiKeyMasking(t *testing.T) {
	t.Run("maskApiKey masks correctly", func(t *testing.T) {
		apikey := "abc123def456"
		masked := maskApiKey(apikey)
		expected := "ab********56" // Fixed: length-4 characters are masked

		if masked != expected {
			t.Errorf("Expected %s, got %s", expected, masked)
		}
	})

	t.Run("maskApiKey masks short value correctly", func(t *testing.T) {
		apikey := "ab45"
		masked := maskApiKey(apikey)
		expected := "****" // Fixed: length-4 characters are masked

		if masked != expected {
			t.Errorf("Expected %s, got %s", expected, masked)
		}
	})

	t.Run("GinHandleErrorAndAbort masks apikey field", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/test", nil)

		err := errors.New("auth error")
		result := GinHandleErrorAndAbort(c, http.StatusUnauthorized, err, String("apikey", "secret123key"))

		if !result {
			t.Error("Should handle error")
		}
		// Note: The masking happens in logging, not in the response JSON
		// This test verifies the function doesn't crash with apikey fields
	})
}
