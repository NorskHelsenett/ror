package telemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
)

func SetupOTel(ctx context.Context, opts ...Option) (shutdown func(context.Context) error, err error) {
	var shutdownFunctions []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFunctions {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFunctions = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	res := newResource(ctx, opts...)

	tracerProvider, err := newTracerProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFunctions = append(shutdownFunctions, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider(res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFunctions = append(shutdownFunctions, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	loggerProvider, err := newLoggerProvider(res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFunctions = append(shutdownFunctions, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return
}
