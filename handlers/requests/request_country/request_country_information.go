package request_country

import (
	"assignment-1/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
RequestCountryInformation retrieves country information from the University-API for a given URL.
Parameter url - The URL to be requested.
Returns universities containing the country information, or an error.
*/
func RequestCountryInformation(url string) (structs.Country, error) {
	var alphaSearch bool
	urlParts := strings.Split(url, "/")
	if urlParts[len(urlParts)-2] == "alpha" {
		alphaSearch = true
	} else {
		alphaSearch = false
	}

	_, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return structs.Country{}, err
	}

	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case res.StatusCode == http.StatusNotFound:
		return structs.Country{}, errors.New(fmt.Sprintf("%d The country was not found", res.StatusCode))
	case res.StatusCode != http.StatusOK:
		return structs.Country{}, errors.New(fmt.Sprintf("The status code returned from the countryAPI: %d", res.StatusCode))
	}

	decoder := json.NewDecoder(res.Body)
	if alphaSearch {
		var country structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}
		return country, nil
	} else {
		var country []structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}
		return country[0], nil
	}
}
