package utility

// File containing helper_functions to retrieve queries from URLs.
import (
	"assignment-1/constants"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

/*
GetLimit Returns the setLimit from the query in the URL, or an error
Parameter URL: the URL containing the limit to be retrieved from
*/
func GetLimit(URL *url.URL) (int, error) {
	var errorReturned error
	var setLimit int
	if URL.Query()["limit"] != nil {
		if l, err := strconv.Atoi(URL.Query()["limit"][0]); err != nil || l < 0 {
			errorReturned = errors.New("the value that was set is not a positive integer. " +
				"0 will be used as the limit")
			setLimit = constants.LIMIT_DEFAULT
		} else {
			setLimit = l
		}
	} else {
		setLimit = constants.LIMIT_DEFAULT
	}
	return setLimit, errorReturned
}

/*
GetFields Returns the specified fields from the query in the URL, or an error
Parameter URL: the URL containing the fields to retrieve information from
*/
func GetFields(URL *url.URL) []string {
	var fields []string
	if URL.Query().Get("fields") != "" {
		fields = strings.Split(URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}
	return fields
}
