package request_country

// File containing helper_functions to request country information from the Country-API
import (
	"assignment-1/handlers/requests"
	"assignment-1/structs"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/*
RequestCountryInformation requests an HTTP GET request to the given URL and decodes the country information from the response.
Param URL: the URL to use in the request
Returns: A list of countries containing the decoded country information, or an error.
*/
func RequestCountryInformation(url string) ([]structs.Country, error) {
	// Send HTTP GET request to specified URL
	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	var c []structs.Country
	if c, err = DecodeCountryInformation(res); err != nil {
		return nil, err
	}
	return c, nil
}

/*
DecodeCountryInformation decodes a JSON response into a list of countries.
Checks whether the JSON response is an array or an object, and decodes it accordingly.
If the response is an object, it appends the decoded country to a list before returning it.
If the response is an empty object, it returns an empty list and an error indicating that no countries were found.
Param res: the HTTP response to decode.
Returns: A list of decoded countries, or an error and an empty list.
*/
func DecodeCountryInformation(res *http.Response) ([]structs.Country, error) {

	// Read the response body into a byte array
	data, _ := io.ReadAll(res.Body)
	// Remove any white space from the byte array
	trimmedData := bytes.TrimLeft(data, " \t\r\n")
	// Determine if the trimmed byte array represents a JSON array or object
	isArray := len(trimmedData) > 0 && trimmedData[0] == '['

	// Create an empty slice to hold the decoded country data
	var countries []structs.Country
	// If the response body represents a JSON array, decode the array into a slice of countries
	if isArray {
		if err := json.Unmarshal(data, &countries); err != nil {
			return nil, err
		}
		return countries, nil
	} else {
		// If the response body represents a JSON object, decode the object into a country struct
		var country structs.Country
		if err := json.Unmarshal(data, &country); err != nil {
			return nil, err
		}
		// If the country does not have a name, return an error
		if country.Name == nil {
			return countries, errors.New("There were no countries found")
		}
		// Append the decoded country data to the slice of countries
		countries = append(countries, country)
		return countries, nil
	}
}
