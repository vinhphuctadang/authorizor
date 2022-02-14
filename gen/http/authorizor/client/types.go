// Code generated by goa v3.5.5, DO NOT EDIT.
//
// authorizor HTTP client types
//
// Command:
// $ goa gen github.com/vinhphuctadang/authorizor/design

package client

import (
	authorizor "github.com/vinhphuctadang/authorizor/gen/authorizor"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "authorizor" service "register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// username
	Username string `form:"username" json:"username" xml:"username"`
	// password
	Password string `form:"password" json:"password" xml:"password"`
}

// RegisterResponseBody is the type of the "authorizor" service "register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// return code of registration
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// jwt refresh token
	RefreshToken *string `form:"refreshToken,omitempty" json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// jwt access token
	AccessToken *string `form:"accessToken,omitempty" json:"accessToken,omitempty" xml:"accessToken,omitempty"`
}

// PingResponseBody is the type of the "authorizor" service "ping" endpoint
// HTTP response body.
type PingResponseBody struct {
	// Return 'ok' if server is running
	Health *string `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "register" endpoint of the "authorizor" service.
func NewRegisterRequestBody(p *authorizor.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{
		Username: p.Username,
		Password: p.Password,
	}
	return body
}

// NewRegisterResultOK builds a "authorizor" service "register" endpoint result
// from a HTTP "OK" response.
func NewRegisterResultOK(body *RegisterResponseBody) *authorizor.RegisterResult {
	v := &authorizor.RegisterResult{
		Code:         *body.Code,
		RefreshToken: *body.RefreshToken,
		AccessToken:  *body.AccessToken,
	}

	return v
}

// NewPingResultOK builds a "authorizor" service "ping" endpoint result from a
// HTTP "OK" response.
func NewPingResultOK(body *PingResponseBody) *authorizor.PingResult {
	v := &authorizor.PingResult{
		Health: *body.Health,
	}

	return v
}

// ValidateRegisterResponseBody runs the validations defined on
// RegisterResponseBody
func ValidateRegisterResponseBody(body *RegisterResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refreshToken", "body"))
	}
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("accessToken", "body"))
	}
	return
}

// ValidatePingResponseBody runs the validations defined on PingResponseBody
func ValidatePingResponseBody(body *PingResponseBody) (err error) {
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	return
}
