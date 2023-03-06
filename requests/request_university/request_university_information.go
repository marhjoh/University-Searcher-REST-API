package request_university

// File containing helper_functions to request university information from the University-API or the cache.
import (
	"assignment-1/contextual_error_messages"
	"assignment-1/predefined"
	"assignment-1/requests"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/*
RequestUniversityInformationByName searches for universities by the given name.
Param search: university name to be searched for.
Return: An array of found universities, or an empty array and an error.
*/
func RequestUniversityInformationByName(search string) ([]predefined.University, error) {
	// Construct a query by appending the search term to the API's URL
	query := "search?name_contains=" + search
	url := predefined.UNIVERSITIESAPI_URL + query

	// Sends a request to the API and returns the response
	return RequestUniversityInformation(url)
}

/*
RequestUniversityInformation requests universities from the University-API using a given URL
Param url: the URL to use in the request.
Returns: A list of found universities, or it returns an empty list and an error.
*/
func RequestUniversityInformation(url string) (universities []predefined.University, err error) {
	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New(contextual_error_messages.GetUniversitiesNotFoundError().Error())
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Returned status code from University-API: %d", res.StatusCode))
	}

	universities, err = DecodeUniversityInformation(res)
	if err != nil {
		return nil, err
	}

	return universities, nil
}

/*
DecodeUniversityInformation Decodes a given response into an array of university/ies.
Param res: The HTTP response to decode.
Returns: An array of the decoded university/ies, or both an empty array and an error.
*/
func DecodeUniversityInformation(res *http.Response) ([]predefined.University, error) {
	// Create a JSON decoder using the response body.
	decoder := json.NewDecoder(res.Body)

	// Check if the response body is empty.
	if res.ContentLength == 0 {
		return nil, errors.New("response body is empty")
	}

	// Check if the response body contains valid JSON.
	var universities []predefined.University
	if err := decoder.Decode(&universities); err != nil {
		return nil, err
	}

	return universities, nil
}
