package handlers

import (
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getHealthHandler struct {
}

func NewGetHealthHandler() *getHealthHandler{
	return &getHealthHandler{}
}

func (h *getHealthHandler) Handle(params operations.GetHealthParams) middleware.Responder {
	logrus.Debug("getHealthHandler got called")
	return operations.NewGetHealthOK()
}