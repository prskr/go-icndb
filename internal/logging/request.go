package logging

import (
	"log/slog"
	"net/http"
	"strings"
)

func RequestLogger(innerHandler http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		var headerValues []any

		for key, value := range request.Header {
			headerValues = append(headerValues, slog.String(key, strings.Join(value, ",")))
		}

		logger.Info(
			"Incoming request",
			slog.String("method", request.Method),
			slog.String("host", request.Host),
			slog.String("url", request.RequestURI),
			slog.String("client", request.RemoteAddr),
			slog.Group("headers", headerValues...),
		)

		innerHandler.ServeHTTP(writer, request)
	})
}
