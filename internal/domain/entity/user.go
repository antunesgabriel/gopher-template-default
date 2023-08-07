package entity

import (
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"net/mail"
	"time"
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
		return domain.NameIsRequired
	}

	if it.Email == "" {
		return domain.EmailIsRequired
	}

	addr, err := mail.ParseAddress(it.Email)

	if err != nil {
		return domain.InvalidEmail
	}

	it.Email = addr.Address

	return nil
}

func (it *User) ValidateNewLocalUser() error {
	if err := it.validateRequiredFields(); err != nil {
		return err
	}

	if it.Password == "" {
		return domain.PasswordIsRequired
	}

	return nil
}

func (it *User) ValidateNewExternalUser() error {
	if err := it.validateRequiredFields(); err != nil {
		return err
	}

	if it.Provider == "" {
		return domain.ProviderIsRequired
	}

	return nil
}

func (it *User) ChangePassword(password string) {
	it.Password = password
}
