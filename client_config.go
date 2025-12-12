package client

import "net/http"

type clientConfig struct {
	httpClient *http.Client
	secretKey  string
	baseURL    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://next-api.useplunk.com",
	}
}
