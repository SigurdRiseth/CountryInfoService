package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	isoCode := "no" //r.URL.Query().Get("two_letter_country_code")
	cityLimitStr := r.URL.Query().Get("limit")

	info := getCountryInfo(isoCode, cityLimitStr)

	response, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func getCountryInfo(isoCode string, cityLimitStr string) utils.CountryInfo { // TODO: add return of error
	url := REST_COUNTRIES_API_URL + isoCode // TODO: Get country code from query parameter
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return utils.CountryInfo{} // Return empty struct if API fails
	}
	defer resp.Body.Close()

	// Check for valid response
	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned status code: %d", resp.StatusCode)
		return utils.CountryInfo{}
	}

	// Decode JSON response
	var apiResponse []utils.APIResponse // The API returns an array
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return utils.CountryInfo{}
	}

	if len(apiResponse) == 0 {
		log.Println("No data found for country.")
		return utils.CountryInfo{}
	}

	country := apiResponse[0] // Extract first result

	// Convert API response to CountryInfo struct
	info := utils.CountryInfo{
		Name:       country.Name.Common,
		Continents: country.Continents,
		Population: country.Population,
		Languages:  country.Languages,
		Borders:    country.Borders,
		Flag:       country.Flag,
		Capital:    country.Capital[0],
		Cities:     nil, // The API does not provide cities
	}

	info.Cities = getCities(isoCode, cityLimitStr)

	return info
}

func getCities(isoCode string, limit string) []string {
	if limit == "" {
		limit = "3" // Default to 3 cities
	}
	return []string{"Oslo", "Bergen", "Trondheim"} // TODO: Implement city fetching through CountriesNow API
}
