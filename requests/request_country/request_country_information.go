package request_country

// File containing helper_functions to request university information from the Country-API or the cache.
import (
	"assignment-1/cache"
	"assignment-1/contextual_error_messages"
	"assignment-1/predefined"
	"assignment-1/requests"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
GetCountryInformationByName Retrieves the country specified by the name. Tries to get the country from cache, and retrieves it
from the API if it's not present.
Param name: Name of the country to retrieve.
Returns: The country to retrieve, and an error (if rasied).
*/
func GetCountryInformationByName(countryName string) (predefined.Country, error) {
	// Get country from cache
	country, err := cache.GetCountryFromCache(countryName)
	if err == nil {
		return country, nil
	}

	// Retrieve from the API
	url := fmt.Sprintf("%sname/%s?fields=%s", predefined.COUNTRIESAPI_URL, countryName, predefined.COUNTRIESAPI_STANDARD_FIELDS)
	countries, err := RequestCountryInformation(url)
	if err != nil {
		return predefined.Country{}, err
	}

	// Find the match.
	for _, country := range countries {
		if strings.EqualFold(country.Name["common"].(string), countryName) ||
			strings.EqualFold(country.Name["official"].(string), countryName) {
			// Add to cache
			if err := cache.AddCountryToCache(country); err != nil {
				log.Println(err)
			}
			return country, nil
		}
	}

	return predefined.Country{}, errors.New(fmt.Sprintf("%s not found in result", countryName))
}

/*
GetCountryInformationByAlphaCode retrieves country information based on the provided AlphaCode, which can be either CCA2 or CCA3.
It first tries to retrieve the country information from the cache. If the information is not present in the cache, it is retrieved from the API.
Param alpha: AlphaCode of the country to retrieve, either as CCA2 or CCA3.
Returns: The country information for the specified AlphaCode, or an error.
*/
func GetCountryInformationByAlphaCode(alphaCode string) (predefined.Country, error) {
	// Get country from cache
	country, err := cache.GetCountryByAlphaCodeFromCache(alphaCode)
	if err == nil {
		return country, nil
	}

	// Retrieve from the API
	url := fmt.Sprintf("%salpha/%s?fields=%s", predefined.COUNTRIESAPI_URL, alphaCode, predefined.COUNTRIESAPI_STANDARD_FIELDS)
	countries, err := RequestCountryInformation(url)
	if err != nil {
		return predefined.Country{}, err
	}

	for _, country := range countries {
		if country.CCA3 == alphaCode || country.CCA2 == alphaCode {
			// Add to cache
			if err := cache.AddCountryToCache(country); err != nil {
				log.Println(err)
			}
			return country, nil
		}
	}

	return predefined.Country{}, errors.New(fmt.Sprintf("%s not found in result", alphaCode))
}

/*
RequestCountryInformation requests an HTTP GET request to the given URL and decodes the country information from the response.
Param URL: the URL to use in the request
Returns: A list of countries containing the decoded country information, or an error.
*/
func RequestCountryInformation(url string) ([]predefined.Country, error) {
	// Send HTTP GET request to specified URL
	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	var country []predefined.Country
	if country, err = DecodeCountryInformation(res); err != nil {
		return nil, err
	}
	return country, nil
}

/*
DecodeCountryInformation decodes a JSON response into a list of countries.
Checks whether the JSON response is an array or an object, and decodes it accordingly.
Param res: the HTTP response to decode.
Returns: A list of decoded countries, or an error and an empty list.
*/
func DecodeCountryInformation(res *http.Response) ([]predefined.Country, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var countries []predefined.Country
	decoder := json.NewDecoder(res.Body)

	// Use json.RawMessage to store the JSON data
	var data json.RawMessage
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	// Use a switch statement to decode the JSON data based on its type
	switch data[0] {
	case '[':
		if err := json.Unmarshal(data, &countries); err != nil {
			return nil, err
		}
	case '{':
		var country predefined.Country
		if err := json.Unmarshal(data, &country); err != nil {
			return nil, err
		}
		if country.Name == nil {
			return countries, contextual_error_messages.GetCountriesNotFoundError()
		}
		countries = append(countries, country)
	default:
		return nil, fmt.Errorf("unexpected JSON data")
	}

	return countries, nil
}
