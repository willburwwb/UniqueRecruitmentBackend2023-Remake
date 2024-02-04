package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	Tracer trace.Tracer = otel.Tracer("UniqueRecruitmentBackend2023DefaultTracer")
)

// SetupTracing - setup otel tracer. return shutdown function
func SetupTracing(appName, mode, reportBackground string) (func(ctx context.Context) error, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(reportBackground)))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in an Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
			attribute.String("environment", mode),
		)),
	)

	otel.SetTracerProvider(tp)
	Tracer = otel.GetTracerProvider().Tracer(appName)
	return func(ctx context.Context) error {
		return tp.Shutdown(ctx)
	}, nil
}
