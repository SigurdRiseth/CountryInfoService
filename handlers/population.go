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
	"strings"
	"time"
)

// IsoCodeResponse represents the response structure for looking up the ISO3 code
// from an ISO2 country code. This is used when making a request to the RestCountries API.
type IsoCodeResponse struct {
	Cca3 string `json:"cca3"`
}

// populationAPIResponse represents the structure of the population data response
// from the external API (e.g., CountriesNow API). It contains the country information
// and the population data for various years.
type populationAPIResponse struct { // TODO: USE struct in utils.structs.go instead?
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string            `json:"country"`
		Code             string            `json:"code"`
		Iso3             string            `json:"iso3"`
		PopulationCounts []utils.YearValue `json:"populationCounts"`
	} `json:"data"`
}

// HandlePopulation processes the population data for a given country based on its ISO2 country code.
//
// This function handles the full flow of fetching population data for a country:
//   - Converts the provided ISO2 country code to an ISO3 code using the `getIso3` function.
//   - Sends a POST request to an external API (CountriesNow API) with the ISO3 code to retrieve population data.
//   - Filters the population data based on the provided year range (if any).
//   - Calculates the mean population for the filtered data.
//   - Constructs a JSON response with the filtered data and mean population.
//
// If any step fails, an appropriate error response is returned to the client.
//
// Parameters:
//   - w: The `http.ResponseWriter` to send the response to the client.
//   - r: The `http.Request` that contains the ISO2 country code in the URL and the optional 'limit' query parameter for year range.
//
// Responses:
//   - If successful, a JSON response with the population data and the mean population is returned.
//   - If there is an error, an error response is returned with a relevant message and status code.
//
// Error Handling:
//   - BadRequest (400): If the ISO2 code cannot be converted to ISO3.
//   - InternalServerError (500): If there is an issue marshaling the JSON or contacting the external API.
//   - ServiceUnavailable (503): If the external API cannot be reached.
//   - UnprocessableEntity (422): If the provided year range is invalid.
//
// This function ensures that population data is accurately retrieved and returned to the client, with proper error handling at each step.
func HandlePopulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := utils.CountriesNowApiUrl + utils.CountriesNowPopulationEndpoint
	isoCode := r.PathValue("two_letter_country_code")
	limit := r.URL.Query().Get("limit")

	log.Printf("Fetching data from API: %s for country code: %s with limit %s", url, isoCode, limit)

	// Convert ISO2 to ISO3
	iso3, err := getIso3(isoCode)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Error fetching ISO3 code from ISO2: "+isoCode)
		return
	}

	// Create JSON payload
	requestBody, err := json.Marshal(map[string]string{"iso3": iso3})
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error marshalling JSON")
		return
	}

	// Make HTTP request to the external API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		handleError(w, http.StatusServiceUnavailable, "Error contacting external API")
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var apiResponse populationAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error decoding JSON response")
		return
	}

	// Filter population data
	filteredValues, err := filterByYearLimit(apiResponse.Data.PopulationCounts, limit)
	if err != nil {
		handleError(w, http.StatusUnprocessableEntity, "Invalid year range format. Expected 'startYear-endYear'.")
		return
	}

	// Calculate the mean population
	sum := 0
	for _, v := range filteredValues {
		sum += v.Value
	}
	mean := 0
	if len(filteredValues) > 0 {
		mean = sum / len(filteredValues)
	}

	// Construct response
	response := utils.APIResponse{
		Error:   false,
		Message: "Population data retrieved successfully",
		Data:    utils.PopulationInfo{Values: filteredValues, Mean: mean},
	}

	// Send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// handleError sends an error response to the client with the provided status code and message.
//
// This function logs the error message and then formats a JSON response with the provided status code,
// a message indicating the error, and a `nil` data field. The response is sent to the client via the
// `http.ResponseWriter`.
//
// Parameters:
//   - w: The `http.ResponseWriter` to send the response to the client.
//   - statusCode: The HTTP status code to be sent in the response (e.g., 400 for bad request, 500 for internal server error).
//   - message: A string containing the error message that will be included in the response body.
//
// Example usage:
//
//	handleError(w, http.StatusBadRequest, "Invalid country code provided")
//
// This function is useful for centralizing error handling and ensuring consistent error responses throughout the API.
func handleError(w http.ResponseWriter, statusCode int, message string) {
	log.Printf("Error: %s", message)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(utils.APIResponse{
		Error:   true,
		Message: message,
		Data:    nil,
	})
}

