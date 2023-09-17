package metrics

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func RequestMetrics(meter metric.Meter) func(innerHandler http.Handler) http.Handler {
	totalRequestCounter, err := meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
	)

	if err != nil {
		panic(err)
	}

	requestDurationHistogram, err := meter.Float64Histogram(
		"http_request_duration",
		metric.WithDescription("Duration of HTTP requests"),
		metric.WithUnit("ms"),
	)

	if err != nil {
		panic(err)
	}

	return func(innerHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			totalRequestCounter.Add(request.Context(), 1, metric.WithAttributes(
				attribute.String("method", request.Method),
				attribute.String("host", request.Host),
				attribute.String("url", request.RequestURI),
			))

			start := time.Now()
			defer func() {
				requestDuration := time.Since(start)
				requestDurationHistogram.Record(
					request.Context(),
					float64(requestDuration.Milliseconds()),
					metric.WithAttributes(
						attribute.String("method", request.Method),
						attribute.String("host", request.Host),
						attribute.String("url", request.RequestURI),
					))
			}()

			innerHandler.ServeHTTP(writer, request)
		})
	}
}
