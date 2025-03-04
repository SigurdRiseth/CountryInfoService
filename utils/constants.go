package utils

// API information
const (
	ApiVersion       = "1.0"
	DefaultCityLimit = 3
)

// Endpoint paths
const (
	BasePath       = "/countryinfo/v1"
	InfoPath       = "/info/{two_letter_country_code}"
	PopulationPath = "/population/{two_letter_country_code}"
	StatusPath     = "/status"
)

// Countries-Now API
const (
	CountriesNowApiUrl             = "http://129.241.150.113:3500/api/v0.1/"
	CountriesNowPopulationEndpoint = "countries/population"
	CountriesNowCityEndpoint       = "countries/cities"
)

// RestCountries API
const (
	RestCountriesApiUrl     = "http://129.241.150.113:8080/v3.1/alpha/"
	RestCountriesFilter     = "?fields=name,continents,population,languages,borders,flag,capital"
	RestCountriesIso3Filter = "?fields=cca3"
)

func GetInfoPath(countryCode string) string {
	return BasePath + InfoPath + countryCode
}

func GetPopulationPath(countryCode string) string {
	return BasePath + PopulationPath + countryCode
}

func GetStatusPath() string {
	return BasePath + StatusPath
}
