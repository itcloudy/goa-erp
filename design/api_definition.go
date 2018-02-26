package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa-erp", func() {
	Title("golang goa-erp use goa")
	Description("golang goa-erp use goa ")
	Contact(func() {
		Name("cloudy")
		Email("272685110@qq.com")
		URL("https://hexiaoyun128.github.io/cloudy-blog")
	})
	License(func() {
		Name("MIT")
		URL("https://hexiaoyun128.github.io/cloudy-blog")
	})
	Docs(func() {
		Description("golang")
		URL("https://hexiaoyun128.github.io/cloudy-blog")
	})
	Host("localhost:8081")
	Scheme("http")
	BasePath("/goa-erp")

	Origin("https://github.com/hexiaoyun128/cloudy-blog", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
