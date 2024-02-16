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

// AddContextKeyField adds a key to look for in a contexts so that we
// can add the context value as a persistent field to all logs
func AddContextKeyField(key string) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}

	l.ContextKeyFields = append(l.ContextKeyFields, key)
	return nil
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site.
func Info(msg string, fields ...Field) {
	l.Info(msg, fields...)
}

// Debug logs a message at InfoLevel. The message includes any fields passed
// at the log site.
func Debug(msg string, fields ...Field) {
	l.Debug(msg, fields...)
}

// Warn logs a message at InfoLevel. The message includes any fields passed
// at the log site.
func Warn(msg string, fields ...Field) {
	l.Warn(msg, fields...)
}

// Error logs a message at InfoLevel. The message includes any fields passed
// at the log site.
func Error(msg string, err error, fields ...Field) {
	fields = append(fields, zap.Error(err))
	l.Error(msg, fields...)
}

// Fatal logs a message at InfoLevel. The message includes any fields passed
// at the log site.
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, err error, fields ...Field) {
	fields = append(fields, zap.Error(err))
	l.Fatal(msg, fields...)
}

// context aware logfunctions

// Infoc logs a message at InfoLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added
func Infoc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.InfoLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Info(msg, fields...)
}

// Debugc logs a message at DebugLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added
func Debugc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.DebugLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Debug(msg, fields...)
}

// Warnc logs a message at WarnLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added
func Warnc(ctx context.Context, msg string, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.WarnLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	l.Warn(msg, fields...)
}

// Errorc logs a message at ErrorLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added
func Errorc(ctx context.Context, msg string, err error, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.ErrorLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	fields = append(fields, zap.Error(err))
	l.Error(msg, fields...)
}

// Fatalc logs a message at FatalLevel with context. The message includes any fields passed
// at the log site and any tracing fields in the attached context, if the context
// contains known fields these are also added
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatalc(ctx context.Context, msg string, err error, fields ...Field) {
	correlateWithTrace(ctx, msg, zap.ErrorLevel, &fields)
	getFieldsFromContext(ctx, &fields)
	fields = append(fields, zap.Error(err))
	l.Fatal(msg, fields...)
}

// Infof logs a message at InfoLevel with context. The message is formated with sprintf.
func Infof(format string, v ...any) {
	err := fmt.Sprintf(format, v...)
	Info(err)
}

// Field functions

func String(key, value string) Field {
	return zap.String(key, value)
}

func Strings(key string, value []string) Field {
	return zap.Strings(key, value)
}

func ByteString(key string, value []byte) Field {
	return zap.ByteString(key, value)
}

func Int(key string, value int) Field {
	return zap.Int(key, value)
}

func Int64(key string, value int64) Field {
	return zap.Int64(key, value)
}

func Uint(key string, value uint) Field {
	return zap.Uint(key, value)
}

func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

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

// tries to get a predifined set of fields from context, if the field is not
// present int the context it is ignored
func getFieldsFromContext(ctx context.Context, fields *[]Field) {
	Info("looking for fields in context", Any("fields", l.ContextKeyFields))
	for _, key := range l.ContextKeyFields {
		Info("found field", Any("key", key))
		if ctx.Value(key) != nil {
			*fields = append(*fields, zap.Any(key, ctx.Value(key)))
		}
	}
}

// all configuration of rlog should be done through config files or
// environment
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

// looks for LOG_LEVEL in environment to set loglevel
func getLogLevelFromConfig() zapcore.Level {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatability, when the API uses viper for config we can remove
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

// looks for LOG_OUTPUT in environment to set ouput destination
// to define more than one output separate them with ','
// outputs must adhere to zaps requirements
func getOutputsFromConfig() []string {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatability, when the API uses viper for config we can remove
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

// looks for optional LOG_OUTPUT_ERROR in environment to set ouput destination
// to define more than one output separate them with ','
// outputs must adhere to zaps requirements
func getErrorOutputsFromConfig() []string {

	//since the api does not use viper yet we do this check to preserve
	//backwards compatability, when the API uses viper for config we can remove
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
