package authorizorapi

import (
	"context"
	"log"

	authorizor "github.com/vinhphuctadang/authorizor/gen/authorizor"
)

// authorizor service example implementation.
// The example methods log the requests and return zero values.
type authorizorsrvc struct {
	logger *log.Logger
}

// NewAuthorizor returns the authorizor service implementation.
func NewAuthorizor(logger *log.Logger) authorizor.Service {
	return &authorizorsrvc{logger}
}

// Register implements register.
func (s *authorizorsrvc) Register(ctx context.Context, p *authorizor.RegisterPayload) (res int, err error) {
	s.logger.Print("authorizor.register")
	return
}

// Ping implements ping.
func (s *authorizorsrvc) Ping(ctx context.Context) (res string, err error) {
	s.logger.Print("authorizor.ping")
	return "ok", nil
}
