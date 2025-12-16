package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, config.baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithSecretKey(t *testing.T) {
	t.Run("SecretKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		key := "secret-key123"

		// Act
		WithSecretKey(key).apply(config)

		// Assert
		assert.Equal(t, key, config.secretKey)
	})
}

func TestWithPublicKey(t *testing.T) {
	t.Run("PublicKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		key := "public-key123"

		// Act
		WithPublicKey(key).apply(config)

		// Assert
		assert.Equal(t, key, config.publicKey)
	})
}
