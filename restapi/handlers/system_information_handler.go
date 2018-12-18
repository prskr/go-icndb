package handlers

import (
	"github.com/baez90/go-icndb/models"
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net"
	"os"
)

type getHostnameHandler struct {
}

type getIPAddressesHandler struct {
}

func NewGetHostnameHandler() *getHostnameHandler {
	return &getHostnameHandler{}
}

func NewGetIPAddressesHandler() *getIPAddressesHandler {
	return &getIPAddressesHandler{}
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

func (h *getIPAddressesHandler) Handle(params operations.GetIPAddressesParams) middleware.Responder {

	addrs, err := net.InterfaceAddrs()

	// extract IP address as strings to return them
	if err == nil {
		addrStr := make([]string, len(addrs))
		for i, addr := range addrs {
			addrStr[i] = addr.String()
		}

		return operations.NewGetIPAddressesOK().WithPayload(&models.IPAddressesResponse{
			Addresses: addrStr,
		})
	}
	return operations.NewGetIPAddressesOK().WithPayload(&models.IPAddressesResponse{
		Addresses: make([]string, 0),
	})
}
