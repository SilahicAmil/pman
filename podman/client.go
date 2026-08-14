package podman

import (
	"net/http"
	"time"
)

// establish connection to podman TCP API

// I think if we have a config in appdata for the port that might be better.
// For now I think like a port 9091 is fine

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
}

func New(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *HTTPClient) Get()    {}
func (h *HTTPClient) Post()   {}
func (h *HTTPClient) Patch()  {}
func (h *HTTPClient) Delete() {}
