package utils

import (
	"net/http"
)

type AuthTransport struct {
	http.Transport
	Username string
	Password string
}

func (t AuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Method == "POST" || r.Method == "PUT" {
		r.Header.Set("Content-Type", "application/json")
	}
	r.SetBasicAuth(t.Username, t.Password)
	return t.Transport.RoundTrip(r)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHttpClient(user string, pass string) HttpClient {
	t := AuthTransport{
		Username: user,
		Password: pass,
	}

	return &http.Client{
		Transport: t,
	}

}
