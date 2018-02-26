package resources

import (
	mt "goa-erp/design/mediaTypes"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AccountResource resource
var AccountResource = Resource("account", func() {
	DefaultMedia(mt.AccountMediaType)
	BasePath("/account")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("list all accounts")
		Response(OK, CollectionOf(mt.AccountMediaType))
	})
})
