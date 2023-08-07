package dto

type LocalAuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LocalAuthOutput struct {
	Token string `json:"token"`
}
