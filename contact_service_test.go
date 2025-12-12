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

func TestContactService_Create(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.ContactsCreateResponse())
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))

	// Act
	contact, response, err := client.Contacts.Create(context.Background(), &ContactCreateRequest{
		Email:      "user@example.com",
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
	jsonContent, _ := json.Marshal(contact)
	assert.JSONEq(t, string(stubs.ContactsCreateResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestContactService_CreateWithError(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))

	// Act
	_, response, err := client.Contacts.Create(context.Background(), &ContactCreateRequest{
		Email:      "user@example.com",
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

func TestContactService_Delete(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusNoContent, nil)
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))

	// Act
	response, err := client.Contacts.Delete(context.Background(), "contact-id")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestContactService_DeleteWithError(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))

	// Act
	response, err := client.Contacts.Delete(context.Background(), "contact-id")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestContactService_List(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	request := new(http.Request)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.ContactsListResponse(), request)
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))
	params := map[string]string{
		"limit":      "10",
		"cursor":     "abc123",
		"subscribed": "true",
		"search":     "user@email.com",
	}

	// Act
	contact, response, err := client.Contacts.List(context.Background(), params)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	jsonContent, _ := json.Marshal(contact)
	assert.JSONEq(t, string(stubs.ContactsListResponse()), string(jsonContent))

	for key, value := range params {
		assert.Equal(t, value, request.URL.Query().Get(key))
	}

	// Teardown
	server.Close()
}

func TestContactService_ListWithError(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL), WithSecretKey(apiKey))

	// Act
	_, response, err := client.Contacts.List(context.Background(), map[string]string{})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
