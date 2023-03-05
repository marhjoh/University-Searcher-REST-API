package contextual_error_messages

import (
	"errors"
)

func GetEncodingError() error {
	return errors.New("There were an error during encoding.")
}

func GetInvalidLimitError() error {
	return errors.New("The limit that was set is not a valid positive integer. The limit has been set as 0.")
}

func GetLimitNotSuppliedError() error {
	return errors.New("There were no limit parameter supplied. The limit has been set as 0.")
}

func GetCountriesNotFoundError() error {
	return errors.New("There were no countries found.")
}

func GetUniversitiesNotFoundError() error {
	return errors.New("There were no universities found.")
}
