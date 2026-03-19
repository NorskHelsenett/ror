package rortracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var noopTracer = noop.NewTracerProvider().Tracer("")

// StartSpan starts a new span using the global tracer singleton.
// If the tracer has not been initialized or tracing is suppressed on the
// context, a noop span is returned that is safe to use but records nothing.
func StartSpan(ctx context.Context, name string, opts ...oteltrace.SpanStartOption) (context.Context, oteltrace.Span) {
	if IsTracingSuppressed(ctx) || globalTracer == nil {
		return noopTracer.Start(ctx, name, opts...)
	}
	return globalTracer.tracer.Start(ctx, name, opts...)
}

// SpanError records an error on the span, sets the span status to Error,
// and returns the original error for convenient inline use:
//
//	return trace.SpanError(span, err)
func SpanError(span oteltrace.Span, err error) error {
	if err == nil || span == nil {
		return err
	}
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	return err
}

// SpanErrorf creates an error with fmt.Errorf, records it on the span,
// sets the span status to Error, and returns the new error.
func SpanErrorf(span oteltrace.Span, format string, args ...any) error {
	err := fmt.Errorf(format, args...)
	return SpanError(span, err)
}

// SpanOk sets the span status to OK, indicating the operation completed
// successfully.
func SpanOk(span oteltrace.Span) {
	if span == nil {
		return
	}
	span.SetStatus(codes.Ok, "")
}
