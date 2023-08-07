package mock

import (
	"context"
	"errors"
)

type mockPasswordHelper struct {
	Hashed  []byte
	IsEqual bool
}

func NewMockPasswordHelper(hashed string, isEqual bool) *mockPasswordHelper {
	m := mockPasswordHelper{
		Hashed:  []byte(hashed),
		IsEqual: isEqual,
	}

	return &m
}

func (it *mockPasswordHelper) Hash(_ string) ([]byte, error) {
	return it.Hashed, nil
}

func (it *mockPasswordHelper) Compare(password, hashedPassword string) error {
	if it.IsEqual || password == hashedPassword {
		return nil
	}

	return errors.New("wrong")
}

type mockJWTHelper struct {
	id    int
	token string
}

func NewMockJWTHelper(id int, token string) *mockJWTHelper {
	m := mockJWTHelper{
		id,
		token,
	}

	return &m
}

func (it *mockJWTHelper) Encode(_ map[string]interface{}) (string, error) {

	return it.token, nil
}

func (it *mockJWTHelper) GetUserID(_ context.Context) (int, error) {

	return it.id, nil
}
