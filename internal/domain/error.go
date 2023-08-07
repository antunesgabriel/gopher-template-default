package domain

import "errors"

var (
	EmailIsRequiredError    error = errors.New("email_is_required")
	InvalidEmailError       error = errors.New("invalid_email")
	NameIsRequiredError     error = errors.New("name_is_required")
	PasswordIsRequiredError error = errors.New("password_is_required")
	ProviderIsRequiredError error = errors.New("provider_is_required")
	UserIsNotRegisterError  error = errors.New("user_is_not_register")
	UserIsNotLocalError     error = errors.New("user_is_not_local")
	InvalidPasswordError    error = errors.New("invalid_password")
	InvalidFieldsError      error = errors.New("invalid_fields")
	UserAlreadyExistError   error = errors.New("user_already_exist")
)
