package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
)

type CheckHealthController struct {
	usecase *usecase.CheckHealthUseCase
}

func NewCheckHealthController(usecase *usecase.CheckHealthUseCase) *CheckHealthController {
	c := CheckHealthController{
		usecase: usecase,
	}

	return &c
}

func (it *CheckHealthController) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := it.usecase.Execute(); err != nil {
		w.WriteHeader(http.StatusFailedDependency)

		output := dto.HealthCheckOutput{
			Up: false,
		}

		response := dto.NewResponse(output, err)

		err := json.NewEncoder(w).Encode(&response)

		if err != nil {
			log.Println(err)

			w.Write([]byte(""))
		}

		return
	}

	output := dto.HealthCheckOutput{
		Up: true,
	}

	response := dto.NewResponse(output, nil)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(&response)

	if err != nil {
		log.Println(err)

		w.Write([]byte(""))
	}
}
