package dto

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestNewResponse(t *testing.T) {
	t.Run("it should create new response with correct payload", func(t *testing.T) {
		data := 42

		response := NewResponse(data, nil)

		if response.Data != data {
			t.Errorf("got %d want %d", response.Data, data)
		}

		if response.Error != "" {
			t.Errorf("got %s want %v", response.Error, "")
		}
	})

	t.Run("it should parse response to json with correct properties", func(t *testing.T) {
		data := 42
		errMessage := "example"

		expected := struct {
			Error string `json:"error"`
			Data  int    `json:"data"`
		}{}

		response := NewResponse(data, errors.New(errMessage))

		payload, err := json.Marshal(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error")
		}

		if err := json.Unmarshal(payload, &expected); err != nil {
			t.Errorf("got %s want %s", err, "no error")
		}

		if expected.Data != data {
			t.Errorf("got %d want %d", expected.Data, data)
		}

		if expected.Error != errMessage {
			t.Errorf("got %s want %s", expected.Error, errMessage)
		}
	})
}
