package htt

import (
	"fmt"
	"net/http"
	"net/url"
)

func (r *Request) Get(URL string) (*Response, error) {
	return r.Do(http.MethodGet, URL)
}

func (r *Request) GetURL(URL *url.URL) (*Response, error) {
	return r.Do(http.MethodGet, URL.String())
}

func (r *Request) Getf(format string, a ...any) (*Response, error) {
	return r.Get(fmt.Sprintf(format, a...))
}

func (r *Request) Post(URL string) (*Response, error) {
	return r.Do(http.MethodPost, URL)
}

func (r *Request) PostURL(URL *url.URL) (*Response, error) {
	return r.Do(http.MethodPost, URL.String())
}

func (r *Request) Postf(format string, a ...any) (*Response, error) {
	return r.Post(fmt.Sprintf(format, a...))
}

func (r *Request) Put(URL string) (*Response, error) {
	return r.Do(http.MethodPut, URL)
}

func (r *Request) PutURL(URL *url.URL) (*Response, error) {
	return r.Do(http.MethodPut, URL.String())
}

func (r *Request) Putf(format string, a ...any) (*Response, error) {
	return r.Put(fmt.Sprintf(format, a...))
}

func (r *Request) Patch(URL string) (*Response, error) {
	return r.Do(http.MethodPatch, URL)
}

func (r *Request) PatchURL(URL *url.URL) (*Response, error) {
	return r.Do(http.MethodPatch, URL.String())
}

func (r *Request) Patchf(format string, a ...any) (*Response, error) {
	return r.Patch(fmt.Sprintf(format, a...))
}

func (r *Request) Delete(URL string) (*Response, error) {
	return r.Do(http.MethodDelete, URL)
}

func (r *Request) DeleteURL(URL *url.URL) (*Response, error) {
	return r.Do(http.MethodDelete, URL.String())
}

func (r *Request) Deletef(format string, a ...any) (*Response, error) {
	return r.Delete(fmt.Sprintf(format, a...))
}
