package handlers

import (
	"assignment-1/contextual_error_messages"
	"assignment-1/handlers/utility"
	"assignment-1/predefined"
	"assignment-1/requests/request_country"
	"assignment-1/requests/request_university"
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

/*
HandlerNeighbourUniversities handles incoming HTTP GET requests for the NeighbourUniversity/ies information(s).
Param w: the http.ResponseWriter that the server uses to write the HTTP response
Param r: the http.Request pointer that contains the incoming request data.
Returns: the combined university and country information in JSON format, or an error.
*/
func HandlerNeighbourUniversities(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Add("content-type", "application/json")

	// Respond with error if method is anything else than GET.
	if r.Method != http.MethodGet {
		http.Error(w, contextual_error_messages.GetInvalidMethodError().Error(), http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the search-terms from the path
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")

	countryQuery := pathList[len(pathList)-2]
	uniQuery := pathList[len(pathList)-1]

	if len(pathList) != 6 {
		http.Error(w, contextual_error_messages.GetInvalidNeighbourUniversityRequestError().Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the specified fields from the query
	fields := utility.GetFields(r.URL)

	// Retrieve the limit from the query
	limit, err := utility.GetLimit(r.URL)
	if err != nil {
		http.Error(w, contextual_error_messages.GetInvalidLimitError().Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the country to find neighbouring countries to
	country, err := request_country.GetCountryInformationByName(countryQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	// Retrieve all neighbouring countries from the "borders" field in the country struct.
	var countries []predefined.Country
	for _, bor := range country.Borders {
		c, _ := request_country.GetCountryInformationByAlphaCode(bor)
		countries = append(countries, c)
	}

	// Retrieve all alphaTwoCodes from the bordering countries, and appends them to a list
	var alphaTwoCodes []string
	for _, c := range countries {
		alphaTwoCodes = append(alphaTwoCodes, c.CCA2)
	}

	// Finds all universities with the given name-search, and retrieve the ones with the matching alphaTwoCode,
	// before combining them with the correct country
	universities, err := request_university.RequestUniversityInformationByName(uniQuery)
	var universityInformation []predefined.UniversityAndCountry
	for _, u := range universities {
		for _, a := range alphaTwoCodes {
			if u.AlphaTwoCode == a {
				c, _ := request_country.GetCountryInformationByAlphaCode(a)
				// If the list of results is smaller than the limit, append the university.
				// If not, break out of the loops.
				if limit == 0 || len(universityInformation) < limit {
					universityInformation = append(universityInformation, predefined.CombineUniversityAndCountry(u, c, fields...))
					break
				} else {
					break
				}
			}
		}
		if len(universityInformation) >= limit && limit != 0 {
			break
		}
	}

	if len(universityInformation) == 0 {
		http.Error(w, contextual_error_messages.GetUniversitiesNotFoundError().Error(), http.StatusNoContent)
		return
	}
	encoder := json.NewEncoder(w)
	// Ensure that the response looks like a json file in the browser.
	encoder.SetIndent("", "\t")
	err = encoder.Encode(universityInformation)
	if err != nil {
		http.Error(w, contextual_error_messages.GetEncodingError().Error(), http.StatusInternalServerError)
		return
	}
}
