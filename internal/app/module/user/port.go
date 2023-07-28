package user

import "context"

// Output

type UserRepository interface {
	FindUserByID(ctx context.Context, id string) (*user, error)
	FindUserByEmail(ctx context.Context, email string) (*user, error)
	Create(ctx context.Context, u *user) error
	Update(ctx context.Context, u *user) (*user, error)
	Delete(ctx context.Context, u *user) error
}

// Inputs

type createUserLocalDTO struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required, email"`
	Password string `json:"password" binding:"required"`
}

type createUserExternalDTO struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
