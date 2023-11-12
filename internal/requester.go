package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type requesterFunc func(context.Context, requestContents) (responseContents, error)

// Implements requesterFunc
// This function contains all the required logic to make HTTP requests according
// to a provided requestContents, then process and return a responseContents.
func HTTPRequester(_ context.Context, requestContents requestContents) (responseContents, error) {
	// Send request.
	client := &http.Client{}
	request, err := http.NewRequest(requestContents.method, requestContents.url, requestContents.body)

	if err != nil {
		return responseContents{}, fmt.Errorf("failed to create new http request: %w", err)
	}

	response, err := client.Do(request)

	if err != nil {
		return responseContents{}, fmt.Errorf("failed to do http request: %w", err)
	}

	defer response.Body.Close()

	// Make response contents.
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return responseContents{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseContents{
		statusCode: response.StatusCode,
		body:       string(body),
	}, nil
}

// Holds all information necessary to make a request.
type requestContents struct {
	method string
	url    string
	body   io.Reader
}

// Implements fmt.Stringer.
// Wraps all relevant information retrieved from the http response.
type responseContents struct {
	statusCode int
	body       string
}

// Generates a simplified string that presents all the http response information
// in the following form:
//
// <statusCode>
// <body>
//
// Example:
//
// 200
// Hello, world!
func (r responseContents) String() string {
	return fmt.Sprintf("%d\n%s", r.statusCode, r.body)
}
