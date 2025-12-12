package stubs

// ContactsCreateResponse returns a stubbed response for creating a contact
func ContactsCreateResponse() []byte {
	return []byte(`
{
  "success": true,
  "data": {
    "id": "string",
    "email": "user@example.com",
    "subscribed": true,
    "data": {},
    "createdAt": "2019-08-24T14:15:22Z",
    "updatedAt": "2019-08-24T14:15:22Z"
  }
}
`)
}

// ContactsDeleteResponse returns a stubbed response for deleting a contact
func ContactsDeleteResponse() []byte {
	return []byte(`
{
  "success": true,
  "data": {
    "message": "string"
  }
}
`)
}

// ContactsListResponse returns a stubbed response for listing a contact
func ContactsListResponse() []byte {
	return []byte(`
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "string",
        "email": "user@example.com",
        "subscribed": true,
        "data": {},
        "createdAt": "2019-08-24T14:15:22Z",
        "updatedAt": "2019-08-24T14:15:22Z"
      }
    ],
    "nextCursor": "string",
    "hasMore": true,
    "total": 0
  }
}
`)
}
