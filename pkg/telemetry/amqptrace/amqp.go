package amqptrace

import (
	"context"

	"go.opentelemetry.io/otel"
)

// AmqpHeadersCarrier adapts an AMQP header map for OpenTelemetry context propagation.
type AmqpHeadersCarrier map[string]interface{}

func (a AmqpHeadersCarrier) Get(key string) string {
	v, ok := a[key]
	if !ok {
		return ""
	}
	s, _ := v.(string)
	return s
}

func (a AmqpHeadersCarrier) Set(key string, value string) {
	a[key] = value
}

func (a AmqpHeadersCarrier) Keys() []string {
	r := make([]string, 0, len(a))
	for k := range a {
		r = append(r, k)
	}
	return r
}

// InjectAMQPHeaders injects the trace context from ctx into a new header map
// suitable for publishing an AMQP message.
func InjectAMQPHeaders(ctx context.Context) map[string]interface{} {
	h := make(AmqpHeadersCarrier)
	otel.GetTextMapPropagator().Inject(ctx, h)
	return h
}

// ExtractAMQPHeaders extracts trace context from AMQP headers into the
// returned context.
func ExtractAMQPHeaders(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, AmqpHeadersCarrier(headers))
}
