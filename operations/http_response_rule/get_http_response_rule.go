// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetHTTPResponseRuleHandlerFunc turns a function with the right signature into a get HTTP response rule handler
type GetHTTPResponseRuleHandlerFunc func(GetHTTPResponseRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHTTPResponseRuleHandlerFunc) Handle(params GetHTTPResponseRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetHTTPResponseRuleHandler interface for that can handle valid get HTTP response rule params
type GetHTTPResponseRuleHandler interface {
	Handle(GetHTTPResponseRuleParams, interface{}) middleware.Responder
}

// NewGetHTTPResponseRule creates a new http.Handler for the get HTTP response rule operation
func NewGetHTTPResponseRule(ctx *middleware.Context, handler GetHTTPResponseRuleHandler) *GetHTTPResponseRule {
	return &GetHTTPResponseRule{Context: ctx, Handler: handler}
}

/*GetHTTPResponseRule swagger:route GET /services/haproxy/configuration/http_response_rules/{id} HTTPResponseRule getHttpResponseRule

Return one HTTP Response Rule

Returns one HTTP Response Rule configuration by it's ID in the specified parent.

*/
type GetHTTPResponseRule struct {
	Context *middleware.Context
	Handler GetHTTPResponseRuleHandler
}

func (o *GetHTTPResponseRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHTTPResponseRuleParams()

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