package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// contactService is the API client for the `/contacts` endpoint
type contactService service

// Create a new contact or update existing (upsert by email)
//
// API Docs: https://next-wiki.useplunk.com/api-reference/contacts/createContact
func (service *contactService) Create(ctx context.Context, params *ContactCreateRequest) (*ContactCreateResponse, *Response, error) {
	request, err := service.client.newRequestWithSecretKey(ctx, http.MethodPost, "/contacts", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	contact := new(ContactCreateResponse)
	if err = json.Unmarshal(*response.Body, contact); err != nil {
		return nil, response, err
	}

	return contact, response, nil
}

// Delete a contact permanently by ID
//
// API Docs: https://next-wiki.useplunk.com/api-reference/contacts/deleteContact
func (service *contactService) Delete(ctx context.Context, contactID string) (*Response, error) {
	request, err := service.client.newRequestWithSecretKey(ctx, http.MethodDelete, "/contacts/"+contactID, nil)
	if err != nil {
		return nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return response, err
	}

	return response, nil
}

// List a paginated list of contacts with cursor-based pagination
//
// API Docs: https://next-wiki.useplunk.com/api-reference/contacts/listContacts
func (service *contactService) List(ctx context.Context, params map[string]string) (*ContactListResponse, *Response, error) {
	request, err := service.client.newRequestWithSecretKey(ctx, http.MethodGet, "/contacts", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(service.client.addURLParams(request, params))
	if err != nil {
		return nil, response, err
	}

	contacts := new(ContactListResponse)
	if err = json.Unmarshal(*response.Body, contacts); err != nil {
		return nil, response, err
	}

	return contacts, response, nil
}
