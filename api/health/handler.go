package health

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var _ httprouter.Handle = Handler{}.Health

func SetupRouter(r *httprouter.Router) {
	handler := Handler{}

	r.GET("/health", handler.Health)
}

type Handler struct {
}

func (h Handler) Health(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
}
