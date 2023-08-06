package infra

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"testing"
	"time"
)

func TestChiJWTHelper_Encode(t *testing.T) {
	t.Run("it should generate token with correct claims", func(t *testing.T) {
		claimsInput := map[string]interface{}{
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			"iat": time.Now().Unix(),
			"id":  42,
		}

		helper := NewChiJWTHelper("vasco")

		token, err := helper.Encode(claimsInput)

		if err != nil {
			t.Errorf("got %s want %s", err, "no expected error")
		}

		if token == "" {
			t.Errorf("got %s want %s", token, "jwt token")
		}
	})
}

func TestChiJWTHelper_FromContext(t *testing.T) {
	t.Run("it should return user logged id", func(t *testing.T) {
		id := 42

		claimsInput := map[string]interface{}{
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			"iat": time.Now().Unix(),
			"id":  id,
		}

		helper := NewChiJWTHelper("vasco")

		tokenString, err := helper.Encode(claimsInput)

		if err != nil {
			t.Errorf("got %s want %s", err, "no expected error")
		}

		if tokenString == "" {
			t.Errorf("got %s want %s", tokenString, "jwt token")
		}

		token, err := jwtauth.VerifyToken(helper.tokenAuth, tokenString)

		ctx := jwtauth.NewContext(context.Background(), token, nil)

		userId, err := helper.GetUserID(ctx)

		if err != nil {
			t.Errorf("got %s want %s", err, "no expected error")
		}

		if userId != id {
			t.Errorf("got %d want %d", userId, id)
		}
	})
}
