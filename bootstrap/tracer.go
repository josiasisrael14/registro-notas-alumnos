package bootstrap

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	"notas/model"
)

func initTracer(ctx context.Context, applicationName string) *trace.TracerProvider {
	endpoint := os.Getenv("ENDPOINT_COLLECTOR_OTLP")
	if endpoint == "" {
		log.Fatalf("the env ENDPOINT_COLLECTOR_OTLP could not be empty")
	}

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(endpoint),
		),
	)
	if err != nil {
		log.Printf("could not new otlptrace: %v", err)
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", applicationName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("could not set resources: %v", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
		trace.WithResource(resources),
	)

	otel.SetTracerProvider(traceProvider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return traceProvider
}

func shutdownTraceProvider(ctx context.Context, logger model.Logger, traceProvider *trace.TracerProvider) {
	if err := traceProvider.Shutdown(ctx); err != nil {
		logger.Errorf("error occurred in shutdown trace provider, err: %v", err)
	}
}
