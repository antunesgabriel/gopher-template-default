package user

import "context"

type UserService struct {
	repository UserRepository
}

func NewCreateUserUseCase(rep UserRepository) *UserService {
	s := UserService{
		repository: rep,
	}

	return &s
}

func (s *UserService) CreateLocal(ctx context.Context, name, email, password string) error {
	u := New(0, name, email, "", password)

	err := u.Validate()

	if err != nil {
		return err
	}

	err = s.repository.Create(ctx, u)

	return err
}
