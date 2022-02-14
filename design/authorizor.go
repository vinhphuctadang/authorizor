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

		Result(func() {
			Field(1, "code", String, func() {
				Description("return code of registration")
				Example("c_success")
			})
			Field(2, "refreshToken", String, func() {
				Description("jwt refresh token")
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
			})
			Field(3, "accessToken", String, func() {
				Description("jwt access token")
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
			})
			Required("code", "refreshToken", "accessToken")
		})

		HTTP(func() {
			POST("/register")
			Response(StatusOK)
		})
	})

	Method("ping", func() {
		Payload(func() {})
		Result(func() {
			Field(1, "health", String, func() {
				Description("Return 'ok' if server is running")
			})
			Required("health")
		})
		HTTP(func() {
			GET("/ping")
		})
	})
})
