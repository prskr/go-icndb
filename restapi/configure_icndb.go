package restapi

import (
	"crypto/tls"
	"github.com/baez90/go-icndb/restapi/handlers"
	"github.com/gobuffalo/packr"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/baez90/go-icndb/internal/pkg/models"

	"github.com/baez90/go-icndb/restapi/operations"
)

//go:generate swagger generate server --target ./ --name ICNDB --spec ./assets/api/swagger.yml

func configureFlags(api *operations.ICNDBAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ICNDBAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	box :=packr.NewBox("../assets/app")
	jokes, err := models.LoadFacts(&box, "jokes.json")

	if err != nil {
		log.Println(err)
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetJokeByIDHandler = operations.GetJokeByIDHandlerFunc(func(params operations.GetJokeByIDParams) middleware.Responder {
		return handlers.NewJokesByIDHandler(jokes).Handle(params)
	})
	api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
		return handlers.NewGetHealthHandler().Handle(params)
	})
	api.GetJokesCountHandler = operations.GetJokesCountHandlerFunc(func(params operations.GetJokesCountParams) middleware.Responder {
		return handlers.NewGetJokesCountHandler(jokes).Handle(params)
	})
	api.GetRandomJokeHandler = operations.GetRandomJokeHandlerFunc(func(params operations.GetRandomJokeParams) middleware.Responder {
		return handlers.NewRandomJokeHandler(jokes).Handle(params)
	})

	api.GetHostnameHandler = operations.GetHostnameHandlerFunc(func(params operations.GetHostnameParams) middleware.Responder {
		return handlers.NewGetHostnameHandler().Handle(params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}