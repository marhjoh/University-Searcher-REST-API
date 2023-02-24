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

func HandlerUniversityInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
	}

	universities := getUniversityInformation(r)

	if len(universities) == 0 {
		http.Error(w, "No universities were found", http.StatusNoContent)
		return
	}

	var fields []string
	if r.URL.Query().Get("fields") != "" {
		fields = strings.Split(r.URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}

	combinedInfo := getCombined(universities, fields)

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(combinedInfo)
	if err != nil {
		http.Error(w, "There was an error during encoding", http.StatusInternalServerError)
		return
	}

}

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
