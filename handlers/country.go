package handler

import (
	"encoding/json"
	"errors"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// GetInfo handles the country info request and returns data in JSON format.
func GetInfo(w http.ResponseWriter, r *http.Request) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters and default values
	isoCode := strings.TrimPrefix(r.URL.Path, "/country/v1/info/")
	cityLimitStr := r.URL.Query().Get("limit")
	if cityLimitStr == "" {
		cityLimitStr = "3" // Default to 3 cities if no limit provided // TODO: Wrong with default value?
	}

	// Fetch country info and handle errors
	info, err := getCountryInfo(isoCode, cityLimitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the response into JSON
	response, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Send the JSON response with HTTP status OK
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
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
	cities, err := getCities(isoCode, cityLimitStr)
	if err != nil {
		log.Printf("Error fetching cities: %v", err)
		return utils.CountryInfo{}, errors.New("error fetching cities")
	}
	info.Cities = cities

	return info, nil
}

// getCities fetches the top cities for the given country, with an optional limit on the number of cities.
func getCities(isoCode, limitStr string) ([]string, error) {
	// Default limit to 3 cities if not provided
	limit := 3
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			log.Printf("Invalid limit value: %s, defaulting to 3", limitStr)
		} else {
			limit = parsedLimit
		}
	}

	// Fetch cities (stubbed for now)
	// TODO: Integrate with a real API for fetching cities
	cities := []string{"Oslo", "Bergen", "Trondheim"}
	if len(cities) > limit {
		cities = cities[:limit] // Slice to limit the number of cities
	}
	return cities, nil
}
