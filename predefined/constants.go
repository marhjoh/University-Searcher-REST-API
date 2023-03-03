package predefined

// This file defines constants used throughout the program.
const (
	// PORT Default port. If the port is not set by environment variables, set the port.
	PORT = "8080"

	// The paths that will be handled by each handler
	UNIVERSITYINFORMATION_PATH = "/unisearcher/v1/uniinfo/"
	NEIGHBOURUNIVERSITIES_PATH = "/unisearcher/v1/neighbourunis/"
	DIAG_PATH                  = "/unisearcher/v1/diag/"

	// The URLS to the different API's
	UNIVERSITIESAPI_URL = "http://universities.hipolabs.com/"
	COUNTRIESAPI_URL    = "https://restcountries.com/v3/"

	// Limit constants
	LIMIT_DEFAULT = 0
	LIMIT_HOURS   = 600

	// The standard fields to request from the country-API
	COUNTRIESAPI_STANDARD_FIELDS = "name,languages,maps,borders,cca2,cca3"
)
