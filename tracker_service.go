package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// trackerService is the API client for the `/contacts` endpoint
type trackerService service

// TrackEvent tracks an event for a contact.
//
// API Docs: https://next-wiki.useplunk.com/api-reference/public-api/trackEvent
func (service *trackerService) TrackEvent(ctx context.Context, params *TrackEventRequest) (*TrackEventResponse, *Response, error) {
	request, err := service.client.newRequestWithPublicKey(ctx, http.MethodPost, "/v1/track", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	event := new(TrackEventResponse)
	if err = json.Unmarshal(*response.Body, event); err != nil {
		return nil, response, err
	}

	return event, response, nil
}
