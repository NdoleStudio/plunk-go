package client

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/NdoleStudio/plunk-go/internal/helpers"
	"github.com/NdoleStudio/plunk-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestTrackerService_TrackEvent(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	request := new(http.Request)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.TrackEventResponse(), request)
	client := New(WithBaseURL(server.URL), WithPublicKey(apiKey))

	// Act
	event, response, err := client.Tracker.TrackEvent(context.Background(), &TrackEventRequest{
		Email:      "user@example.com",
		Event:      "purchase",
		Subscribed: true,
		Data: map[string]any{
			"firstName": "John",
			"lastName":  "Doe",
			"plan":      "premium",
		},
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "Bearer "+apiKey, request.Header.Get("Authorization"))
	jsonContent, _ := json.Marshal(event)
	assert.JSONEq(t, string(stubs.TrackEventResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestTrackerService_TrackEventWithError(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL), WithPublicKey(apiKey))

	// Act
	_, response, err := client.Tracker.TrackEvent(context.Background(), &TrackEventRequest{
		Email:      "user@example.com",
		Event:      "purchase",
		Subscribed: true,
		Data: map[string]any{
			"firstName": "John",
			"lastName":  "Doe",
			"plan":      "premium",
		},
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
