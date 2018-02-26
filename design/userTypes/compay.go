package userTypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CompanyPayload user types
var CompanyPayload = Type("CompanyPayload", func() {
	Attribute("name", String)
})
