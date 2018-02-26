package main

import (
	"github.com/goadesign/goa"
	"goa-erp/app"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here

	res := app.AccountCollection{}
	return ctx.OK(res)
	// AccountController_List: end_implement
}
