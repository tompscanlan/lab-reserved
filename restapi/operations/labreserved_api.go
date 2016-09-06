package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewLabreservedAPI creates a new Labreserved instance
func NewLabreservedAPI(spec *loads.Document) *LabreservedAPI {
	o := &LabreservedAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*LabreservedAPI REST interface for lab items reserved by a person until a date */
type LabreservedAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GetItemNameHandler sets the operation handler for the get item name operation
	GetItemNameHandler GetItemNameHandler
	// GetItemsHandler sets the operation handler for the get items operation
	GetItemsHandler GetItemsHandler
	// GetUsersHandler sets the operation handler for the get users operation
	GetUsersHandler GetUsersHandler
	// PostItemHandler sets the operation handler for the post item operation
	PostItemHandler PostItemHandler
	// PostItemNameReservationHandler sets the operation handler for the post item name reservation operation
	PostItemNameReservationHandler PostItemNameReservationHandler
	// PostUserHandler sets the operation handler for the post user operation
	PostUserHandler PostUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *LabreservedAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LabreservedAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *LabreservedAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LabreservedAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LabreservedAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LabreservedAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LabreservedAPI
func (o *LabreservedAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetItemNameHandler == nil {
		unregistered = append(unregistered, "GetItemNameHandler")
	}

	if o.GetItemsHandler == nil {
		unregistered = append(unregistered, "GetItemsHandler")
	}

	if o.GetUsersHandler == nil {
		unregistered = append(unregistered, "GetUsersHandler")
	}

	if o.PostItemHandler == nil {
		unregistered = append(unregistered, "PostItemHandler")
	}

	if o.PostItemNameReservationHandler == nil {
		unregistered = append(unregistered, "PostItemNameReservationHandler")
	}

	if o.PostUserHandler == nil {
		unregistered = append(unregistered, "PostUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LabreservedAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LabreservedAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *LabreservedAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *LabreservedAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *LabreservedAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *LabreservedAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/item/{name}"] = NewGetItemName(o.context, o.GetItemNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items"] = NewGetItems(o.context, o.GetItemsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = NewGetUsers(o.context, o.GetUsersHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/item"] = NewPostItem(o.context, o.PostItemHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/item/{name}/reservation"] = NewPostItemNameReservation(o.context, o.PostItemNameReservationHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = NewPostUser(o.context, o.PostUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LabreservedAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
