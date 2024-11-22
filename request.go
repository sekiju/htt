package htt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// BodyEncoding represents the encoding type for the body of an HTTP request.
type BodyEncoding int

const (
	// NoEncoding indicates that the body will be sent as-is, without additional encoding.
	NoEncoding BodyEncoding = iota
	// JsonEncoding indicates that the body will be serialized to JSON format.
	JsonEncoding
)

// Request represents an HTTP request with customizable headers, body, and client settings.
type Request struct {
	headers      map[string]string // Custom headers for the HTTP request.
	body         interface{}       // The body of the HTTP request.
	bodyEncoding BodyEncoding      // Encoding type for the request body.
	client       *http.Client      // HTTP client used to execute the request.
}

// New creates and initializes a new Request instance with default settings.
func New() *Request {
	return &Request{
		headers:      map[string]string{},
		bodyEncoding: JsonEncoding,
		client:       Client,
	}
}

// SetHeader sets a single header key-value pair for the request.
func (r *Request) SetHeader(key, value string) *Request {
	r.headers[key] = value
	return r
}

// SetHeaders sets multiple headers for the request by replacing the current header map.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.headers = headers
	return r
}

// Body sets the body of the request. The format of the body depends on the encoding type.
func (r *Request) Body(v interface{}) *Request {
	r.body = v
	return r
}

// BodyEncoding sets the encoding type for the request body.
func (r *Request) BodyEncoding(encoding BodyEncoding) *Request {
	r.bodyEncoding = encoding
	return r
}

// Client sets a custom HTTP client for executing the request.
func (r *Request) Client(client *http.Client) *Request {
	r.client = client
	return r
}

// Do execute the HTTP request with the specified method and URL.
// It applies the configured headers, body, and context.
func (r *Request) Do(method, URL string) (*Response, error) {
	var body io.Reader
	if r.body != nil {
		switch r.bodyEncoding {
		case JsonEncoding:
			serialized, err := json.Marshal(r.body)
			if err != nil {
				return nil, err
			}
			body = bytes.NewBuffer(serialized)
		case NoEncoding:
			switch v := r.body.(type) {
			case []byte:
				body = bytes.NewReader(v)
			case io.Reader:
				body = v
			case string:
				body = bytes.NewReader([]byte(v))
			}
		}
	}

	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	if r.bodyEncoding == JsonEncoding {
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}
