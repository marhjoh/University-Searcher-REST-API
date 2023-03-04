package cache

//This file contains the cache and the helper_functions that operate on it.
import (
	"assignment-1/predefined"
	"errors"
	"fmt"
	"strings"
	"time"
)

var cachedCountry = make(map[string]predefined.Country)

/*
AddCountryToCache country to be added to the cache.
Param country: The country to be added.
Returns: an error if the country is already cached.
*/
func AddCountryToCache(country predefined.Country) error {
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
func GetCountryFromCache(cachedCountryName string) (predefined.Country, error) {
	if country, exists := cachedCountry[strings.ToTitle(cachedCountryName)]; exists {
		if time.Since(country.Cache).Hours() > predefined.LIMIT_HOURS {
			// Delete the cached country if it has surpassed the limit
			delete(cachedCountry, country.Name["common"].(string))
			return predefined.Country{}, errors.New(fmt.Sprintf("%s was cached, but it has gone over %v hours since it was cached.",
				cachedCountryName, predefined.LIMIT_HOURS))
		}
		return country, nil
	} else {
		return predefined.Country{}, errors.New(fmt.Sprintf("%s is not cached.", cachedCountryName))
	}
}

/*
GetCountryByAlphaCodeFromCache retrieves a cached country by its alpha code (CCA2 or CCA3).
Param alphaCode: the Alpha code (CCA2 or CCA3) of the cached country to be retrieved.
Returns: The cached country, or an empty struct and an error.
*/
func GetCountryByAlphaCodeFromCache(alphaCode string) (predefined.Country, error) {
	for _, country := range cachedCountry {
		if strings.ToUpper(country.CCA3) == strings.ToUpper(alphaCode) {
			if time.Since(country.Cache).Hours() > predefined.LIMIT_HOURS {
				// Delete the cached country if it has surpassed the limit
				delete(cachedCountry, country.Name["common"].(string))
				return predefined.Country{}, errors.New(fmt.Sprintf("%s was cached, but it has gone over %v hours since it was cached.",
					cachedCountry, predefined.LIMIT_HOURS))
			}
			return country, nil
		} else if strings.ToUpper(country.CCA2) == strings.ToUpper(alphaCode) {
			if time.Since(country.Cache).Hours() > predefined.LIMIT_HOURS {
				// Delete the cached country if it has surpassed the limit
				delete(cachedCountry, country.Name["common"].(string))
				return predefined.Country{}, errors.New(fmt.Sprintf("%s was cached, but it has gone over %v hours since it was cached.",
					cachedCountry, predefined.LIMIT_HOURS))
			}
			return country, nil
		}
	}
	return predefined.Country{}, errors.New(fmt.Sprintf("%s is not cached.", alphaCode))
}
