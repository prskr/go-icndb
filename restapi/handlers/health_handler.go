package handlers

import (
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type getHealthHandler struct {

}

func NewGetHealthHandler() *getHealthHandler{
	return &getHealthHandler{}
}

func (h *getHealthHandler) Handle(params operations.GetHealthParams) middleware.Responder {
	return operations.NewGetHealthOK()
}