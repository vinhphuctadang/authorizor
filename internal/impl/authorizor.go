package impl

import (
	"context"
	"net/http"

	"github.com/vinhphuctadang/authorizor/gen/authorizor"
)

type AuthorizorServiceImpl struct {
	authorizor.Service
}

func (svc *AuthorizorServiceImpl) Register(context.Context, *authorizor.RegisterPayload) (res int, err error) {

	return http.StatusOK, nil
}

func (svc *AuthorizorServiceImpl) Ping(context.Context) (res string, err error) {
	return "ok", nil
}
