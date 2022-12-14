// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostTriggerHandlerFunc turns a function with the right signature into a post trigger handler
type PostTriggerHandlerFunc func(PostTriggerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTriggerHandlerFunc) Handle(params PostTriggerParams) middleware.Responder {
	return fn(params)
}

// PostTriggerHandler interface for that can handle valid post trigger params
type PostTriggerHandler interface {
	Handle(PostTriggerParams) middleware.Responder
}

// NewPostTrigger creates a new http.Handler for the post trigger operation
func NewPostTrigger(ctx *middleware.Context, handler PostTriggerHandler) *PostTrigger {
	return &PostTrigger{Context: ctx, Handler: handler}
}

/*
	PostTrigger swagger:route POST /trigger trigger postTrigger

Save new achievement
*/
type PostTrigger struct {
	Context *middleware.Context
	Handler PostTriggerHandler
}

func (o *PostTrigger) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostTriggerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
