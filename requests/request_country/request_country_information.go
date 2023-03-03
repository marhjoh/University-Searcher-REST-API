package request_country

// File containing helper_functions to request country information from the Country-API
import (
	"assignment-1/cache"
	"assignment-1/predefined"
	"assignment-1/requests"
	"bytes"
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
Returns: The country to retrieve, if found, and an error, if rasied.
*/
func GetCountryInformationByName(countryName string) (predefined.Country, error) {
	//Tries to get country from cache
	if c, err := cache.GetCountryFromCache(countryName); err == nil {
		return c, nil
	} else {
		// Country not in cache, retrieves from the API instead
		url := fmt.Sprintf("%sname/%s?fields=%s", predefined.COUNTRIESAPI_URL, countryName, predefined.COUNTRIESAPI_STANDARD_FIELDS)
		var countries []predefined.Country
		if countries, err = RequestCountryInformation(url); err != nil {
			return predefined.Country{}, err
		}

		// Goes through the list of retrieved country to find the correct one.
		for _, c := range countries {
			if strings.ToLower(c.Name["common"].(string)) == strings.ToLower(countryName) ||
				strings.ToLower(c.Name["official"].(string)) == strings.ToLower(countryName) {
				// Adds the newly retrieved country to the cache
				if err = cache.AddCountryToCache(c); err != nil {
					log.Println(err)
				}
				return c, nil
			}
		}

		return predefined.Country{}, errors.New(fmt.Sprintf("%s not found in result", countryName))
	}

}

/*
GetCountryInformationByAlphaCode retrieves country information based on the provided AlphaCode, which can be either CCA2 or CCA3.
Tries to get the country from cache, and retrieves it from the API if it's not present.
Param AlphaCode: AlphaCode of the country to retrieve, either as CCA2 or CCA3.
Returns: The country to retrieve, if found, and an error, if rasied.
*/

/*
GetCountryInformationByAlphaCode retrieves country information based on the provided AlphaCode, which can be either CCA2 or CCA3.
It first tries to retrieve the country information from the cache. If the information is not present in the cache, it is retrieved from the API.
Param alpha: AlphaCode of the country to retrieve, either as CCA2 or CCA3.
Returns: The country information for the specified AlphaCode, if found, and an error otherwise.
*/
func GetCountryInformationByAlphaCode(alphaCode string) (predefined.Country, error) {
	// Tries to get country from cache
	if country, err := cache.GetCountryByAlphaCodeFromCache(alphaCode); err == nil {
		return country, nil
	} else {
		// Country not found in cache, retrieves from the API instead
		url := fmt.Sprintf("%salpha/%s?fields=%s", predefined.COUNTRIESAPI_URL, alphaCode, predefined.COUNTRIESAPI_STANDARD_FIELDS)
		var countries []predefined.Country
		if countries, err = RequestCountryInformation(url); err != nil {
			return predefined.Country{}, err
		}

		for _, country := range countries {
			if country.CCA3 == alphaCode || country.CCA2 == alphaCode {
				// Adds the newly retrieved country to the cache
				if err = cache.AddCountryToCache(country); err != nil {
					log.Println(err)
				}
				return country, nil
			}
		}

		return predefined.Country{}, errors.New(fmt.Sprintf("%s not found in result", alphaCode))
	}
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
	var c []predefined.Country
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
func DecodeCountryInformation(res *http.Response) ([]predefined.Country, error) {

	// Read the response body into a byte array
	data, _ := io.ReadAll(res.Body)
	// Remove any white space from the byte array
	trimmedData := bytes.TrimLeft(data, " \t\r\n")
	// Determine if the trimmed byte array represents a JSON array or object
	isArray := len(trimmedData) > 0 && trimmedData[0] == '['

	// Create an empty slice to hold the decoded country data
	var countries []predefined.Country
	// If the response body represents a JSON array, decode the array into a slice of countries
	if isArray {
		if err := json.Unmarshal(data, &countries); err != nil {
			return nil, err
		}
		return countries, nil
	} else {
		// If the response body represents a JSON object, decode the object into a country struct
		var country predefined.Country
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
