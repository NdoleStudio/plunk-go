package stubs

// ContactsCreateResponse returns a stubbed response for creating a contact
func ContactsCreateResponse() []byte {
	return []byte(`
{
	"id": "string",
	"email": "user@example.com",
	"subscribed": true,
	"projectId": "44654d15-35aa-4573-a767-616015403763",
	"data": {
		"firstName": "John",
		"lastName": "Doe",
		"plan": "premium"
	},
	"createdAt": "2019-08-24T14:15:22Z",
	"updatedAt": "2019-08-24T14:15:22Z",
	"_meta": {
		"isNew":false,
		"isUpdate":false
	}
}
`)
}

// ContactsListResponse returns a stubbed response for listing a contact
func ContactsListResponse() []byte {
	return []byte(`
{
  "data": [
    {
      "id": "7d195254-bc27-4937-a73d-370ec0fb39b2",
      "email": "user@example.com",
      "data": {
        "firstName": "John",
        "lastName": "Doe",
        "plan": "premium"
      },
      "subscribed": true,
      "projectId": "44654d15-35aa-4573-a767-616015403763",
      "createdAt": "2025-12-16T17:53:23.347Z",
      "updatedAt": "2026-01-05T20:19:54.936Z"
    }
  ],
  "total": 1,
  "nextCursor": null,
  "hasMore": false
}
`)
}
