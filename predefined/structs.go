package predefined

// This file defines structs and a function to work with data related to universities and countries.
import "time"

// Diagnose represents diagnostic information about the API server.
type Diagnose struct {
	UniversitiesApi string `json:"universitiesapi"`
	CountriesApi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          int    `json:"uptime"`
}

// University represents data about a university.
type University struct {
	Name         string   `json:"name"`
	Country      string   `json:"country"`
	AlphaTwoCode string   `json:"alpha_two_code"`
	WebPages     []string `json:"web_pages"`
}

// Country represents data about a country.
type Country struct {
	Name      map[string]interface{} `json:"name"`
	Languages map[string]string      `json:"languages"`
	Maps      map[string]string      `json:"maps"`
	Borders   []string               `json:"borders"`
	CCA3      string                 `json:"cca3"`
	CCA2      string                 `json:"cca2"`
	Cache     time.Time
}

// UniversityAndCountry struct that contains data about a university and its country.
type UniversityAndCountry struct {
	Name      string            `json:"name,omitempty"`
	Country   string            `json:"country,omitempty"`
	Isocode   string            `json:"isocode,omitempty"`
	WebPages  []string          `json:"webpages,omitempty"`
	Languages map[string]string `json:"languages,omitempty"`
	Map       string            `json:"map,omitempty"`
}

/*
CombineUniversityAndCountry combines data about a university and a country.
Returns: a UniversityAndCountry object
*/
func CombineUniversityAndCountry(u University, c Country, fields ...string) UniversityAndCountry {
	var universityInformation UniversityAndCountry
	if len(fields) > 0 {
		for _, f := range fields {
			switch f {
			case "name":
				universityInformation.Name = u.Name

			case "country":
				universityInformation.Country = u.Country

			case "isocode":
				universityInformation.Isocode = u.AlphaTwoCode

			case "webpages":
				universityInformation.WebPages = u.WebPages

			case "languages":
				universityInformation.Languages = c.Languages

			case "map":
				universityInformation.Map = c.Maps["openStreetMaps"]
			}
		}
	} else {
		universityInformation = UniversityAndCountry{
			Name:      u.Name,
			Country:   u.Country,
			Isocode:   u.AlphaTwoCode,
			WebPages:  u.WebPages,
			Languages: c.Languages,
			Map:       c.Maps["openStreetMaps"],
		}
	}
	return universityInformation

}
