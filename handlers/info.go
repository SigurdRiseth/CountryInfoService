package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
	"strconv"
)

// Country represents the structure of the API response
type Country struct {
	Name       Name              `json:"name"`
	Capital    []string          `json:"capital"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Population int               `json:"population"`
	Continents []string          `json:"continents"`
}

// Name represents the naming details of the country
type Name struct {
	Common     string                `json:"common"`
	Official   string                `json:"official"`
	NativeName map[string]NativeName `json:"nativeName"`
}

// NativeName represents the native language details
type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

// HandleInfo processes requests to retrieve country information based on an ISO2 country code.
// It fetches details about the country, including cities, with an optional limit on the number of cities returned.
//
// Request Parameters:
//   - "two_letter_country_code" (path parameter): The ISO2 country code (e.g., "US" for the United States).
//   - "limit" (query parameter, optional): A string representing the number of cities to retrieve.
//
// Response:
//   - Returns a JSON response containing country information or an error message.
//
// HTTP Status Codes:
//   - 200 OK: Successfully retrieved country information.
//   - 400 Bad Request: Invalid input parameters.
//   - 500 Internal Server Error: Failed to retrieve country information.
//
// Example Usage:
//
//	GET /info/US?limit=10 -> Retrieves information about the United States with a limit of 10 cities.
//
// Returns an error if fetching country information fails.
func HandleInfo(w http.ResponseWriter, r *http.Request) error {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters and default values
	isoCode := r.PathValue("two_letter_country_code")
	cityLimitStr := r.URL.Query().Get("limit")

	// Fetch country info and handle errors
	info, err := getCountryInfo(isoCode, cityLimitStr)

	// Construct response based on error presence
	apiResponse := utils.APIResponse{
		Error:   err != nil,
		Message: "Country information retrieved successfully",
		Data:    info,
	}

	// Set error message if an error occurred
	if err != nil {
		apiResponse.Message = err.Error()
		apiResponse.Data = nil
	}

	// Marshal and send the JSON response
	response, _ := json.Marshal(apiResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return err
}

// getCountryInfo retrieves country information from an external API based on the provided ISO2 country code.
// It fetches details such as country name, continents, population, languages, borders, flag, capital, and cities.
//
// Parameters:
//   - isoCode (string): The ISO2 country code (e.g., "US" for the United States).
//   - cityLimitStr (string): A string representing the maximum number of cities to retrieve (optional).
//
// Returns:
//   - utils.CountryInfo: A struct containing the retrieved country information.
//   - error: An error if the request fails or the response cannot be processed.
//
// Function Workflow:
//   - Constructs the API request URL.
//   - Makes an HTTP GET request to fetch country data.
//   - Decodes the JSON response into a Country struct.
//   - Extracts relevant country data into a utils.CountryInfo struct.
//   - Fetches cities using an additional API call and applies an optional limit.
//
// Errors:
//   - Returns an error if the API request fails, the response is invalid, or JSON decoding fails.
//
// Example Usage:
//
//	info, err := getCountryInfo("US", "10")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(info)
func getCountryInfo(isoCode, cityLimitStr string) (utils.CountryInfo, error) {
	// Construct the URL for the external API
	url := utils.RestCountriesApiUrl + isoCode + utils.RestCountriesFilter
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
		log.Printf("Rest-Countries-API returned status code: %d", resp.StatusCode)
		return utils.CountryInfo{}, fmt.Errorf("API returned error status code %d", resp.StatusCode)
	}

	// Decode the JSON response into the appropriate struct
	var apiResponse Country
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return utils.CountryInfo{}, errors.New("error decoding JSON")
	}

	// Extract country data from the response
	country := apiResponse
	info := utils.CountryInfo{
		Name:       country.Name.Common,
		Continents: country.Continents,
		Population: country.Population,
		Languages:  country.Languages,
		Borders:    country.Borders,
		Flag:       country.Flag,
		Capital:    country.Capital[0],
		Cities:     nil,
	}

	// Fetch cities based on the country code and limit
	citiesFromAPI, err := fetchCitiesFromAPI(isoCode)
	if err != nil {
		return utils.CountryInfo{}, errors.New("error fetching cities")
	}
	cities := limitCities(citiesFromAPI.Data, cityLimitStr)
	info.Cities = cities

	return info, nil
}

// limitCities limits the number of cities returned based on a given limit string.
// If the limit is invalid or not provided, a default limit is used.
//
// Parameters:
//   - cities ([]string): A slice of city names to be limited.
//   - limitString (string): The limit as a string; expected to be a positive integer.
//
// Returns:
//   - []string: A slice of cities, limited to the specified number.
//
// Behavior:
//   - Converts `limitString` to an integer. If conversion fails or limit is non-positive, defaults to `utils.DefaultCityLimit`.
//   - Ensures that the limit does not exceed the available number of cities.
//   - Returns the original slice if its length is less than or equal to the limit.
//
// Example Usage:
//
//	cities := []string{"Oslo", "Bergen", "Trondheim", "Stavanger"}
//	limitedCities := limitCities(cities, "2")  // Output: ["Oslo", "Bergen"]
func limitCities(cities []string, limitString string) []string {
	// Convert limitString to an integer, fallback to defaultLimit on error
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit <= 0 {
		log.Printf("Invalid limit value: %s, defaulting to %d", limitString, utils.DefaultCityLimit)
		limit = utils.DefaultCityLimit
	}

	// Ensure limit does not exceed the length of cities
	if len(cities) > limit {
		return cities[:limit]
	}
	return cities
}

// fetchCitiesFromAPI fetches a list of cities for a given country ISO2 code from the Countries-Now API.
//
// Parameters:
//   - isoCode (string): The two-letter country code (ISO2).
//
// Returns:
//   - *utils.APIResponseString: A pointer to the API response containing city data if successful, otherwise nil.
//   - error: An error if the request fails or the response contains an error message.
//
// Behavior:
//   - Constructs a request payload with the given `isoCode`.
//   - Sends a POST request to the Countries-Now API to retrieve city data.
//   - Handles potential errors, including JSON encoding/decoding issues and API response failures.
//   - Checks if the API response contains an error and returns an appropriate error message.
//
// Example Usage:
//
//	response, err := fetchCitiesFromAPI("NO")
//	if err != nil {
//	    log.Println("Error fetching cities:", err)
//	} else {
//	    log.Println("Cities:", response.Data)
//	}
func fetchCitiesFromAPI(isoCode string) (*utils.APIResponseString, error) {
	// Construct the URL for the external API
	url := utils.CountriesNowApiUrl + utils.CountriesNowCityEndpoint
	log.Println("Fetching city data from API:", url)

	// Construct the request payload
	requestBody, err := json.Marshal(map[string]string{"iso2": isoCode})
	if err != nil {
		return nil, errors.New("failed to encode request payload for city data")
	}

	// Make HTTP request to the external API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("failed to reach Countries-Now API for city data")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Countries-Now API returned error status code: %d", resp.StatusCode)
	}

	var apiResponse utils.APIResponseString
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, errors.New("failed to decode Countries-Now API response")
	}

	if apiResponse.Error {
		return nil, errors.New("Countries-Now API returned an error: " + apiResponse.Message)
	}

	return &apiResponse, nil
}
