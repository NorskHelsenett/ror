# RLog - Structured Logging for ROR

## Why RLog

Rlog is a wrapper for the Uber Zap logging framework. The reason we want to wrap
zap and not use it directly is that it makes it easier to change logger in the
future if necessary. RLog was made when we wanted to switch from Logrus to Zap
and encountered this exact problem.

## What RLog offers

Rlog offers a global logger with both context (go context) and non context aware
logging functions. Context aware functions will add certain predefined fields
from a context and also correlate any trace information with the logs.

Rlog offers leveled and structured logging in the same manner as Zap does. The
levels supported are Info, Debug, Warn, Error and Fatal. Logs are normally
output as JSON but can be output in a more readable form when developing, its
important to note that in prod all logs shall be in the JSON format.

## Features

- **Structured Logging**: Log entries include structured fields for better filtering and analysis
- **Context-Aware Logging**: Automatically extract trace IDs and context values for correlated logs
- **OpenTelemetry Integration**: Correlate logs with OpenTelemetry traces
- **Multiple Output Targets**: Send logs to files, stdout/stderr, or other targets
- **Environment-Based Configuration**: Easy configuration via environment variables
- **HTTP Request Logging**: Built-in middleware for Gin web framework
- **Field Helpers**: Convenience functions for adding typed fields to log entries

## Configuration

Rlog is configured using environment variables, from either an env file or the
environment its running in.

### LOG_LEVEL

Sets the minimum log level to output. Valid values are:
- `debug`: Most verbose level, includes detailed diagnostic information
- `trace`: Alias for debug level
- `info`: Standard information messages (default)
- `warn`: Warning messages
- `error`: Error messages only

Example:
```
LOG_LEVEL=debug
```

### LOG_OUTPUT

Set where the logs should be sent, defaults to stderr. Logs can be sent to a
file an url or stdout/stderr. Logs can be sent to multiple locations with a
comma separated string.

Example:
```
LOG_OUTPUT="/home/user/foo/.ror/log,stderr"
```
This example logs to both a file and stderr

### LOG_OUTPUT_ERROR

Set where error level logs should be sent, works the same as LOG_OUTPUT.
If not specified, error logs will go to the same destination as regular logs.

Example:
```
LOG_OUTPUT_ERROR="/home/user/foo/.ror/error.log"
```

### LOG_DEVELOP

Set whether or not logs should be in JSON format or a more human-readable format.
- `false` or unset: JSON format (default for production)
- `true`: Human-readable console format (for development)

Example:
```
LOG_DEVELOP=true
```

## Usage Examples

### Basic Logging

```go
package main

import (
    "github.com/yourusername/ror/pkg/rlog"
)

func main() {
    // Basic logging at different levels
    rlog.Debug("Debug message", rlog.String("component", "example"))
    rlog.Info("Information message", rlog.Int("count", 42))
    rlog.Warn("Warning message")
    rlog.Error("Error occurred", err, rlog.String("operation", "file_read"))
    
    // CAUTION: This will terminate the program
    // rlog.Fatal("Fatal error", err)
}
```

### Context-Aware Logging

```go
package main

import (
    "context"
    
    "github.com/yourusername/ror/pkg/rlog"
)

func processRequest(ctx context.Context, userID string) {
    // Add context keys that should be included in all logs
    rlog.AddContextKeyField("request_id")
    rlog.AddContextKeyField("user_id")
    
    // Context with values
    ctx = context.WithValue(ctx, "request_id", "req-123")
    ctx = context.WithValue(ctx, "user_id", userID)
    
    // Log with context - will include request_id and user_id automatically
    rlog.Infoc(ctx, "Processing request", rlog.String("action", "user_login"))
}
```

### HTTP Middleware

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourusername/ror/pkg/rlog"
)

func setupRouter() *gin.Engine {
    router := gin.New()
    
    // Add the rlog middleware for HTTP request logging
    router.Use(rlog.LogMiddleware())
    
    // Add routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    
    return router
}
```

### Formatted Logging

```go
package main

import (
    "github.com/yourusername/ror/pkg/rlog"
)

func main() {
    // Printf-style logging (at info level)
    rlog.Infof("User %s logged in from %s", "john_doe", "192.168.1.1")
}
```

## Integration with OpenTelemetry

Rlog automatically correlates logs with OpenTelemetry traces when using context-aware logging functions. When a span is active in the provided context, the following trace information is added to log entries:

- `trace_id`: The OpenTelemetry trace ID
- `trace_flags`: The OpenTelemetry trace flags
- `span_id`: The OpenTelemetry span ID (if the span is recording)

Additionally, log entries are recorded as span events in the active span.

## Best Practices

1. **Use Structured Logging**: Always include relevant fields using the field helpers rather than embedding them in the message string
2. **Use Context-Aware Logging**: When a context is available, prefer the context-aware logging functions (Infoc, Debugc, etc.)
3. **Log Levels**: Use appropriate log levels based on the importance and severity of the message
4. **Error Handling**: Always include the error object when logging errors using the Error or Fatal functions
5. **Sensitive Data**: Never log sensitive information such as passwords, tokens, or personal identifiable information

