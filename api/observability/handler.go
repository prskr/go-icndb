package observability

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(router chi.Router) {
	handler := Handler{}

	router.Get("/health", handler.Health)
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
}

type Handler struct {
}

func (h Handler) Health(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
