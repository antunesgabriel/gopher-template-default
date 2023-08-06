package mock

import "errors"

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

func (it *mockPasswordHelper) Compare(_ string, __ string) error {
	if it.IsEqual {
		return nil
	}

	return errors.New("wrong")
}
