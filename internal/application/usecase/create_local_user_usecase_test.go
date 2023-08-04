package usecase

import (
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/validation"
	"testing"
)

func TestCreateLocalUserUseCase(t *testing.T) {
	t.Run("it should not send invalid users to the repository", func(t *testing.T) {
		validUser := entity.NewUser(1, "amazing user", "example@email.io", "", "secret")
		invalidUser := entity.NewUser(0, "invalid user", "invalid_email", "", "secret")

		mock := mockUserRepository{
			[]*entity.User{},
		}

		usecase := NewCreateLocalUserUseCase(&mock)

		expectedErr := errors.New(string(validation.InvalidEmail))

		if err := usecase.Execute(invalidUser.Name, invalidUser.Email, invalidUser.Password); err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("got %s want %s", err, expectedErr)
		}

		if err := usecase.Execute(validUser.Name, validUser.Email, validUser.Password); err != nil {
			t.Errorf("got %s want %s", err, expectedErr)
		}

		for _, user := range mock.Users {
			if user.Name == "wrong user" {
				b, _ := json.Marshal(user)

				t.Errorf("got %s want %s", b, "any user")
			}
		}
	})
}
