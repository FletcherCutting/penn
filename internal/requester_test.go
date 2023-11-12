package internal

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_ResponsePayload_String(t *testing.T) {
	testCases := []struct {
		name             string
		responseContents responseContents
	}{
		{
			name: "Basic",
			responseContents: responseContents{
				statusCode: 200,
				body:       "Hello, world!",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.responseContents.String()
			snaps.MatchSnapshot(t, str)
		})
	}
}
