package htt

import (
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

// Response wraps the standard *http.Response to provide additional helper methods
// for reading and parsing the response body in different formats.
type Response struct {
	*http.Response
}

// JSON reads the response body and unmarshal it into the provided object (v) as JSON.
func (r *Response) JSON(v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

// Text reads the response body and returns it as a string.
func (r *Response) Text() (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Bytes reads the response body and returns it as a byte slice.
func (r *Response) Bytes() ([]byte, error) {
	return io.ReadAll(r.Body)
}
