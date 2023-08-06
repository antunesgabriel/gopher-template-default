package controller

import (
	"encoding/json"
	"log"
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

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response := dto.NewResponse(nil, err)

		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode(&response)

		if err != nil {
			log.Println(err)

			w.Write([]byte(""))
		}

		return
	}

	if err := it.usecase.Execute(input.Name, input.Email, input.Password); err != nil {
		response := dto.NewResponse(nil, err)

		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode(&response)

		if err != nil {
			log.Println(err)

			w.Write([]byte(""))
		}

		return
	}

	response := dto.NewResponse(nil, nil)

	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(&response)

	if err != nil {
		log.Println(err)

		w.Write([]byte(""))
	}
}
