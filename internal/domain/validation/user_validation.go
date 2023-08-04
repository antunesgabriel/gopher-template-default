package validation

type UserValidation string

var (
	EmailIsRequired    UserValidation = "email_is_required"
	InvalidEmail       UserValidation = "invalid_email"
	NameIsRequired     UserValidation = "name_is_required"
	PasswordIsRequired UserValidation = "password_is_required"
	ProviderIsRequired UserValidation = "provider_is_required"
)
