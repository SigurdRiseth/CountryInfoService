package utils

import "fmt"

// Country struct for displaying the country information
type Country struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int64             `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

// PopulationHistory struct, utilizing yearValue struct for displaying a country's population history
type PopulationHistory struct {
	Mean   int64       `json:"mean"`
	Values []yearValue `json:"values"`
}

// yearValue struct for displaying the year and corresponding population value
type yearValue struct {
	Year  int   `json:"year"`
	Value int64 `json:"value"`
}

// APIStatus struct for displaying the status of the APIs
type APIStatus struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           string `json:"uptime"`
}

// NewAPIStatus creates a new APIStatus instance with the provided API statuses and uptime.
//
// Parameters:
// - CountriesNowAPI: The status of the CountriesNow API.
// - RestCountriesAPI: The status of the RestCountries API.
// - uptime: The uptime of the API in seconds.
//
// Returns:
// - A pointer to an APIStatus struct containing the provided information.
func NewAPIStatus(CountriesNowAPI, RestCountriesAPI string, uptime float64) *APIStatus {
	return &APIStatus{
		CountriesNowAPI:  CountriesNowAPI,
		RestCountriesAPI: RestCountriesAPI,
		Version:          ApiVersion,
		Uptime:           fmt.Sprintf("%.0f seconds", uptime),
	}
}

type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"` // Use map for key-value pairs
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

// ErrorResponse represents the JSON error message structure
type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type PopulationInfo struct {
	Mean   int         `json:"mean"`
	Values []YearValue `json:"values"`
}

type YearValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type APIResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Allows any type of data
}

type APIResponseString struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Data    []string `json:"data"` // Allows any type of data
}
