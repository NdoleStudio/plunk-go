package stubs

// ContactsCreateResponse returns a stubbed response for creating a contact
func ContactsCreateResponse() []byte {
	return []byte(`
{
	"id": "string",
	"email": "user@example.com",
	"subscribed": true,
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
  	"contacts": [
		{
			"id": "string",
			"email": "user@example.com",
			"subscribed": true,
			"data": {
				"firstName": "John",
				"lastName": "Doe",
				"plan": "premium"
			},
			"createdAt": "2019-08-24T14:15:22Z",
			"updatedAt": "2019-08-24T14:15:22Z"
		}
    ],
    "nextCursor": null,
    "hasMore": false,
    "total": 1
}
`)
}
