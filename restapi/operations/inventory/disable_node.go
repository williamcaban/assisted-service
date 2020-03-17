// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DisableNodeHandlerFunc turns a function with the right signature into a disable node handler
type DisableNodeHandlerFunc func(DisableNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DisableNodeHandlerFunc) Handle(params DisableNodeParams) middleware.Responder {
	return fn(params)
}

// DisableNodeHandler interface for that can handle valid disable node params
type DisableNodeHandler interface {
	Handle(DisableNodeParams) middleware.Responder
}

// NewDisableNode creates a new http.Handler for the disable node operation
func NewDisableNode(ctx *middleware.Context, handler DisableNodeHandler) *DisableNode {
	return &DisableNode{Context: ctx, Handler: handler}
}

/*DisableNode swagger:route DELETE /nodes/{node_id}/actions/enable inventory disableNode

Disable a node for use

*/
type DisableNode struct {
	Context *middleware.Context
	Handler DisableNodeHandler
}

func (o *DisableNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDisableNodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
