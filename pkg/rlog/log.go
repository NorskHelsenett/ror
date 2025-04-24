// Package rlog provides a structured logging library for the ROR project.
//
// rlog is a wrapper around the uber-go/zap logging package that provides:
// - Structured logging with strongly typed fields
// - Context-aware logging with automatic extraction of request IDs and trace data
// - OpenTelemetry integration for distributed tracing
// - Configurable outputs via environment variables
// - Support for both development (human-readable) and production (JSON) formats
// - HTTP middleware for Gin web framework
//
// The package is configured via environment variables:
// - LOG_LEVEL: Sets the minimum log level (debug, info, warn, error)
// - LOG_OUTPUT: Specifies where logs are written (stderr by default, can be files or multiple targets)
// - LOG_OUTPUT_ERROR: Specifies where error logs are written
// - LOG_DEVELOP: When "true", outputs human-readable logs instead of JSON
//
// Basic usage:
//
//	rlog.Info("This is an informational message", rlog.String("key", "value"))
//	rlog.Error("An error occurred", err, rlog.Int("status", 500))
//
//	// With context (includes trace IDs and context values automatically)
//	rlog.Infoc(ctx, "Processing request", rlog.String("user", "admin"))
package rlog

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CorrelationIdType int

const (
	RequestIdKey CorrelationIdType = iota
	SessionIdKey
	LOG_LEVEL        = "LOG_LEVEL"
	LOG_OUTPUT       = "LOG_OUTPUT"
	LOG_OUTPUT_ERROR = "LOG_OUTPUT_ERROR"
	LOG_DEVELOP      = "LOG_DEVELOP"
)

type Field = zap.Field

var (
	l        Logger
	logLevel zapcore.Level
)

type Logger struct {
	*zap.Logger
	ContextKeyFields []string
}

func init() {
	InitializeRlog()
}

// InitializeRlog initializes the global logger based on configuration.
// It creates either a default or development logger configuration depending on environment
// settings and initializes a global logger instance.
// If initialization fails, it will panic with an error message.
func InitializeRlog() {
	zapConfig := createDefaultRLogConfig()
	if getIsDevelop() {
		zapConfig = createDevelopRLogConfig()
	}

	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(fmt.Errorf("unable to initialize logger: %w", err))
	}

	l = Logger{
		Logger:           zapLogger,
		ContextKeyFields: []string{},
	}

	Debug("global logger initialized", String("level", logLevel.String()))
}

// GetLogLevel returns the current log level as a string.
// The log level indicates the minimum severity of messages that will be logged.
func GetLogLevel() string {
	return logLevel.String()
}

