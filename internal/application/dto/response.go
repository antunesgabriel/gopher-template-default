package dto

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewResponse(data interface{}, err error) *Response {
	r := Response{
		Data: data,
	}

	if err != nil {
		r.Error = err.Error()
	}

	return &r
}
