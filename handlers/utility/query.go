package utility

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
	limit := URL.Query().Get("limit")
	if limit == "" {
		return predefined.LIMIT_DEFAULT, nil
	}

	setLimit, err := strconv.Atoi(limit)
	if err != nil || setLimit < 0 {
		return predefined.LIMIT_DEFAULT, contextual_error_messages.GetInvalidLimitError()
	}

	return setLimit, nil
}

/*
GetFields Returns the specified fields from the query in the URL, or an error
Parameter URL: the URL containing the fields to retrieve information from
Returns: A slice of strings containing the specified fields, or an empty slice.
*/
func GetFields(URL *url.URL) []string {
	if fields := URL.Query().Get("fields"); fields != "" {
		return strings.Split(fields, ",")
	}
	return nil
}
