package request_university

// File containing helper_functions to request universities from the university-API
import (
	"assignment-1/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

/*
RequestUniversityInformation Requests an array of universities with a URL using the University-API.
Param url: the URL to be requested
Returns: An array of universities, or both an error and an empty array.
*/
func RequestUniversityInformation(url string) ([]structs.University, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("content-type", "application/json")

	if err != nil {
		return []structs.University{}, err
	}
	client := &http.Client{}
	res, err := client.Do(r)

	switch {
	case res.StatusCode != http.StatusOK:
		return []structs.University{}, errors.New(fmt.Sprintf("The status code returned from the universityAPI: %d", res.StatusCode))
	case res.StatusCode == http.StatusNotFound:
		return []structs.University{}, errors.New(fmt.Sprintf("%d There were no university found", res.StatusCode))
	}

	decoder := json.NewDecoder(res.Body)

	var universities []structs.University
	if err := decoder.Decode(&universities); err != nil {
		log.Println(err)
	}

	return universities, nil
}
