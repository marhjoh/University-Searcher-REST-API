package contextual_error_messages

// This file contains functions that return error messages used in the program.
import (
	"errors"
	"fmt"
)

/*
GetEncodingError returns an error
Returns: an error indicating that there was an error during encoding.
*/
func GetEncodingError() error {
	return errors.New("There were an error during encoding.")
}

/*
GetInvalidLimitError returns an error
Returns: an error indicating that the limit set is invalid.
*/
func GetInvalidLimitError() error {
	return errors.New("The limit that was set is not a valid positive integer.")
}

/*
GetCountriesNotFoundError returns an error
Returns: an error indicating that no countries were found.
*/
func GetCountriesNotFoundError() error {
	return errors.New("There were no countries found.")
}

/*
GetUniversitiesNotFoundError returns an error
Returns: an error indicating that no universities were found.
*/
func GetUniversitiesNotFoundError() error {
	return errors.New("There were no universities found.")
}

/*
GetInvalidNeighbourUniversityRequestError returns an error
Returns: an error indicating an invalid request made for getting neighbouring universities.
It also provides information on the valid format and a reference to the documentation.
*/
func GetInvalidNeighbourUniversityRequestError() error {
	return errors.New(fmt.Sprintf("The request that was made is not a valid request. "+
		"Format: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?{fields={:field1,field2,...}}&{limit={:number}}}.\n\n"+
		"%s", GetDocumentationError().Error()))
}

/*
GetInvalidMethodError returns an error
Returns: an error indicating that the method used is invalid.
*/
func GetInvalidMethodError() error {
	return errors.New("Method is not supported. Currently only GET are supported.")
}

/*
GetDocumentationError returns an error
Returns: an error indicating that the user should check the README file.
*/
func GetDocumentationError() error {
	return errors.New(fmt.Sprintf("Please check the README for further description" +
		"\nhttps://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2023-workspace/marhjoh/assignment-1/-/blob/master/README.md"))
}
