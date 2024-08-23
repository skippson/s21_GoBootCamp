// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"day04/ex01/server/restapi/operations"
)

type BuyCandyCreatedBody struct {
	Thanks string `json:"thanks"`
	Change int64  `json:"change"`
}

var candyPrices = map[string]int{
	"CE": 5,
	"AA": 3,
	"NT": 2,
	"DE": 4,
	"YR": 1,
}

//go:generate swagger generate server --target ../../src --name CandyServer --spec ../swagger.json --principal interface{}

func configureFlags(api *operations.CandyServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandyServerAPI) http.Handler {
	// configure the api here
	var buyCandyHandler operations.BuyCandyHandlerFunc = func(params operations.BuyCandyParams) middleware.Responder {
		reqBody := params.Order

		if *reqBody.CandyCount < 0 {
			err := errors.New(http.StatusBadRequest, "candyCount must be a positive integer")
			return middleware.Error(http.StatusBadRequest, err)
		}

		if _, ok := candyPrices[*reqBody.CandyType]; !ok {
			err := errors.New(http.StatusBadRequest, "invalid canduTYpe")
			return middleware.Error(http.StatusBadRequest, err)

		}

		candyPrice := int64(candyPrices[*reqBody.CandyType])
		totalCost := candyPrice * *reqBody.CandyCount

		if totalCost > *reqBody.Money {
			diff := totalCost - *reqBody.Money
			err := errors.New(http.StatusBadRequest, fmt.Sprintf("You need %d more money!", diff))
			return middleware.Error(http.StatusBadRequest, err)
		}

		change := int64(*reqBody.Money - totalCost)
		resp := operations.BuyCandyCreatedBody{
			Thanks: "Thank you!",
			Change: change,
		}
		response := operations.NewBuyCandyCreated()
		response.SetPayload(&resp)

		return response
	}
	api.BuyCandyHandler = buyCandyHandler
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BuyCandyHandler == nil {
		api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.BuyCandy has not yet been implemented")
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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
