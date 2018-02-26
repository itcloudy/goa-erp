package resources

import (
	mt "goa-erp/design/mediaTypes"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CompanyResource resource
var CompanyResource = Resource("company", func() {
	DefaultMedia(mt.CompanyMediaType)
	BasePath("/company")
	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("list all companys")
		Response(OK, CollectionOf(mt.CompanyMediaType))
	})
})
