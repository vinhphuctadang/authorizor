package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("authorizor", func() {
	Title("The authorizor Service")
	Description("HTTP Service for managing identity")
	Server("identity", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var _ = Service("authorizor", func() {
	Description("Service for authenticating and authorizing")
	Method("register", func() {
		Payload(func() {
			Attribute("username", String, "username")
			Attribute("password", String, "password")
			Required("username", "password")
		})

		Result(Int)
		HTTP(func() {
			POST("/register")
			Response(StatusOK)
		})
	})

	Method("ping", func() {
		Payload(func() {})
		Result(String)
		HTTP(func() {
			GET("/ping")
		})
	})
})