// AddContextKeyField adds a key to look for in contexts so that we
// can add the context value as a persistent field to all logs.
// This allows automatic inclusion of specified context values in log entries.
//
// Parameters:
//   - key: The context key to extract values from. Must be non-empty.
//
// Returns:
//   - An error if the key is empty, nil otherwise.
func AddContextKeyField(key string) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}

	l.ContextKeyFields = append(l.ContextKeyFields, key)
	return nil
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site.
//
// Parameters:
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Info(msg string, fields ...Field) {
	l.Info(msg, fields...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site.
//
// Parameters:
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Debug(msg string, fields ...Field) {
	l.Debug(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site.
//
// Parameters:
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Warn(msg string, fields ...Field) {
	l.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site and the error.
//
// Parameters:
//   - msg: The log message
//   - err: The error to log
//   - fields: Optional fields to add context to the log entry
func Error(msg string, err error, fields ...Field) {
	fields = append(fields, zap.Error(err))
	l.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site and the error.
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
//
// Parameters:
//   - msg: The log message
//   - err: The error to log
//   - fields: Optional fields to add context to the log entry
func Fatal(msg string, err error, fields ...Field) {
	fields = append(fields, zap.Error(err))
	l.Fatal(msg, fields...)
}

// context aware logfunctions

// Infoc logs a message at InfoLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added.
//
// Parameters:
//   - ctx: The context which may contain trace information
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Infoc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.InfoLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Info(msg, fields...)
}

// Debugc logs a message at DebugLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added.
//
// Parameters:
//   - ctx: The context which may contain trace information
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Debugc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.DebugLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Debug(msg, fields...)
}

// Warnc logs a message at WarnLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added.
//
// Parameters:
//   - ctx: The context which may contain trace information
//   - msg: The log message
//   - fields: Optional fields to add context to the log entry
func Warnc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.WarnLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Warn(msg, fields...)
}

// Errorc logs a message at ErrorLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added.
//
// Parameters:
//   - ctx: The context which may contain trace information
//   - msg: The log message
//   - err: The error to log
//   - fields: Optional fields to add context to the log entry
func Errorc(ctx context.Context, msg string, err error, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.ErrorLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	fields = append(fields, zap.Error(err))
	l.Error(msg, fields...)
}

// Fatalc logs a message at FatalLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added.
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
//
// Parameters:
//   - ctx: The context which may contain trace information
//   - msg: The log message
//   - err: The error to log
//   - fields: Optional fields to add context to the log entry
func Fatalc(ctx context.Context, msg string, err error, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.ErrorLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	fields = append(fields, zap.Error(err))
	l.Fatal(msg, fields...)
}

// Infof logs a message at InfoLevel with context. The message is formated with sprintf.
//
// Parameters:
//   - format: A format string for the message
//   - v: Values to be formatted into the message
func Infof(format string, v ...any) {
	err := fmt.Sprintf(format, v...)
	Info(err)
}

// Field functions

// String creates a field with the given key and string value.
//
// Parameters:
//   - key: The field key
//   - value: The string value
//
// Returns:
//   - A Field object that can be used in logging functions
func String(key, value string) Field {
	return zap.String(key, value)
}

// Strings creates a field with the given key and string slice value.
//
// Parameters:
//   - key: The field key
//   - value: The string slice value
//
// Returns:
//   - A Field object that can be used in logging functions
func Strings(key string, value []string) Field {
	return zap.Strings(key, value)
}

// ByteString creates a field with the given key and byte slice value.
//
// Parameters:
//   - key: The field key
//   - value: The byte slice value
//
// Returns:
//   - A Field object that can be used in logging functions
func ByteString(key string, value []byte) Field {
	return zap.ByteString(key, value)
}

// Int creates a field with the given key and integer value.
//
// Parameters:
//   - key: The field key
//   - value: The integer value
//
// Returns:
//   - A Field object that can be used in logging functions
func Int(key string, value int) Field {
	return zap.Int(key, value)
}

// Int64 creates a field with the given key and int64 value.
//
// Parameters:
//   - key: The field key
//   - value: The int64 value
//
// Returns:
//   - A Field object that can be used in logging functions
func Int64(key string, value int64) Field {
	return zap.Int64(key, value)
}

// Uint creates a field with the given key and unsigned integer value.
//
// Parameters:
//   - key: The field key
//   - value: The unsigned integer value
//
// Returns:
//   - A Field object that can be used in logging functions
func Uint(key string, value uint) Field {
	return zap.Uint(key, value)
}

// Any creates a field with the given key and arbitrary value.
// Any handles JSON marshaling for arbitrary objects.
//
// Parameters:
//   - key: The field key
//   - value: The value to be logged (will be marshaled to JSON)
//
// Returns:
//   - A Field object that can be used in logging functions
func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

// correlateWithTrace adds OpenTelemetry trace context information to log fields.
// It extracts trace ID, span ID, and trace flags from the current span in the provided context
// and adds them as fields to the log entry. It also records the log as a span event.
//
// Parameters:
//   - ctx: The context containing the current span
//   - msg: The log message to be recorded as a span event
//   - lvl: The log level of the message
//   - fields: Pointer to the slice of fields to append trace information to
func correlateWithTrace(ctx context.Context, msg string, lvl zapcore.Level, fields *[]Field) {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().TraceID().IsValid() {
		return
	}
	var attributes []attribute.KeyValue
	*fields = append(*fields, zap.String("trace_id", span.SpanContext().TraceID().String()))
	*fields = append(*fields, zap.String("trace_flags", span.SpanContext().TraceFlags().String()))
	attributes = append(attributes, attribute.Key("log.severity").String(lvl.CapitalString()))
	attributes = append(attributes, attribute.Key("log.message").String(msg))
	span.AddEvent("log", trace.WithAttributes(attributes...))
	if span.IsRecording() {
		*fields = append(*fields, zap.String("span_id", span.SpanContext().SpanID().String()))
	}
}

// getFieldsFromContext extracts predefined fields from the context and adds them to the log fields.
// It looks for any keys registered in the logger's ContextKeyFields and adds their values
// to the log entry if present in the context.
//
// Parameters:
//   - ctx: The context from which to extract field values
//   - fields: Pointer to the slice of fields to append context values to
func getFieldsFromContext(ctx context.Context, fields *[]Field) {
	for _, key := range l.ContextKeyFields {
		if ctx.Value(key) != nil {
			*fields = append(*fields, zap.Any(key, ctx.Value(key)))
		}
	}
}

// createDefaultRLogConfig creates a default production configuration for the zap logger.
// It configures JSON output format with production settings and reads the log level
// and output paths from environment variables or configuration.
//
// Returns:
//   - A zap.Config with production settings
func createDefaultRLogConfig() zap.Config {

	infoLevel := getLogLevelFromConfig()

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(infoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      getOutputsFromConfig(),
		ErrorOutputPaths: getErrorOutputsFromConfig(),
	}
}

// createDevelopRLogConfig creates a development configuration for the zap logger.
// It configures a console output format with development settings and reads the log level
// and output paths from environment variables or configuration.
//
// Returns:
//   - A zap.Config with development settings
func createDevelopRLogConfig() zap.Config {
	infoLevel := getLogLevelFromConfig()

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(infoLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      getOutputsFromConfig(),
		ErrorOutputPaths: getErrorOutputsFromConfig(),
	}
}

// getLogLevelFromConfig retrieves the log level from environment variables or configuration.
// It checks for a LOG_LEVEL environment variable first, then falls back to viper config.
// Valid log levels are "debug", "trace", "info", "warn", and "error".
// If the log level is not set or is invalid, it defaults to "info".
//
// Returns:
//   - The configured zapcore.Level
func getLogLevelFromConfig() zapcore.Level {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatibility, when the API uses viper for config we can remove
	//this
	logLevelConfig, present := os.LookupEnv("LOG_LEVEL")
	if !present {
		logLevelConfig = viper.GetString(LOG_LEVEL)
	}
	if len(logLevelConfig) != 0 {
		switch strings.ToLower(logLevelConfig) {
		case "debug":
			logLevel = zap.DebugLevel
		case "trace":
			logLevel = zap.DebugLevel
		case "info":
			logLevel = zap.InfoLevel
		case "warn":
			logLevel = zap.WarnLevel
		case "error":
			logLevel = zap.ErrorLevel
		default:
			logLevel = zap.InfoLevel
			fmt.Printf("unsupported loglevel: %s, defaulting to info\n", logLevelConfig)
		}
	} else {
		logLevel = zap.InfoLevel
	}

	return logLevel
}

// getOutputsFromConfig retrieves the log output destinations from environment variables or configuration.
// It checks for a LOG_OUTPUT environment variable first, then falls back to viper config.
// Multiple outputs can be specified by separating them with commas.
// If no outputs are specified, it defaults to "stderr".
//
// Returns:
//   - A slice of output destination strings
func getOutputsFromConfig() []string {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatibility, when the API uses viper for config we can remove
	//this
	outputsString, present := os.LookupEnv("LOG_OUTPUT")
	if !present {
		outputsString = viper.GetString(LOG_OUTPUT)
	}

	// if LOG_OUTPUT is not set, we default to stderr
	if len(outputsString) == 0 {
		return []string{"stderr"}
	}

	outputs := strings.Split(outputsString, ",")
	return outputs
}

// getErrorOutputsFromConfig retrieves the error log output destinations from environment variables or configuration.
// It checks for a LOG_OUTPUT_ERROR environment variable first, then falls back to viper config.
// Multiple outputs can be specified by separating them with commas.
// If no error outputs are specified, it uses the same outputs as regular logs.
//
// Returns:
//   - A slice of error output destination strings
func getErrorOutputsFromConfig() []string {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatibility, when the API uses viper for config we can remove
	//this
	outputsString, present := os.LookupEnv("LOG_OUTPUT_ERROR")
	if !present {
		outputsString = viper.GetString(LOG_OUTPUT_ERROR)
	}

	// if LOG_OUTPUT_ERROR is not set, we want to get outputs from LOG_OUTPUT
	// instad
	if len(outputsString) == 0 {
		return getOutputsFromConfig()
	}

	outputs := strings.Split(outputsString, ",")

	return outputs
}

// getIsDevelop checks if development mode is enabled for logging.
// It checks for a LOG_DEVELOP environment variable first, then falls back to viper config.
// Development mode is enabled if the value is "true" (case-insensitive).
//
// Returns:
//   - true if development mode is enabled, false otherwise
func getIsDevelop() bool {
	isdevelop, present := os.LookupEnv("LOG_DEVELOP")
	if !present {
		isdevelop = viper.GetString(LOG_DEVELOP)
	}

	if strings.ToLower(isdevelop) == "true" {
		return true
	}

	return false
}
