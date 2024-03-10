package testutils

import (
	"encoding/json"
	"io"
)

// ReadJSON reads a JSON object from the given body and decodes it
func ReadJSON[T interface{}](body io.ReadCloser, target T) T {
	if err := json.NewDecoder(body).Decode(&target); err != nil {
		panic(err)
	}
	return target
}
