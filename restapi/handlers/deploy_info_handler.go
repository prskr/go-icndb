package handlers

import (
	"github.com/baez90/go-icndb/models"
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"os"
)

type getDeployEnvHandler struct {
}

func NewGetDeployEnvHandler() *getDeployEnvHandler {
	return &getDeployEnvHandler{}
}

func (h *getDeployEnvHandler) Handle(params operations.GetDeployEnvParams) middleware.Responder {

	deployEnv := "production"

	if envVar := os.Getenv("DEPLOYMENT_ENV"); envVar != "" {
		deployEnv = envVar
	}

	return operations.NewGetDeployEnvOK().WithPayload(&models.DeployEnvResponse{
		Env: deployEnv,
	})
}
