package userTypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AccountPayload user types
var AccountPayload = Type("AccountPayload", func() {
	Attribute("username", String, func() {
		MinLength(6)
		Example("cloudy")
	})
})
