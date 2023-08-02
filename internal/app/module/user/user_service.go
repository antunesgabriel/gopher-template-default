package user

import "context"

type UserService struct {
	repository UserRepository
}

func NewUserService(rep UserRepository) *UserService {
	s := UserService{
		repository: rep,
	}

	return &s
}

func (s *UserService) CreateLocal(name, email, password string) error {
	ctx := context.Background()

	u := New(0, name, email, "", password)

	err := u.Validate()

	if err != nil {
		return err
	}

	err = s.repository.Create(ctx, u)

	return err
}
