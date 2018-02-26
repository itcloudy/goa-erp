package mediaTypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CompanyMediaType company
var CompanyMediaType = MediaType("application/vnd.company+json", func() {
	Description("system user table")
	Attributes(func() {
		Attribute("id", Integer, "ID of company", func() {
			Example(1)
		})
		Attribute("name", String, "name of company", func() {
			Example("弛菡")
		})
		Attribute("create_date", DateTime, "Date of create")
		Attribute("update_date", DateTime, "Date of update")
		Attribute("create_uid", Integer, "company id of create", func() {
			Example(1)
		})
		Attribute("update_uid", Integer, "company id of update", func() {
			Example(1)
		})
		Required("id", "name")
	})
	View("default", func() {
		Description("company detail")
		Attribute("id")
		Attribute("name")
		Attribute("create_date")
		Attribute("update_date")
		Attribute("create_uid")
		Attribute("update_uid")

	})
})
