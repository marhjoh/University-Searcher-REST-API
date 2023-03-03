package handlers

import (
	"assignment-1/predefined"
	"assignment-1/uptime"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
HandlerDiag is a handler for the /diag endpoint.
Param w: the http.ResponseWriter that the server uses to write the HTTP response
Param r: the http.Request pointer that contains the incoming request data.
Returns:
*/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Add("content-type", "application/json")

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
		return
	}

	// Get diagnose information.
	diagnose := getDiagnose()

	// Encode the diagnose information as JSON and send it in the response.
	encoder := json.NewEncoder(w)
	err := encoder.Encode(diagnose)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}
}

/*
getDiagnose returns a diagnose struct containing information about the uptime and status of the universities and countries APIs.
Returns: a diagnose struct containing information about the uptime and status of the universities and countries APIs.
*/
func getDiagnose() predefined.Diagnose {

	// Check the status of the universities API.
	url := predefined.UNIVERSITIESAPI_URL
	universityApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("There was an error in creating university API request: %e", err.Error())
	}
	universityApiRequest.Header.Add("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(universityApiRequest)
	if err != nil {
		fmt.Errorf("There was an error in university API response: %e", err.Error())
	}
	universityApiStatus := res.StatusCode

	// Check the status of the countries API.
	url = predefined.COUNTRIESAPI_URL + "all"
	countriesApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("There was an error in creating country API request: %e", err.Error())
	}
	res, err = client.Do(countriesApiRequest)
	if err != nil {
		fmt.Errorf("There was an error in countries API response: %e", err.Error())
	}
	countriesApiStatus := res.StatusCode

	// Return a diagnose struct containing information about the uptime and status of the universities and countries APIs.
	return predefined.Diagnose{
		UniversitiesApi: fmt.Sprintf("%d", universityApiStatus),
		CountriesApi:    fmt.Sprintf("%d", countriesApiStatus),
		Version:         "v1",
		Uptime:          uptime.GetUptime(),
	}
}
