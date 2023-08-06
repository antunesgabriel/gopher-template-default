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

func NewUser(id int64, name, email, provider, password string) *User {
	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Provider: provider,
		Password: password,
	}

	return &u
}

func (it *User) validateRequiredFields() error {
	if it.Name == "" {
		return errors.New(string(validation.NameIsRequired))
	}

	if it.Email == "" {
		return errors.New(string(validation.EmailIsRequired))
	}

	addr, err := mail.ParseAddress(it.Email)

	if err != nil {
		return errors.New(string(validation.InvalidEmail))
	}

	it.Email = addr.Address

	return nil
}

func (it *User) ValidateNewLocalUser() error {
	if err := it.validateRequiredFields(); err != nil {
		return err
	}

	if it.Password == "" {
		return errors.New(string(validation.PasswordIsRequired))
	}

	return nil
}

func (it *User) ValidateNewExternalUser() error {
	if err := it.validateRequiredFields(); err != nil {
		return err
	}

	if it.Provider == "" {
		return errors.New(string(validation.ProviderIsRequired))
	}

	return nil
}

func (it *User) ChangePassword(password string) {
	it.Password = password
}
