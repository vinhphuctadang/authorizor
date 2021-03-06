// Code generated by goa v3.5.5, DO NOT EDIT.
//
// authorizor HTTP server types
//
// Command:
// $ goa gen github.com/vinhphuctadang/authorizor/design

package server

import (
	authorizor "github.com/vinhphuctadang/authorizor/gen/authorizor"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "authorizor" service "register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// RegisterResponseBody is the type of the "authorizor" service "register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// return code of registration
	Code string `form:"code" json:"code" xml:"code"`
	// jwt refresh token
	RefreshToken string `form:"refreshToken" json:"refreshToken" xml:"refreshToken"`
	// jwt access token
	AccessToken string `form:"accessToken" json:"accessToken" xml:"accessToken"`
}

// PingResponseBody is the type of the "authorizor" service "ping" endpoint
// HTTP response body.
type PingResponseBody struct {
	// Return 'ok' if server is running
	Health string `form:"health" json:"health" xml:"health"`
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "authorizor" service.
func NewRegisterResponseBody(res *authorizor.RegisterResult) *RegisterResponseBody {
	body := &RegisterResponseBody{
		Code:         res.Code,
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}
	return body
}

// NewPingResponseBody builds the HTTP response body from the result of the
// "ping" endpoint of the "authorizor" service.
func NewPingResponseBody(res *authorizor.PingResult) *PingResponseBody {
	body := &PingResponseBody{
		Health: res.Health,
	}
	return body
}

// NewRegisterPayload builds a authorizor service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *authorizor.RegisterPayload {
	v := &authorizor.RegisterPayload{
		Username: *body.Username,
		Password: *body.Password,
	}

	return v
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}
