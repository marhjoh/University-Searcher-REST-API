package handlers

import (
	"assignment-1/cmd/handlers/constants"
	"assignment-1/cmd/handlers/requests"
	"assignment-1/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

/*
HandlerUniversityInformation is an HTTP handler function for getting information about universities.
Param w: The http.ResponseWriter
Param r: The http.Request pointer.
*/
func HandlerUniversityInformation(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
	}

	// Get a list of universities matching the search term in the request URL
	universities := getUniversityInformation(r)

	// Return an error if no universities were found
	if len(universities) == 0 {
		http.Error(w, "No universities were found", http.StatusNoContent)
		return
	}

	// Get the list of fields to include in the response
	var fields []string
	if r.URL.Query().Get("fields") != "" {
		fields = strings.Split(r.URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}

	// Get the combined information about universities and countries
	combinedInfo := getCombined(universities, fields)

	// Set the content type of the response to JSON and encode the combined information as a JSON object
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(combinedInfo)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}
}

/*
getUniversityInformation is a helper function that retrieves information about universities matching the search term in the request URL.
Param r: The http.Request pointer.
*/
func getUniversityInformation(r *http.Request) []structs.University {
	search := path.Base(r.URL.Path)
	if search == "" {
		return nil
	}
	query := "search?name=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	universities, _ := requests.RequestUniversityInformation(url)
	if universities == nil {
		return nil
	}

	return universities
}

/*
getCombined is a helper function that combines information about universities and their countries.
Param universities: The universities to combine.
Param fields: fields to be included in the combined structs
Returns: a list of combined universities and countries
*/
func getCombined(universities []structs.University, fields []string) []structs.UniversityAndCountry {
	var universityInformations []structs.UniversityAndCountry
	for _, u := range universities {
		url := fmt.Sprintf("%sname/%s?fields=languages,maps", constants.COUNTRIESAPI_URL, u.Country)
		c, _ := requests.RequestCountryInformation(url)
		universityInformation := structs.CombineUniversityAndCountry(u, c, fields...)
		universityInformations = append(universityInformations, universityInformation)
	}
	return universityInformations
}
