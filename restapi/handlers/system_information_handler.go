package handlers

import (
	"github.com/baez90/go-icndb/models"
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"os"
)

type getHostnameHandler struct {

}

func NewGetHostnameHandler() *getHostnameHandler{
	return &getHostnameHandler{}
}

func (h *getHostnameHandler) Handle(params operations.GetHostnameParams) middleware.Responder {

	hostname, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
		return operations.NewGetHostnameInternalServerError()
	}

	return operations.NewGetHostnameOK().WithPayload(&models.HostnameResponse{
		Hostname: &hostname,
	})
}