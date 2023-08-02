package health

import "encoding/json"

type health struct {
	Status string `json:"status"`
}

func NewHealth(status string) *health {
	s := health{
		Status: status,
	}

	return &s
}

func (h *health) ToByte() ([]byte, error) {
	data, err := json.Marshal(h)

	return data, err
}
