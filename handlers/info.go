package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
)

// HandleInfo handles the country info request and returns data in JSON format.
func HandleInfo(w http.ResponseWriter, r *http.Request) error {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters and default values
	isoCode := r.PathValue("two_letter_country_code")
	cityLimitStr := r.URL.Query().Get("limit")

	// Fetch country info and handle errors
	info, err := getCountryInfo(isoCode, cityLimitStr)
	if err != nil {
		return err
	}

	// Marshal the response into JSON
	response, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// Send the JSON response with HTTP status OK
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		return err
	}

	return nil
}

// getCountryInfo fetches country data from an external API.
func getCountryInfo(isoCode, cityLimitStr string) (utils.CountryInfo, error) {
	url := utils.REST_COUNTRIES_API_URL + isoCode
	log.Printf("Fetching data from API: %s for country code: %s with limit %s", url, isoCode, cityLimitStr)

	// Make HTTP request to the external API
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return utils.CountryInfo{}, errors.New("error contacting API")
	}
	defer resp.Body.Close()
	/**
	Filter
	GET
	Get only specified countries information
	https://countriesnow.space/api/v0.1/countries/info?returns=currency,flag,unicodeFlag,dialCode
	Get only specified countries information
	*/

	// Ensure the response is successful
	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned status code: %d", resp.StatusCode)
		return utils.CountryInfo{}, errors.New("API returned error status code")
	}

	// Decode the JSON response into the appropriate struct
	var apiResponse []utils.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return utils.CountryInfo{}, errors.New("error decoding JSON")
	}

	if len(apiResponse) == 0 {
		log.Println("No data found for country")
		return utils.CountryInfo{}, errors.New("no data found for country")
	}

	// Extract country data from the response
	country := apiResponse[0]
	info := utils.CountryInfo{
		Name:       country.Name.Common,
		Continents: country.Continents,
		Population: country.Population,
		Languages:  country.Languages,
		Borders:    country.Borders,
		Flag:       country.Flag,
		Capital:    country.Capital[0], // Assuming there's always at least one capital
		Cities:     nil,                // Placeholder for cities
	}

	// Fetch cities based on the country code and limit
	APIResponse, err := fetchCitiesFromAPI(isoCode)
	if err != nil {
		return utils.CountryInfo{}, errors.New("error fetching cities")
	}
	cities := limitCities(APIResponse.Data, 10)
	info.Cities = cities

	return info, nil
}

// limitCities ensures the returned list of cities does not exceed the given limit.
func limitCities(cities []string, limit int) []string {
	if len(cities) > limit {
		return cities[:limit]
	}
	return cities
}

// APIResponse defines the expected structure of the API response.
type APIResponse struct {
	Error   bool     `json:"error"`
	Message string   `json:"msg"`
	Data    []string `json:"data"`
}

// fetchCitiesFromAPI sends a request to the Countries-Now API to fetch city data.
func fetchCitiesFromAPI(isoCode string) (*APIResponse, error) {
	url := utils.COUNTRIES_NOW_API_URL + "countries/cities"
	log.Println("Fetching city data from API:", url)

	requestBody, err := json.Marshal(map[string]string{"iso2": isoCode})
	if err != nil {
		return nil, errors.New("failed to encode request payload")
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("failed to reach Countries-Now API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Countries-Now API returned error status code: %d", resp.StatusCode)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, errors.New("failed to decode Countries-Now API response")
	}

	if apiResponse.Error {
		return nil, errors.New("Countries-Now API returned an error: " + apiResponse.Message)
	}

	return &apiResponse, nil
}
