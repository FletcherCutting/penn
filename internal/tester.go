package internal

import (
	"context"
	"fmt"
	"net/http"
)

// This is a temporary function for testing purposes.
func Request(ctx context.Context, requester requesterFunc) {
	// Read test files
	// Foreach test, send request
	// Compare output to success criteria

	url := "http://localhost:8123/hello"
	method := http.MethodGet
	requestContents := requestContents{
		method: method,
		url:    url,
		body:   nil,
	}

	responseContents, err := requester(ctx, requestContents)

	if err != nil {
		panic(fmt.Sprintf("failed to send request: %s", err))
	}

	fmt.Println(responseContents)
}
