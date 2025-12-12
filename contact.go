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
	Contact
	Meta struct {
		IsNew    bool `json:"isNew"`
		IsUpdate bool `json:"isUpdate"`
	} `json:"_meta"`
}

// ContactListResponse represents the response from listing contacts
type ContactListResponse struct {
	Contacts   []Contact `json:"contacts"`
	Total      int       `json:"total"`
	HasMore    bool      `json:"hasMore"`
	NextCursor *string   `json:"nextCursor"`
}
