package infra

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcryptPasswordHelper(t *testing.T) {
	t.Run("Hash() should encrypt the password it received", func(t *testing.T) {
		helper := NewBcryptPasswordHelper()

		password := "nana"

		encryptBytes, err := helper.Hash(password)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if err := bcrypt.CompareHashAndPassword(encryptBytes, []byte(password)); err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}
	})

	t.Run("Compare() should compare an encrypted password with a password and return error if they are different", func(t *testing.T) {
		helper := NewBcryptPasswordHelper()

		password := "vasco"

		encryptBytes, _ := helper.Hash(password)

		if err := helper.Compare("wrong", string(encryptBytes)); err == nil {
			t.Errorf("got %s want %s", "nil", "error")
		}

		if err := helper.Compare(password, string(encryptBytes)); err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}
	})
}
