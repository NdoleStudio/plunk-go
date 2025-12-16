package client

import "time"

// TrackEventRequest represents the payload for tracking events
type TrackEventRequest struct {
	Email      string         `json:"email"`
	Event      string         `json:"event"`
	Subscribed bool           `json:"subscribed"`
	Data       map[string]any `json:"data,omitempty"`
}

// TrackEventResponse represents the response from tracking an event
type TrackEventResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Contact   string    `json:"contact"`
		Event     string    `json:"event"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"data"`
}
