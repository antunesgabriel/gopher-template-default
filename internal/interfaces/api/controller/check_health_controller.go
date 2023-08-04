package controller

import (
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
	w.Header().Add("Content-Type", "application/json")

	if err := it.usecase.Execute(); err != nil {
		w.WriteHeader(http.StatusFailedDependency)

		output := dto.HealthCheckOutput{
			Up: false,
		}

		response, _ := dto.NewResponse(output, err).ToByte()

		w.Write(response)

		return
	}

	output := dto.HealthCheckOutput{
		Up: true,
	}

	response, _ := dto.NewResponse(output, nil).ToByte()

	log.Println("->", string(response))

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
