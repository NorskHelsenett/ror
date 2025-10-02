package trace

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tracerProvider *sdktrace.TracerProvider
	sleepTime      = time.Minute
	GrpcEndpoint   = "localhost:4317"
)

func newTracerProvider(traceExporter sdktrace.SpanExporter, serviceName string) *sdktrace.TracerProvider {

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		),
	)

	if err != nil {
		return nil
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1))),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
	)
}

func newExporter(ctx context.Context, grpcEndpoint string) (*otlptrace.Exporter, error) {

	if grpcEndpoint != "" {
		GrpcEndpoint = grpcEndpoint
	}

	conn, err := grpc.NewClient(
		GrpcEndpoint,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func initTracerProvider(ctx context.Context, serviceName string, grpcEndpoint string) (func(context.Context) error, error) {

	traceExporter, err := newExporter(ctx, grpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	tracerProvider = newTracerProvider(traceExporter, serviceName)
	if tracerProvider == nil {
		return nil, fmt.Errorf("failed to create trace provider")
	}

	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tracerProvider.Shutdown, nil
}

func ConnectTracer(stop chan struct{}, serviceName string, grpcEndpoint string) {
	ctx := context.Background()
	rlog.Infoc(ctx, "Connecting to opentelemetry collector...")
	shutdown, err := initTracerProvider(ctx, serviceName, grpcEndpoint)
	for err != nil {
		rlog.Errorc(ctx, "could not connect to opentelemetry", err)
		if sleepTime <= time.Minute*256 {
			rlog.Infoc(ctx, fmt.Sprintf("Retrying in %s", sleepTime))
			time.Sleep(sleepTime)
			sleepTime = 2 * sleepTime
		} else {
			break
		}
		rlog.Infoc(ctx, "Connecting to opentelemetry collector...")
		shutdown, err = initTracerProvider(ctx, serviceName, grpcEndpoint)
	}
	if err == nil {
		rlog.Infoc(ctx, "Connected successfully to opentelemetry collector on "+GrpcEndpoint)
	}
	defer func() {
		rlog.Infoc(ctx, "Shutting down TracerProvider")
		if err := shutdown(ctx); err != nil {
			rlog.Errorc(ctx, "failed to shutdown TracerProvider", err)
		} else {
			rlog.Infoc(ctx, "TracerProvider shut down successfully")
		}
	}()
	<-stop
}

func Tracer(serviceName string) oteltrace.Tracer {
	if tracerProvider == nil {
		rlog.Warn("TracerProvider is not initialized. Returning a no-op tracer.")
		return otel.Tracer("noop")
	}
	return tracerProvider.Tracer(serviceName)
}

func StartTracing(stop chan struct{}, cancelChan chan os.Signal, serviceName string, grpcEndpoint string) {
	go func() {
		ConnectTracer(stop, serviceName, grpcEndpoint)
		sig := <-cancelChan
		rlog.Info("Caught signal", rlog.Any("signal", sig))
		stop <- struct{}{}
	}()
}
