package user

import (
	"encoding/json"
	"log"
	"net/http"

	"gihub.com/antunesgabriel/gopher-template-default/internal/app/shared"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	c := UserController{
		userService: userService,
	}

	return &c
}

func (c *UserController) StoreLocal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto := createUserLocalDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println("[ERROR]: ", err.Error())

		resp, _ := shared.NewResponse(nil, err.Error()).ToByte()

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write(resp)

		return
	}

	if err := c.userService.CreateLocal(ctx, dto.Name, dto.Email, dto.Password); err != nil {
		log.Println("[ERROR]: ", err.Error())

		resp, _ := shared.NewResponse(nil, err.Error()).ToByte()

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write(resp)

		return
	}

	resp, _ := shared.NewResponse(nil, "").ToByte()

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
