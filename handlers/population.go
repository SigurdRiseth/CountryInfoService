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
)

type IsoCodeResponse struct {
	Cca3 string `json:"cca3"`
}

type populationAPIResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string            `json:"country"`
		Code             string            `json:"code"`
		Iso3             string            `json:"iso3"`
		PopulationCounts []utils.YearValue `json:"populationCounts"`
	} `json:"data"`
}

func HandlePopulation(w http.ResponseWriter, r *http.Request) error {
	url := utils.COUNTRIES_NOW_API_URL + "countries/population"

	isoCode := r.PathValue("two_letter_country_code")
	limit := r.URL.Query().Get("limit")
	log.Printf("Fetching data from API: %s for country code: %s with limit %s", url, isoCode, limit)

	var err error

	// Now you can access the ISO code
	isoCode, err = getIso3(isoCode)
	if err != nil {
		return fmt.Errorf("countries ISO3 not found")
	}
	log.Printf("ISO Code: %s", isoCode)

	w.Header().Set("Content-Type", "application/json")

	// Create JSON payload
	requestBody, err := json.Marshal(map[string]string{
		"iso3": isoCode,
	})
	if err != nil {
		return err
	}

	// Make HTTP request to the external API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var apiResponse populationAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return err
	}

	// Filter data based on limit
	filteredValues, err := filterByYearLimit(apiResponse.Data.PopulationCounts, limit)
	if err != nil {
		log.Printf("Error filtering data: %v", err)
		return err
	}

	// Process the data to match your PopulationInfo struct
	population := utils.PopulationInfo{
		Values: filteredValues,
	}

	// Calculate the mean population
	sum := 0
	for _, v := range filteredValues {
		sum += v.Value
	}
	if len(filteredValues) > 0 {
		population.Mean = sum / len(filteredValues)
	}

	// Use the PopulationInfo struct as needed
	log.Printf("Population Info: %+v", population)

	// Send the response
	json.NewEncoder(w).Encode(population)

	return nil
}

func getIso3(isoCode string) (string, error) {
	isoCodeResponse, err := http.Get(utils.REST_COUNTRIES_API_URL + isoCode + "?fields=cca3")
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return "", errors.New("error contacting API")
	}
	defer isoCodeResponse.Body.Close()

	// Check if the response status is 200 OK
	if isoCodeResponse.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", isoCodeResponse.StatusCode)
	}

	var result IsoCodeResponse
	if err := json.NewDecoder(isoCodeResponse.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}
	return result.Cca3, nil
}

func filterByYearLimit(values []utils.YearValue, limit string) ([]utils.YearValue, error) {
	if limit == "" {
		return values, nil
	}

	// Split the limit into start and end year
	years := strings.Split(limit, "-")
	if len(years) != 2 {
		return nil, fmt.Errorf("invalid limit format, expected 'start-end'")
	}

	// Convert to integers
	startYear, err := strconv.Atoi(years[0])
	if err != nil {
		return nil, fmt.Errorf("invalid start year")
	}
	endYear, err := strconv.Atoi(years[1])
	if err != nil {
		return nil, fmt.Errorf("invalid end year")
	}

	// Filter the data
	var filtered []utils.YearValue
	for _, v := range values {
		if v.Year >= startYear && v.Year <= endYear {
			filtered = append(filtered, v)
		}
	}
	return filtered, nil
}
