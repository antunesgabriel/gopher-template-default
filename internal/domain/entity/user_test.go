package entity

import (
	"github.com/antunesgabriel/gopher-template-default/internal/domain/validation"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("it should create new user with correct params", func(t *testing.T) {
		id := int64(1)
		name := "Antunes"
		email := "antunes@test.io"
		provider := "github"
		password := "secret"

		u := New(id, name, email, provider, password)

		if u == nil {
			t.Errorf("got %s want %s", "nil", "user instance")
		}

		if u.Name != name {
			t.Errorf("got %s want %s", u.Name, name)
		}

		if u.Email != email {
			t.Errorf("got %s want %s", u.Email, email)
		}

		if u.Provider != provider {
			t.Errorf("got %s want %s", u.Provider, provider)
		}

		if u.Password != password {
			t.Errorf("got %s want %s", u.Password, password)
		}
	})

	t.Run("CheckIfNewUserIsValid()", func(t *testing.T) {

		t.Run("it should returns error if new user name is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := "google"
			password := "secrets"

			inValidUser := New(id, "", email, provider, password)
			validUser := New(id, name, email, provider, password)

			expectedErr := inValidUser.CheckIfNewUserIsValid()
			noExpectedErr := validUser.CheckIfNewUserIsValid()

			if expectedErr == nil || expectedErr.Error() != string(validation.NameIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.NameIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})

		t.Run("it should returns error if new user email is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := "google"
			password := "secrets"

			inValidUser := New(id, name, "", provider, password)
			validUser := New(id, name, email, provider, password)

			expectedErr := inValidUser.CheckIfNewUserIsValid()
			noExpectedErr := validUser.CheckIfNewUserIsValid()

			if expectedErr == nil || expectedErr.Error() != string(validation.EmailIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.EmailIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})

		t.Run("it should returns error if password and provider are empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := "google"
			password := "secrets"

			inValidUser := New(id, name, email, "", "")
			validUser := New(id, name, email, provider, password)

			expectedErr := inValidUser.CheckIfNewUserIsValid()
			noExpectedErr := validUser.CheckIfNewUserIsValid()

			if expectedErr == nil || expectedErr.Error() != string(validation.PasswordOrProviderIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.PasswordOrProviderIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})
	})
}