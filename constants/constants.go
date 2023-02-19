package constants

// This file defines constants used throughout the program.
const (
	// PORT Default port. If the port is not set by environment variables, set the port.
	PORT = "8080"

	// The paths that will be handles by each handler
	DEFAULT_PATH       = "/unisearcher/"
	UNIINFO_PATH       = "/unisearcher/v1/uniinfo/"
	NEIGHBOURUNIS_PATH = "/unisearcher/v1/neighbourunis/"
	DIAG_PATH          = "/unisearcher/v1/diag/"

	// The URLS to the different API's
	UNIVERSITIESAPI_URL = "http://universities.hipolabs.com/"
	COUNTRIESAPI_URL    = "https://restcountries.com/v3.1/"
)
