package contextual_error_messages

import (
	"errors"
	"fmt"
)

func GetEncodingError() error {
	return errors.New("There were an error during encoding.")
}

func GetInvalidLimitError() error {
	return errors.New("The limit that was set is not a valid positive integer. The limit has been set as 0.")
}

func GetCountriesNotFoundError() error {
	return errors.New("There were no countries found.")
}

func GetUniversitiesNotFoundError() error {
	return errors.New("There were no universities found.")
}

func GetInvalidNeighbourUniversityRequestError() error {
	return errors.New(fmt.Sprintf("The request that was made is not a valid request. "+
		"Format: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?{fields={:field1,field2,...}}&{limit={:number}}}.\n\n"+
		"%s", GetDocumentationError().Error()))
}

func GetInvalidMethodError() error {
	return errors.New("Method is not supported. Currently only GET are supported.")
}

func GetDocumentationError() error {
	return errors.New(fmt.Sprintf("Please check the README for further description" +
		"\nhttps://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2023-workspace/marhjoh/assignment-1/-/blob/master/README.md"))
}
