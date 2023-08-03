package health

import "net/http"

type HealthController struct {
	healthService *HealthService
}

func NewHealthController(healthService *HealthService) *HealthController {
	c := HealthController{
		healthService: healthService,
	}

	return &c
}

func (c *HealthController) CheckHealth(w http.ResponseWriter, _ *http.Request) {
	health, err := c.healthService.Check()

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Header().Add("Content-Type", "application/json")

		output, err := NewHealth("down").ToByte()

		if err != nil {
			w.WriteHeader(http.StatusFailedDependency)
			w.Write([]byte("{\"up\": \"down\"}"))

			return
		}

		w.Write(output)

		return
	}

	b, _ := health.ToByte()

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

}
