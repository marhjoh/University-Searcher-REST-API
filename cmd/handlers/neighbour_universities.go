package handlers

import (
	"assignment-1/cmd/handlers/constants"
	"assignment-1/cmd/handlers/requests"
	"assignment-1/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
HandlerNeighbourUniversities handles requests for neighbour universities.
Param w: The http.ResponseWriter
Param r: The http.Request pointer.
*/
func HandlerNeighbourUniversities(w http.ResponseWriter, r *http.Request) {
	// Only GET method is allowed.
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
		return
	}

	// Extracts the country and university query from the URL path.
	pathList := strings.Split(r.URL.Path, "/")
	countryQuery := pathList[len(pathList)-2]
	uniQuery := pathList[len(pathList)-1]

	var fields []string
	if r.URL.Query().Get("fields") != "" {
		// Splits the fields query parameter into a slice of strings.
		fields = strings.Split(r.URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}

	var limit int
	if r.URL.Query()["limit"] != nil {
		// Converts the limit query parameter to an integer and sets a default value if it is invalid.
		if l, err := strconv.Atoi(r.URL.Query()["limit"][0]); err != nil || l < 0 {
			http.Error(w, "The limit is not a valid positive integer. Using 0 as limit.", http.StatusBadRequest)
			limit = constants.LIMIT_DEFAULT
		} else {
			limit = l
		}
	} else {
		limit = constants.LIMIT_DEFAULT
	}

	// Retrieves the borders of the country from the countries API.
	country, err := requests.RequestCountryInformation(fmt.Sprintf("%sname/%s?fields=borders", constants.COUNTRIESAPI_URL, countryQuery))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var countries []structs.Country
	for _, bor := range country.Borders {
		// Retrieves country information for each border from the countries API.
		url := fmt.Sprintf("%salpha/%s?fields=name,languages,maps", constants.COUNTRIESAPI_URL, bor)
		c, _ := requests.RequestCountryInformation(url)
		countries = append(countries, c)
	}

	var universityInformation []structs.UniversityAndCountry
	for _, c := range countries {
		// Searches for universities that match the university query in the current country.
		query := fmt.Sprintf("search?name=%s&country=%s", uniQuery, c.Name["common"].(string))
		url := constants.UNIVERSITIESAPI_URL + query
		universities, _ := requests.RequestUniversityInformation(url)
		for _, u := range universities {
			// Limits the number of universities to return and combines them with country information.
			if limit == 0 || len(universityInformation) < limit {
				universityInformation = append(universityInformation, structs.CombineUniversityAndCountry(u, c, fields...))
			} else {
				break
			}
		}
		if len(universityInformation) >= limit {
			break
		}
	}
	if len(universityInformation) == 0 {
		http.Error(w, "No universities were found", http.StatusNoContent)
		return
	}

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err = encoder.Encode(universityInformation)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}
}
