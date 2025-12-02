package growwapi

import "net/http"

// Client to access groww apis
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new Client
func NewClient(httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return Client{httpClient: httpClient}
}
