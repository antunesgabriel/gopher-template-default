package user

import "context"

// Output

type UserRepository interface {
	FindUserByID(ctx context.Context, id int64) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, u *User) error
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
