package cache

//This file contains the cache and the helper_functions that operate on it.
import (
	"assignment-1/constants"
	"assignment-1/structs"
	"errors"
	"fmt"
	"strings"
	"time"
)

var cachedCountry = make(map[string]structs.Country)

/*
AddCountryToCache country to be added to the cache.
Param country: The country to be added.
Returns: an error if the country is already cached.
*/
func AddCountryToCache(country structs.Country) error {
	var cachedCountryName string
	cachedCountryName = strings.ToTitle(country.Name["common"].(string))
	if _, exists := cachedCountry[cachedCountryName]; !exists {
		country.Cache = time.Now()
		cachedCountry[cachedCountryName] = country
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s is already cached.", cachedCountryName))
	}
}

/*
GetCountryFromCache retrieves a cached country.
Param cachedCountryName: The name of the country to retrieve from the cache.
Returns: The cached country if it exists and is not too old, or an empty struct and an error otherwise.
*/
func GetCountryFromCache(cachedCountryName string) (structs.Country, error) {
	if c, exists := cachedCountry[strings.ToTitle(cachedCountryName)]; exists {
		if time.Since(c.Cache).Hours() > constants.LIMIT_HOURS {
			// Delete the cached country if it has surpassed the limit
			delete(cachedCountry, c.Name["common"].(string))
			return structs.Country{}, errors.New(fmt.Sprintf("%s was cached, but it has gone over %v hours since it was cached.",
				cachedCountryName, constants.LIMIT_HOURS))
		}
		return c, nil
	} else {
		return structs.Country{}, errors.New(fmt.Sprintf("%s is not cached.", cachedCountryName))
	}
}
