package user

import (
	"errors"
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

func (u *User) Validate() error {
	if u.ID == 0 {
		return errors.New("")
	}

	return nil
}
