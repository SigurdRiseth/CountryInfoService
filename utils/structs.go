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

// CountryInfo struct for displaying the country information
type APIResponse struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName map[string]struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nativeName"`
	} `json:"name"`
	TLD         []string `json:"tld"`
	CCA2        string   `json:"cca2"`
	CCN3        string   `json:"ccn3"`
	CCA3        string   `json:"cca3"`
	CIOC        string   `json:"cioc"`
	Independent bool     `json:"independent"`
	Status      string   `json:"status"`
	UNMember    bool     `json:"unMember"`
	Currencies  map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	IDD struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital      []string          `json:"capital"`
	AltSpellings []string          `json:"altSpellings"`
	Region       string            `json:"region"`
	Subregion    string            `json:"subregion"`
	Languages    map[string]string `json:"languages"`
	Translations map[string]struct {
		Official string `json:"official"`
		Common   string `json:"common"`
	} `json:"translations"`
	LatLng     []float64 `json:"latlng"`
	Landlocked bool      `json:"landlocked"`
	Borders    []string  `json:"borders"`
	Area       float64   `json:"area"`
	Demonyms   map[string]struct {
		F string `json:"f"`
		M string `json:"m"`
	} `json:"demonyms"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int                `json:"population"`
	Gini       map[string]float64 `json:"gini"`
	FIFA       string             `json:"fifa"`
	Car        struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
		Alt string `json:"alt"`
	} `json:"flags"`
	CoatOfArms struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
	} `json:"coatOfArms"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		LatLng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}
