package handlers

import (
	"assignment-1/cmd/handlers/constants"
	"assignment-1/structs"
	"assignment-1/uptime"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
HandlerDiag is a handler for the /diag endpoint.
Param w: The http.ResponseWriter
Param r: The http.Request pointer.
*/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
		return
	}

	// Get diagnose information.
	diagnose := getDiagnose()

	// Set response content type to JSON.
	w.Header().Add("content-type", "application/json")

	// Encode the diagnose information as JSON and send it in the response.
	encoder := json.NewEncoder(w)
	err := encoder.Encode(diagnose)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}
}

/*
getDiagnose returns a Diagnose struct containing information about the uptime and status of the universities and countries APIs.
*/
func getDiagnose() structs.Diagnose {

	// Check the status of the universities API.
	url := constants.UNIVERSITIESAPI_URL
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
	url = constants.COUNTRIESAPI_URL + "all"
	countriesApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("There was an error in creating country API request: %e", err.Error())
	}
	res, err = client.Do(countriesApiRequest)
	if err != nil {
		fmt.Errorf("There was an error in countries API response: %e", err.Error())
	}
	countriesApiStatus := res.StatusCode

	// Return a Diagnose struct containing information about the uptime and status of the universities and countries APIs.
	return structs.Diagnose{
		UniversitiesApi: fmt.Sprintf("%d", universityApiStatus),
		CountriesApi:    fmt.Sprintf("%d", countriesApiStatus),
		Version:         "v1",
		Uptime:          uptime.GetUptime(),
	}
}
