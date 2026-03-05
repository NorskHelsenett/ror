package trace

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc/credentials"
)

var (
	sleepTime    = time.Minute
	GrpcEndpoint = "localhost:4317"
)

func initTracerProvider(ctx context.Context, serviceName string, grpcEndpoint string) (func(context.Context) error, error) {
	if grpcEndpoint != "" {
		GrpcEndpoint = grpcEndpoint
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	exporterOpts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(GrpcEndpoint),
	}
	if strings.HasSuffix(GrpcEndpoint, ":443") {
		exporterOpts = append(exporterOpts, otlptracegrpc.WithTLSCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		exporterOpts = append(exporterOpts, otlptracegrpc.WithInsecure())
	}

	traceExporter, err := otlptracegrpc.New(ctx, exporterOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1))),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

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

func StartTracing(stop chan struct{}, cancelChan chan os.Signal, serviceName string, grpcEndpoint string) {
	go func() {
		ConnectTracer(stop, serviceName, grpcEndpoint)
		sig := <-cancelChan
		rlog.Info("Caught signal", rlog.Any("signal", sig))
		stop <- struct{}{}
	}()
}
