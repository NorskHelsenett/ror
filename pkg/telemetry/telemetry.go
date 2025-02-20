package telemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func SetupOTel(ctx context.Context, opts ...Option) (shutdown func(context.Context) error, err error) {
	// we enable all signals by default
	cfg := config{
		WithLogger: true,
		WithMeter:  true,
		WithTracer: true,
	}
	cfg = applyEnvOptions(cfg)
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	res := newResource(ctx, opts...)

	if cfg.WithTracer {
		var tracerProvider *trace.TracerProvider
		tracerProvider, err = newTracerProvider(ctx, res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
		otel.SetTracerProvider(tracerProvider)
	}

	if cfg.WithMeter {
		var meterProvider *metric.MeterProvider
		meterProvider, err = newMeterProvider(res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
		otel.SetMeterProvider(meterProvider)
	}

	if cfg.WithLogger {
		var loggerProvider *log.LoggerProvider
		loggerProvider, err = newLoggerProvider(res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
		global.SetLoggerProvider(loggerProvider)
	}

	return
}
