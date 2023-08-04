package controller

import (
	"encoding/json"
	"net/http"

	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
)

type CreateLocalUserController struct {
	usecase *usecase.CreateLocalUserUseCase
}

func NewCreateLocalUserController(uc *usecase.CreateLocalUserUseCase) *CreateLocalUserController {
	c := CreateLocalUserController{
		usecase: uc,
	}

	return &c
}

func (it *CreateLocalUserController) Handle(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateUserLocalInput{}

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resp, _ := dto.NewResponse(nil, err).ToByte()

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)

		return
	}

	if err := it.usecase.Execute(input.Name, input.Email, input.Password); err != nil {
		resp, _ := dto.NewResponse(nil, err).ToByte()

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)

		return
	}

	resp, _ := dto.NewResponse(nil, nil).ToByte()

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
