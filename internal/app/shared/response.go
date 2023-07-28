package shared

import "encoding/json"

type response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewResponse(data any, errorMessage string) *response {
	r := response{
		Data:  data,
		Error: errorMessage,
	}

	return &r
}

func (r *response) ToByte() ([]byte, error) {
	return json.Marshal(r)
}
