package requests

// File containing a helper_function to create and do requests
import (
	"net/http"
	"strings"
)

/*
CreateAndDoRequest creates an HTTP request with the given method and URL.
Param method: HTTP method to use in the request.
Param url: URL to be used in the request.
Returns: An HTTP response and error, or nil and an error.
*/
func CreateAndDoRequest(method string, url string) (*http.Response, error) {
	// Encode any spaces in the URL to "%20"
	url = strings.ReplaceAll(url, " ", "%20")
	// Create the request object
	req, err := http.NewRequest(method, url, nil)
	// Set the request header to indicate that the content type is JSON
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	// Send the request using the HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
