package usecase

import (
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"testing"
)

func TestCreateLocalUserUseCase(t *testing.T) {
	t.Run("it should not send invalid users to the repository", func(t *testing.T) {
		validUser := dto.CreateUserLocalInput{
			Name:     "Valid User",
			Email:    "valid@email.com",
			Password: "password",
		}

		invalidUser := dto.CreateUserLocalInput{
			Name:     "Invalid User",
			Email:    "its_not_a_email",
			Password: "password",
		}

		mockRepository := mock.NewMockUserRepository()

		mockHelper := mock.NewMockPasswordHelper("#4asd3", true)

		uc := NewCreateLocalUserUseCase(mockRepository, mockHelper)

		if err := uc.Execute(&invalidUser); err == nil || !errors.Is(err, domain.InvalidEmailError) {
			t.Errorf("got %s want %s", err, domain.InvalidEmailError)
		}

		if err := uc.Execute(&validUser); err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		for _, user := range mockRepository.Users {
			if user.Name == "invalid user" {
				b, _ := json.Marshal(user)

				t.Errorf("got %s want %s", b, "any user")
			}
		}
	})

	t.Run("it should encrypt user password before pass to repository", func(t *testing.T) {
		userInput := dto.CreateUserLocalInput{
			Name:     "Archimedes",
			Email:    "archimedes@gopher.io",
			Password: "eureka",
		}

		mockRepository := mock.NewMockUserRepository()

		mockHelper := mock.NewMockPasswordHelper("#4asd3", true)

		uc := NewCreateLocalUserUseCase(mockRepository, mockHelper)

		if err := uc.Execute(&userInput); err != nil {
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

	t.Run("it should return error UserAlreadyExistError if user already exist", func(t *testing.T) {
		mockRepository := mock.NewMockUserRepository()

		email := "mouse@email.com"

		mockRepository.Users = []*entity.User{
			entity.NewUser(3, "Mouse", email, "", "password"),
		}

		mockHelper := mock.NewMockPasswordHelper("#4asd3", true)

		userInput := dto.CreateUserLocalInput{
			Name:     "Another Mouse",
			Email:    email,
			Password: "password",
		}

		uc := NewCreateLocalUserUseCase(mockRepository, mockHelper)

		uc.Execute(&userInput)
	})
}
