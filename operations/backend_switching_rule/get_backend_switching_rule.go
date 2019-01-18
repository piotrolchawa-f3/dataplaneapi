// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBackendSwitchingRuleHandlerFunc turns a function with the right signature into a get backend switching rule handler
type GetBackendSwitchingRuleHandlerFunc func(GetBackendSwitchingRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBackendSwitchingRuleHandlerFunc) Handle(params GetBackendSwitchingRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetBackendSwitchingRuleHandler interface for that can handle valid get backend switching rule params
type GetBackendSwitchingRuleHandler interface {
	Handle(GetBackendSwitchingRuleParams, interface{}) middleware.Responder
}

// NewGetBackendSwitchingRule creates a new http.Handler for the get backend switching rule operation
func NewGetBackendSwitchingRule(ctx *middleware.Context, handler GetBackendSwitchingRuleHandler) *GetBackendSwitchingRule {
	return &GetBackendSwitchingRule{Context: ctx, Handler: handler}
}

/*GetBackendSwitchingRule swagger:route GET /services/haproxy/configuration/backend_switching_rules/{id} BackendSwitchingRule getBackendSwitchingRule

Return one Backend Switching Rule

Returns one Backend Switching Rule configuration by it's ID in the specified frontend.

*/
type GetBackendSwitchingRule struct {
	Context *middleware.Context
	Handler GetBackendSwitchingRuleHandler
}

func (o *GetBackendSwitchingRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBackendSwitchingRuleParams()

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