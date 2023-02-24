package handlers

import (
	"assignment-1/cmd/handlers/constants"
	"assignment-1/structs"
	"assignment-1/uptime"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerDiag(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
		return
	}

	diagnose := getDiagnose()

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(diagnose)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}

}

func getDiagnose() structs.Diagnose {

	url := constants.UNIVERSITIESAPI_URL
	universityApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("There was an error in creating university API request: %e", err.Error())
	}
	universityApiRequest.Header.Add("content-type", "application/json")

	url = constants.COUNTRIESAPI_URL + "all"
	countriesApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("There was an error in creating country API request: %e", err.Error())
	}

	client := http.Client{}
	res, err := client.Do(universityApiRequest)
	if err != nil {
		fmt.Errorf("There was an error in university API response: %e", err.Error())
	}

	universityApiStatus := res.StatusCode

	res, err = client.Do(countriesApiRequest)
	if err != nil {
		fmt.Errorf("There was an error in countries API response: %e", err.Error())
	}

	countriesApiStatus := res.StatusCode

	return structs.Diagnose{
		UniversitiesApi: fmt.Sprintf("%d", universityApiStatus),
		CountriesApi:    fmt.Sprintf("%d", countriesApiStatus),
		Version:         "v1",
		Uptime:          uptime.GetUptime(),
	}

}
