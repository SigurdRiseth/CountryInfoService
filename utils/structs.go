package utils

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
	CountriesNowAPI  int     `json:"countriesnowapi"`
	RestCountriesAPI int     `json:"restcountriesapi"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}
