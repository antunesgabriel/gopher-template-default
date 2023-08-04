package validation

type USER_VALIDATION string

var (
	EMAIL_IS_REQUIRED                 USER_VALIDATION = "email_is_required"
	INVALID_EMAIL                     USER_VALIDATION = "invalid_email"
	NAME_IS_REQUIRED                  USER_VALIDATION = "name_is_required"
	PASSWORD_IS_REQUIRED_IN_THAT_CASE USER_VALIDATION = "password_is_required_in_that_case"
)
