package utility

// File containing helper_functions to retrieve queries from URLs.
import (
	"assignment-1/contextual_error_messages"
	"assignment-1/predefined"
	"net/url"
	"strconv"
	"strings"
)

/*
GetLimit Returns the setLimit from the query in the URL, or an error
Parameter URL: the URL containing the limit to be retrieved from
Returns: the value of limit (int), and an error if the limit is not applicable.
*/
func GetLimit(URL *url.URL) (int, error) {
	var errorReturned error
	var setLimit int
	if URL.Query()["limit"] != nil {
		if l, err := strconv.Atoi(URL.Query()["limit"][0]); err != nil || l < 0 {
			errorReturned = contextual_error_messages.GetInvalidLimitError()
			setLimit = predefined.LIMIT_DEFAULT
		} else {
			setLimit = l
		}
	} else {
		setLimit = predefined.LIMIT_DEFAULT
	}
	return setLimit, errorReturned
}

/*
GetFields Returns the specified fields from the query in the URL, or an error
Parameter URL: the URL containing the fields to retrieve information from
Returns: A slice of strings containing the specified fields, or an empty slice.
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
