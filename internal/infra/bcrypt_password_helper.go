package infra

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHelper struct{}

func NewBcryptPasswordHelper() *BcryptPasswordHelper {
	h := BcryptPasswordHelper{}

	return &h
}

func (it *BcryptPasswordHelper) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (it *BcryptPasswordHelper) Compare(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
