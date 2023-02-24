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

func HandlerNeighbourUniversities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Other methods than GET are not supported.", http.StatusMethodNotAllowed)
		return
	}

	pathList := strings.Split(r.URL.Path, "/")
	countryQuery := pathList[len(pathList)-2]
	uniQuery := pathList[len(pathList)-1]

	var fields []string
	if r.URL.Query().Get("fields") != "" {
		fields = strings.Split(r.URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}

	var limit int
	if r.URL.Query()["limit"] != nil {
		if l, err := strconv.Atoi(r.URL.Query()["limit"][0]); err != nil || l < 0 {
			http.Error(w, "The limit is not a valid positive integer. Using 0 as limit.", http.StatusBadRequest)
			limit = constants.LIMIT_DEFAULT
		} else {
			limit = l
		}
	} else {
		limit = constants.LIMIT_DEFAULT
	}

	country, err := requests.RequestCountryInformation(fmt.Sprintf("%sname/%s?fields=borders", constants.COUNTRIESAPI_URL, countryQuery))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	var countries []structs.Country
	for _, bor := range country.Borders {
		url := fmt.Sprintf("%salpha/%s?fields=name,languages,maps", constants.COUNTRIESAPI_URL, bor)
		c, _ := requests.RequestCountryInformation(url)
		countries = append(countries, c)
	}

	var universityInformation []structs.UniversityAndCountry
	for _, c := range countries {
		query := fmt.Sprintf("search?name=%s&country=%s", uniQuery, c.Name["common"].(string))
		url := constants.UNIVERSITIESAPI_URL + query
		universities, _ := requests.RequestUniversityInformation(url)
		for _, u := range universities {
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
