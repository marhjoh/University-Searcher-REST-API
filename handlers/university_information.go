package handlers

import (
	"assignment-1/handlers/utility"
	"assignment-1/predefined"
	"assignment-1/requests/request_country"
	"assignment-1/requests/request_university"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

/*
HandlerUniversityInformation handles incoming HTTP requests for University information
Param w: the http.ResponseWriter that the server uses to write the HTTP response
Param r: the http.Request pointer that contains the incoming request data.
Returns:
*/
func HandlerUniversityInformation(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Add("content-type", "application/json")

	// Check if the HTTP method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	// Retrieve the university information based on the query parameters
	universities := GetUniversityInformation(r)

	// Return an HTTP status code indicating that no content was found if no universities are found
	if len(universities) == 0 {
		http.Error(w, "No universities were found", http.StatusNoContent)
		return
	}

	// Retrieve the specified fields from the query
	fields := utility.GetFields(r.URL)

	// Retrieve the limit from the query
	limit, err := utility.GetLimit(r.URL)
	if err != nil {
		// Add not valid limit error to response
		http.Error(w, "The limit that was set is not a valid positive integer. Limit set as 0", http.StatusBadRequest)
		return
	}

	// Retrieve the combined university information.
	combinedUniversityInformation := GetCombinedUniversityInformation(universities, fields, limit)

	// Encode the combined university information
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(combinedUniversityInformation)
	if err != nil {
		http.Error(w, "There were an error during encoding", http.StatusInternalServerError)
		return
	}

}

/*
GetUniversityInformation retrieves information about universities by their name.
Param r: the http.Request pointer that contains the incoming request data.
Returns: An array of found university objects that match the search string, or nil.
*/
func GetUniversityInformation(r *http.Request) []predefined.University {
	// Retrieve the search-string
	search := path.Base(r.URL.Path)
	if search == "" {
		return nil
	}
	// Retrieve the resulting universities
	universities, _ := request_university.RequestUniversityInformationByName(search)
	if universities == nil {
		return nil
	}
	return universities
}

/*
GetCombinedUniversityInformation combines a list of information about universities with their corresponding country.
Param universities: A list of universities to combine with their corresponding countries.
Param fields: A list of fields to be included in the combined structs.
Param limit: An integer representing the maximum number of results to be returned.
Returns: A list of combined universities and countries
*/
func GetCombinedUniversityInformation(universities []predefined.University, fields []string, limit int) []predefined.UniversityAndCountry {
	var universityInformationList []predefined.UniversityAndCountry
	// Request information about the country of the university
	for _, university := range universities {
		country, err := request_country.GetCountryInformationByAlphaCode(university.AlphaTwoCode)
		if err != nil {
			// Print the error message if it occurs
			fmt.Println(err)
		}
		// Combine the university and country information
		universityInformation := predefined.CombineUniversityAndCountry(university, country, fields...)
		if limit == 0 || len(universityInformationList) < limit {
			// Add the combined university and country information to the list if the limit has not been reached
			universityInformationList = append(universityInformationList, universityInformation)
		} else {
			// Exit the loop if the limit has been reached
			break
		}
	}
	return universityInformationList
}
