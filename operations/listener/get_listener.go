// Code generated by go-swagger; DO NOT EDIT.

package listener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetListenerHandlerFunc turns a function with the right signature into a get listener handler
type GetListenerHandlerFunc func(GetListenerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetListenerHandlerFunc) Handle(params GetListenerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetListenerHandler interface for that can handle valid get listener params
type GetListenerHandler interface {
	Handle(GetListenerParams, interface{}) middleware.Responder
}

// NewGetListener creates a new http.Handler for the get listener operation
func NewGetListener(ctx *middleware.Context, handler GetListenerHandler) *GetListener {
	return &GetListener{Context: ctx, Handler: handler}
}

/*GetListener swagger:route GET /services/haproxy/configuration/listeners/{name} Listener getListener

Return one listener

Returns one listener configuration by it's name in the specified frontend.

*/
type GetListener struct {
	Context *middleware.Context
	Handler GetListenerHandler
}

func (o *GetListener) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetListenerParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}