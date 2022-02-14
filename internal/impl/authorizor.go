package impl

import (
	"context"
	"log"

	"github.com/vinhphuctadang/authorizor/gen/authorizor"
)

type authorizorService struct {
	logger *log.Logger
	authorizor.Service
}

func NewAuthorizor(logger *log.Logger) (authorizor.Service, error) {
	return &authorizorService{
		logger: logger,
	}, nil
}

func (svc *authorizorService) Register(context.Context, *authorizor.RegisterPayload) (res *authorizor.RegisterResult, err error) {

	return &authorizor.RegisterResult{
		Code: "c_ok",
	}, nil
}

func (svc *authorizorService) Ping(context.Context) (res *authorizor.PingResult, err error) {
	return &authorizor.PingResult{
		Health: "ok",
	}, nil
}
