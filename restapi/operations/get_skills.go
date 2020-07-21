// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSkillsHandlerFunc turns a function with the right signature into a get skills handler
type GetSkillsHandlerFunc func(GetSkillsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSkillsHandlerFunc) Handle(params GetSkillsParams) middleware.Responder {
	return fn(params)
}

// GetSkillsHandler interface for that can handle valid get skills params
type GetSkillsHandler interface {
	Handle(GetSkillsParams) middleware.Responder
}

// NewGetSkills creates a new http.Handler for the get skills operation
func NewGetSkills(ctx *middleware.Context, handler GetSkillsHandler) *GetSkills {
	return &GetSkills{Context: ctx, Handler: handler}
}

/*GetSkills swagger:route GET /skills getSkills

GetSkills get skills API

*/
type GetSkills struct {
	Context *middleware.Context
	Handler GetSkillsHandler
}

func (o *GetSkills) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSkillsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}