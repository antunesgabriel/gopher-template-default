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

		u := NewUser(id, name, email, provider, password)

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

	t.Run("ValidateNewLocalUser()", func(t *testing.T) {

		t.Run("it should returns error if new user name is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := "google"
			password := "secrets"

			inValidUser := NewUser(id, "", email, provider, password)
			validUser := NewUser(id, name, email, provider, password)

			expectedErr := inValidUser.ValidateNewLocalUser()
			noExpectedErr := validUser.ValidateNewLocalUser()

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

			inValidUser := NewUser(id, name, "", provider, password)
			validUser := NewUser(id, name, email, provider, password)

			expectedErr := inValidUser.ValidateNewLocalUser()
			noExpectedErr := validUser.ValidateNewLocalUser()

			if expectedErr == nil || expectedErr.Error() != string(validation.EmailIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.EmailIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})

		t.Run("if is local user it should return error if password is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := ""
			password := "secrets"

			inValidUser := NewUser(id, name, email, provider, "")
			validUser := NewUser(id, name, email, provider, password)

			expectedErr := inValidUser.ValidateNewLocalUser()
			noExpectedErr := validUser.ValidateNewLocalUser()

			if expectedErr == nil || expectedErr.Error() != string(validation.PasswordIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.PasswordIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})

		t.Run("if is local user it should return error if password is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := ""
			password := "secrets"

			inValidUser := NewUser(id, name, email, provider, "")
			validUser := NewUser(id, name, email, provider, password)

			expectedErr := inValidUser.ValidateNewLocalUser()
			noExpectedErr := validUser.ValidateNewLocalUser()

			if expectedErr == nil || expectedErr.Error() != string(validation.PasswordIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.PasswordIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})

		t.Run("if is external user it should return error if provider is empty", func(t *testing.T) {
			id := int64(1)
			name := "Dias"
			email := "dias@test.io"
			provider := "google"
			password := ""

			inValidUser := NewUser(id, name, email, "", password)
			validUser := NewUser(id, name, email, provider, password)

			expectedErr := inValidUser.ValidateNewExternalUser()
			noExpectedErr := validUser.ValidateNewExternalUser()

			if expectedErr == nil || expectedErr.Error() != string(validation.ProviderIsRequired) {
				t.Errorf("got %s want %s", expectedErr, validation.ProviderIsRequired)
			}

			if noExpectedErr != nil {
				t.Errorf("got %s want %s", noExpectedErr, "nil")
			}
		})
	})

	t.Run("ChangePassword() should change user password", func(t *testing.T) {
		user := NewUser(42, "Octavianus Augustus", "octavianus@augustus.com", "", "first")

		if user.Password != "first" {
			t.Errorf("got %s want %s", user.Password, "first")
		}

		user.ChangePassword("roman")

		if user.Password != "roman" {
			t.Errorf("got %s want %s", user.Password, "roman")
		}
	})
}
