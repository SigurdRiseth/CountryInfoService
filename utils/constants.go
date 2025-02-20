package utils

// Define the paths for the different endpoints
const (
	InfoPath            = "/country/v1/info/{two_letter_country_code}"
	PopulationPath      = "/country/v1/population/{two_letter_country_code}"
	StatusPath          = "/country/v1/status"
	ApiVersion          = "1.0"
	CountriesNowApiUrl  = "http://129.241.150.113:3500/api/v0.1/"
	RestCountriesApiUrl = "http://129.241.150.113:8080/v3.1/alpha/"
	DefaultCityLimit    = 3
	RestCountriesFilter = "?fields=name,continents,population,languages,borders,flag,capital"
)