// getIso3 retrieves the ISO3 code corresponding to a given ISO2 code (two-letter country code).
//
// It makes an HTTP request to the Rest-Countries API to fetch the ISO3 code. The function
// includes error handling for failed HTTP requests, non-200 responses, and issues decoding the API response.
//
// The function has a timeout of 5 seconds to avoid hanging requests in case the external API is unresponsive.
//
// Parameters:
//   - isoCode: A string representing the two-letter ISO2 country code for which the ISO3 code is requested.
//
// Returns:
//   - A string representing the ISO3 code if the lookup is successful.
//   - An error if there is a problem making the request, receiving the response, or decoding the result.
//
// Example usage:
//
//	iso3Code, err := getIso3("NO") // Retrieves ISO3 code for Norway (NO)
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("ISO3 Code:", iso3Code)
//	}
//
// Error cases:
//   - If the request to the Rest-Countries API fails, an error is returned with a message indicating the issue.
//   - If the response status code is not 200 OK, an error is returned with the unexpected status code.
//   - If the response body cannot be decoded or contains an invalid ISO3 code, an error is returned.
func getIso3(isoCode string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second} // Timeout to prevent hanging requests

	url := utils.RestCountriesApiUrl + isoCode + utils.RestCountriesIso3Filter
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for ISO3 lookup (%s): %v", isoCode, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error contacting Rest-Countries API for %s: %v", isoCode, err)
		return "", errors.New("failed to reach Rest-Countries API")
	}
	defer resp.Body.Close()

	// Check if the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code (%d) from Rest-Countries API for %s", resp.StatusCode, isoCode)
	}

	// Decode JSON response
	var result IsoCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode JSON response for %s: %v", isoCode, err)
	}

	// Ensure the API response contains a valid ISO3 code
	if result.Cca3 == "" {
		return "", fmt.Errorf("empty ISO3 code received for %s", isoCode)
	}

	return result.Cca3, nil
}

// filterByYearLimit filters the given slice of YearValue structs based on a specified year range.
//
// If no limit is provided (empty string), the original slice of values is returned without any filtering.
//
// The `limit` parameter should be a string in the format 'start-end', where:
// - 'start' is the starting year (inclusive)
// - 'end' is the ending year (inclusive)
//
// If the `limit` format is invalid or the years are not integers, an error will be returned.
// Additionally, if the start year is greater than the end year, an error will be returned.
//
// Example usage:
//
//	values := []YearValue{
//	    {Year: 2000, Value: 100},
//	    {Year: 2005, Value: 150},
//	    {Year: 2010, Value: 200},
//	}
//	filteredValues, err := filterByYearLimit(values, "2000-2005")
//	// filteredValues will contain the YearValue structs for the years 2000 and 2005
//
// Parameters:
//   - values: a slice of YearValue structs to be filtered.
//   - limit: a string in the format 'start-end' specifying the year range to filter by.
//
// Returns:
//   - A slice of YearValue structs filtered by the specified year range.
//   - An error if the limit format is invalid or if other validation issues occur.
func filterByYearLimit(values []utils.YearValue, limit string) ([]utils.YearValue, error) {
	// If no limit is provided, return the original values
	if limit == "" {
		return values, nil
	}

	// Split the limit into start and end year
	years := strings.Split(limit, "-")
	if len(years) != 2 {
		return nil, fmt.Errorf("invalid limit format, expected 'start-end' but got '%s'", limit)
	}

	// Convert start year
	startYear, err := strconv.Atoi(years[0])
	if err != nil {
		return nil, fmt.Errorf("invalid start year '%s', expected an integer", years[0])
	}

	// Convert end year
	endYear, err := strconv.Atoi(years[1])
	if err != nil {
		return nil, fmt.Errorf("invalid end year '%s', expected an integer", years[1])
	}

	// Validate the year range
	if startYear > endYear {
		return nil, fmt.Errorf("invalid limit format, start year %d cannot be greater than end year %d", startYear, endYear)
	}

	// Filter the data based on the year range
	var filtered []utils.YearValue
	for _, v := range values {
		if v.Year >= startYear && v.Year <= endYear {
			filtered = append(filtered, v)
		}
	}

	// Return the filtered data
	return filtered, nil
}
