package contextual_error_messages

import (
	"errors"
)

func GetErrorDuringEncodingError() error {
	return errors.New("There were an error during encoding.")
}

func GetNotValidLimitError() error {
	return errors.New("The limit that was set is not a valid positive integer. Limit set as 0.")
}

func GetNoLimitGivenError() error {
	return errors.New("There were no limit given.")
}

func GetNoCountriesFoundError() error {
	return errors.New("There were no countries found.")
}

func GetNoUniversitiesFoundError() error {
	return errors.New("There were no universities found.")
}
