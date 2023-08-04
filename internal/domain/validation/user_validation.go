package validation

type UserValidation string

var (
	EmailIsRequired              UserValidation = "email_is_required"
	InvalidEmail                 UserValidation = "invalid_email"
	NameIsRequired               UserValidation = "name_is_required"
	PasswordOrProviderIsRequired UserValidation = "password_or_provider_is_required"
)
