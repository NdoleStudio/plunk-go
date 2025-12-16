package stubs

// TrackEventResponse returns a stubbed response for tracking an event
func TrackEventResponse() []byte {
	return []byte(`
{
	  "success": true,
	  "data": {
			"contact": "6c53eb74-f3cd-4953-84da-6b88703b980a",
			"event": "7a592ad5-59dc-40bd-aec0-ff201d121f4e",
			"timestamp": "2025-12-16T16:45:24.649Z"
	  }
}
`)
}
