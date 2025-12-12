package client

import "time"

// Contact represents a contact object in the system
type Contact struct {
	ID         string         `json:"id"`
	Email      string         `json:"email"`
	Subscribed bool           `json:"subscribed"`
	Data       map[string]any `json:"data"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

// ContactCreateRequest represents the payload to create or update a contact
type ContactCreateRequest struct {
	Email      string         `json:"email"`
	Subscribed bool           `json:"subscribed"`
	Data       map[string]any `json:"data,omitempty"`
}

// ContactListRequest represents the payload to list contacts with optional filters
type ContactListRequest struct {
	Limit      *int    `json:"limit"`
	Cursor     *string `json:"cursor"`
	Subscribed *bool   `json:"subscribed"`
	Search     *string `json:"search"`
}

// ContactCreateResponse represents the response from creating or updating a contact
type ContactCreateResponse struct {
	Success bool    `json:"success"`
	Data    Contact `json:"data"`
}

// ContactDeleteResponse represents the response from deleting a contact
type ContactDeleteResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Message string `json:"message"`
	} `json:"data"`
}

// ContactListResponse represents the response from listing contacts
type ContactListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items      []Contact `json:"items"`
		NextCursor string    `json:"nextCursor"`
		HasMore    bool      `json:"hasMore"`
		Total      int       `json:"total"`
	} `json:"data"`
}
