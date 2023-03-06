package handlers

import (
	"assignment-1/contextual_error_messages"
	"assignment-1/handlers/utility"
	"assignment-1/predefined"
	"assignment-1/requests/request_country"
	"assignment-1/requests/request_university"
	"encoding/json"
	"log"
	"net/http"
	"path"
)

/*
HandlerUniversityInformation handles incoming HTTP requests for University information
Param w: the http.ResponseWriter that the server uses to write the HTTP response
Param r: the http.Request pointer that contains the incoming request data.
Returns: nothing
*/
func HandlerUniversityInformation(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Set("Content-Type", "application/json")

	// Ensure that only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, contextual_error_messages.GetInvalidMethodError().Error(), http.StatusMethodNotAllowed)
		return
	}

	// Get the search string from the URL path
	search := path.Base(r.URL.Path)
	if search == "" {
		http.Error(w, "no search string provided", http.StatusBadRequest)
		return
	}

	// Retrieve universities that match the search string
	universities, err := request_university.RequestUniversityInformationByName(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return an HTTP status code indicating that no content was found if no universities are found
	if len(universities) == 0 {
		http.Error(w, contextual_error_messages.GetUniversitiesNotFoundError().Error(), http.StatusNoContent)
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

	// Retrieve the combined university information.
	combinedUniversityInformation := GetCombinedUniversityInformation(universities, fields, limit)

	// Encode the combined university information
	if err := json.NewEncoder(w).Encode(combinedUniversityInformation); err != nil {
		http.Error(w, contextual_error_messages.GetEncodingError().Error(), http.StatusInternalServerError)
		return
	}
}

/*
GetCombinedUniversityInformation combines a list of information about universities with their corresponding country.
Param universities: A list of universities to combine with their corresponding countries.
Param fields: A list of fields to be included in the combined structs.
Param limit: An integer representing the maximum number of results to be returned.
Returns: A list of combined universities and countries
*/
func GetCombinedUniversityInformation(universities []predefined.University, fields []string, limit int) []predefined.UniversityAndCountry {
	universityInformationList := make([]predefined.UniversityAndCountry, 0, len(universities))
	for _, university := range universities {
		country, err := request_country.GetCountryInformationByAlphaCode(university.AlphaTwoCode)
		if err != nil {
			log.Printf("failed to get country information for university %v: %v", university.Name, err)
			continue
		}
		universityInformation := predefined.CombineUniversityAndCountry(university, country, fields...)
		if limit == 0 || len(universityInformationList) < limit {
			universityInformationList = append(universityInformationList, universityInformation)
		} else {
			break
		}
	}
	return universityInformationList
}
