package mediaTypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Account is the system user
var AccountMediaType = MediaType("application/vnd.account+json", func() {
	Description("system user table")
	Attributes(func() {
		Attribute("id", Integer, "ID of account", func() {
			Example(1)
		})
		Attribute("username", String, "username of account", func() {
			Example("cloudy")
		})
		Attribute("password", String, "password of account", func() {
			Example("cloudy")
		})
		Attribute("email", String, "email of account", func() {
			Format("email")
			Example("272685110@qq.com")
		})
		Attribute("create_date", DateTime, "Date of create")
		Attribute("update_date", DateTime, "Date of update")
		Attribute("create_uid", Integer, "account id of create", func() {
			Example(1)
		})
		Attribute("update_uid", Integer, "account id of update", func() {
			Example(1)
		})
		Required("id", "username", "password", "email")
	})
	View("default", func() {
		Description("account detail")
		Attribute("id")
		Attribute("username")
		Attribute("email")
		Attribute("create_date")
		Attribute("update_date")
		Attribute("create_uid")
		Attribute("update_uid")

	})
})
