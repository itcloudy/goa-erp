package main

import (
	"github.com/goadesign/goa"
	"goa-erp/app"
)

// CompanyController implements the company resource.
type CompanyController struct {
	*goa.Controller
}

// NewCompanyController creates a company controller.
func NewCompanyController(service *goa.Service) *CompanyController {
	return &CompanyController{Controller: service.NewController("CompanyController")}
}

// List runs the list action.
func (c *CompanyController) List(ctx *app.ListCompanyContext) error {
	// CompanyController_List: start_implement

	// Put your logic here

	res := app.CompanyCollection{}
	return ctx.OK(res)
	// CompanyController_List: end_implement
}
