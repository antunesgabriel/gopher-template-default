package entity

import (
	"errors"
	"net/mail"
	"time"

	"github.com/antunesgabriel/gopher-template-default/internal/domain/validation"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Provider  string     `json:"provider,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func New(id int64, name, email, provider, password string) *User {
	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Provider: provider,
		Password: password,
	}

	return &u
}

func (u *User) CheckIfNewUserIsValid() error {
	if u.Name == "" {
		return errors.New(string(validation.NameIsRequired))
	}

	if u.Email == "" {
		return errors.New(string(validation.EmailIsRequired))
	}

	if u.Provider == "" && u.Password == "" {
		return errors.New(string(validation.PasswordOrProviderIsRequired))
	}

	addr, err := mail.ParseAddress(u.Email)

	if err != nil {
		return errors.New(string(validation.InvalidEmail))
	}

	u.Email = addr.Address

	return nil
}
