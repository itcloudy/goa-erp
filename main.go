//go:generate goagen bootstrap -d goa-erp/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"goa-erp/app"
)

func main() {
	// Create service
	service := goa.New("goa-erp")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "company" controller
	c2 := NewCompanyController(service)
	app.MountCompanyController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
