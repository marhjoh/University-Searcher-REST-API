package requests

// File containing a helper_function to create and do requests
import (
	"assignment-1/httpclient"
	"net/http"
	"strings"
)

// Reuse the HTTP client to prevent creating a new one for each request
var client = httpclient.Client

/*
CreateAndDoRequest creates an HTTP request with the given method and URL.
Param method: HTTP method to use in the request.
Param url: URL to be used in the request.
Returns: An HTTP response and error, or nil and an error.
*/
func CreateAndDoRequest(method string, url string) (*http.Response, error) {
	// Encode any spaces in the URL to "%20"
	url = strings.ReplaceAll(url, " ", "%20")

	// Create the request object with the specified method and URL
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Set the request header to indicate that the content type is JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the HTTP client
	return client.Do(req)
}
