package usecase

import (
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"testing"
)

func TestLocalAuthUseCase_Execute(t *testing.T) {
	t.Run("it should return error UserIsNotRegister user no exist", func(t *testing.T) {
		buildLocalAuthUseCase("xpto", true, 1, "rrrrr")
	})
}

func buildLocalAuthUseCase(fakeHash string, isEqual bool, id int, token string) *LocalAuthUseCase {

	mockRepository := mock.NewMockUserRepository()
	passwordHelper := mock.NewMockPasswordHelper(fakeHash, isEqual)
	jwtHelper := mock.NewMockJWTHelper(id, token)

	uc := LocalAuthUseCase{
		repository:     mockRepository,
		passwordHelper: passwordHelper,
		jwtHelper:      jwtHelper,
	}

	return &uc
}
