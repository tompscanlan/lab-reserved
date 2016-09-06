package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/tompscanlan/labreserved"
	"github.com/tompscanlan/labreserved/models"
	"github.com/tompscanlan/labreserved/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.LabreservedAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	labreserved.AddFlags(api)
}

func configureAPI(api *operations.LabreservedAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetItemsHandler = operations.GetItemsHandlerFunc(func(params operations.GetItemsParams) middleware.Responder {
		items := operations.NewGetItemsOK()
		items.SetPayload(labreserved.AllItems)
		return items
	})
	api.PostItemHandler = operations.PostItemHandlerFunc(func(params operations.PostItemParams) middleware.Responder {
		labreserved.AllItems[*params.Additem.Name] = *params.Additem
		item := operations.NewPostItemOK()
		i, ok := labreserved.AllItems[*params.Additem.Name]
		if ok {
			item.SetPayload(&i)
			return item
		} else {
			err := operations.NewPostItemBadRequest()
			err.SetPayload("failed to add and/or find item to lab map")
			return err
		}

	})
	api.GetUsersHandler = operations.GetUsersHandlerFunc(func(params operations.GetUsersParams) middleware.Responder {
		users := operations.NewGetUsersOK()
		users.SetPayload(labreserved.AllUsers)
		return users
	})
	api.PostUserHandler = operations.PostUserHandlerFunc(func(params operations.PostUserParams) middleware.Responder {
		labreserved.AllUsers[*params.Adduser.Name] = *params.Adduser
		user := operations.NewPostUserOK()
		u, ok := labreserved.AllUsers[*params.Adduser.Name]
		if ok {
			user.SetPayload(&u)
			return user
		} else {
			err := operations.NewPostUserBadRequest()
			err.SetPayload("failed to add and/or find user to lab map")
			return err
		}
	})
	api.GetItemNameHandler = operations.GetItemNameHandlerFunc(func(params operations.GetItemNameParams) middleware.Responder {
		resp := operations.NewGetItemNameOK()
		item := labreserved.AllItems[params.Name]
		resp.SetPayload(&item)
		return resp
	})
	api.PostItemNameReservationHandler = operations.PostItemNameReservationHandlerFunc(func(params operations.PostItemNameReservationParams) middleware.Responder {
		item, ok := labreserved.AllItems[params.Name]
		if !ok {
			err := operations.NewPostItemNameReservationBadRequest()
			err.SetPayload(fmt.Sprintf("failed to find item named %s. Can't add reservation.", params.Name))
			return err
		}

		item.Reserve(*params.Reservation.Username, models.StrfmtDateTimeToTime(params.Reservation.Begin), int(*params.Reservation.Hoursheld))
		labreserved.AllItems[params.Name] = item
		resp := operations.NewPostItemNameReservationOK()
		resp.SetPayload(params.Reservation)
		return resp

		//		if err != nil {
		//			err := operations.NewPostItemNameReservationBadRequest()
		//			err.SetPayload(err)
		//			return err
		//		}
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
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
