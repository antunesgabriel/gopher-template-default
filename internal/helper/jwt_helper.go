package helper

import "context"

type JWTHelper interface {
	Encode(payload map[string]interface{}) (string, error)
	GetUserID(ctx context.Context) (int, error)
}
