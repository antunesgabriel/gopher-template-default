package usecase

import (
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/validation"
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"testing"
)

func TestCreateLocalUserUseCase(t *testing.T) {
	t.Run("it should not send invalid users to the repository", func(t *testing.T) {
		validUser := entity.NewUser(1, "amazing user", "example@email.io", "", "secret")
		invalidUser := entity.NewUser(0, "invalid user", "invalid_email", "", "secret")

		mockRepository := mock.NewMockUserRepository()

		mockHelper := mock.NewMockPasswordHelper("#4asd3", true)

		uc := NewCreateLocalUserUseCase(mockRepository, mockHelper)

		expectedErr := errors.New(string(validation.InvalidEmail))

		if err := uc.Execute(invalidUser.Name, invalidUser.Email, invalidUser.Password); err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("got %s want %s", err, expectedErr)
		}

		if err := uc.Execute(validUser.Name, validUser.Email, validUser.Password); err != nil {
			t.Errorf("got %s want %s", err, expectedErr)
		}

		for _, user := range mockRepository.Users {
			if user.Name == "invalid user" {
				b, _ := json.Marshal(user)

				t.Errorf("got %s want %s", b, "any user")
			}
		}
	})

	t.Run("it should encrypt user password before pass to repository", func(t *testing.T) {
		userInput := entity.NewUser(7, "Archimedes", "archimedes@gopher.io", "", "eureka")

		mockRepository := mock.NewMockUserRepository()

		mockHelper := mock.NewMockPasswordHelper("#4asd3", true)

		uc := NewCreateLocalUserUseCase(mockRepository, mockHelper)

		if err := uc.Execute(userInput.Name, userInput.Email, userInput.Password); err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		var storedUser *entity.User

		for _, user := range mockRepository.Users {
			if user.Name == "Archimedes" {
				storedUser = user
			}
		}

		if storedUser == nil {
			t.Errorf("got %s want %s", "user don't stored", "user stored")
		}

		if storedUser.Password != string(mockHelper.Hashed) {
			t.Errorf("got %s want %s", storedUser.Password, string(mockHelper.Hashed))
		}
	})
}
