package controller

import (
	"encoding/json"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"net/http"
)

type AuthLocalController struct {
	usecase *usecase.LocalAuthUseCase
}

func NewAuthLocalController(usecase *usecase.LocalAuthUseCase) *AuthLocalController {
	c := AuthLocalController{
		usecase: usecase,
	}

	return &c
}

func (it AuthLocalController) Handle(w http.ResponseWriter, r *http.Request) {
	input := dto.LocalAuthInput{}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := dto.NewResponse(nil, err)

		_ = json.NewEncoder(w).Encode(&response)

		return
	}

	token, err := it.usecase.Execute(&input)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := dto.NewResponse(nil, err)

		_ = json.NewEncoder(w).Encode(&response)

		return
	}

	output := dto.LocalAuthOutput{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := dto.NewResponse(output, nil)

	_ = json.NewEncoder(w).Encode(&response)
}
