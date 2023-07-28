package user

import "errors"

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Provider string `json:"provider,omitempty"`
	Password string `json:"password,omitempty"`
}

func New(id, name, email, provider, password string) *user {
	u := user{
		ID:       id,
		Name:     name,
		Email:    email,
		Provider: provider,
		Password: password,
	}

	return &u
}

func (u *user) Validate() error {
	if u.ID == "" {
		return errors.New("")
	}

	return nil
}
