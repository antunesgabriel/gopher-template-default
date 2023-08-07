package domain

import "errors"

var (
	EmailIsRequired    error = errors.New("email_is_required")
	InvalidEmail       error = errors.New("invalid_email")
	NameIsRequired     error = errors.New("name_is_required")
	PasswordIsRequired error = errors.New("password_is_required")
	ProviderIsRequired error = errors.New("provider_is_required")
	UserIsNotRegister  error = errors.New("user_is_not_register")
	UserIsNotLocal     error = errors.New("user_is_not_local")
	InvalidPassword    error = errors.New("invalid_password")
	InvalidFields      error = errors.New("invalid_fields")
)
