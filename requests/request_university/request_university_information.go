package request_university

// File containing helper_functions to request university information from the University-API
import (
	"assignment-1/predefined"
	"assignment-1/requests"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
func RequestUniversityInformation(url string) ([]predefined.University, error) {
	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	switch {
	case res.StatusCode == http.StatusNotFound:
		return []predefined.University{}, errors.New(fmt.Sprintf("%d There were no university found", res.StatusCode))
	case res.StatusCode != http.StatusOK:
		return []predefined.University{}, errors.New(fmt.Sprintf("The status code returned from the universityAPI: %d", res.StatusCode))
	}
	var universities []predefined.University
	if universities, err = DecodeUniversityInformation(res); err != nil {
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
	var universities []predefined.University
	if err := decoder.Decode(&universities); err != nil {
		log.Println(err)
		return nil, err
	}
	return universities, nil
}
