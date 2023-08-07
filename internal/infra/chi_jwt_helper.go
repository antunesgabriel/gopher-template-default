package infra

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"github.com/go-chi/jwtauth/v5"
)

type ChiJWTHelper struct {
	tokenAuth *jwtauth.JWTAuth
}

func NewChiJWTHelper(signKey config.SignKey) *ChiJWTHelper {
	tokenAuth := jwtauth.New("HS256", []byte(signKey), nil)

	h := ChiJWTHelper{
		tokenAuth: tokenAuth,
	}

	return &h
}

func (it ChiJWTHelper) Encode(payload map[string]interface{}) (string, error) {
	_, token, err := it.tokenAuth.Encode(payload)

	return token, err
}

func (it ChiJWTHelper) GetUserID(ctx context.Context) (int, error) {
	_, claims, err := jwtauth.FromContext(ctx)

	value := claims["id"].(float64)

	return int(value), err
}
