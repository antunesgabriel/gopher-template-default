package dto

type CreateUserLocalInput struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required, email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserExternalDTO struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
