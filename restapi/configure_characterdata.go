// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"character-data-microservice/restapi/operations"
)

//go:generate swagger generate server --target ../../character-data-microservice --name Characterdata --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.CharacterdataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CharacterdataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GetClassHandler == nil {
		api.GetClassHandler = operations.GetClassHandlerFunc(func(params operations.GetClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetClass has not yet been implemented")
		})
	}
	if api.GetItemHandler == nil {
		api.GetItemHandler = operations.GetItemHandlerFunc(func(params operations.GetItemParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetItem has not yet been implemented")
		})
	}
	if api.GetItemsHandler == nil {
		api.GetItemsHandler = operations.GetItemsHandlerFunc(func(params operations.GetItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetItems has not yet been implemented")
		})
	}
	if api.GetSkillHandler == nil {
		api.GetSkillHandler = operations.GetSkillHandlerFunc(func(params operations.GetSkillParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSkill has not yet been implemented")
		})
	}
	if api.GetSkillsHandler == nil {
		api.GetSkillsHandler = operations.GetSkillsHandlerFunc(func(params operations.GetSkillsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSkills has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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