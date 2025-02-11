package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
)

func HandlePopulation(w http.ResponseWriter, r *http.Request) error {
	url := utils.COUNTRIES_NOW_API_URL + "countries/population"

	isoCode := r.PathValue("two_letter_country_code")
	log.Printf("Fetching data from API: %s for country code: %s", url, isoCode)
	//limit := r.URL.Query().Get("limit")

	w.Header().Set("Content-Type", "application/json")

	// Create JSON payload
	requestBody, err := json.Marshal(map[string]string{
		"iso3": "NOR", // TODO: API endpoint does not support ISO2 codes, only ISO3 and country names
	}) // TODO: Endpoint must support iso2 codes, and must therefor translate iso2 to iso3/country name
	if err != nil {
		return err
	}

	// Make HTTP request to the external API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Decode the response
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	// Check for errors in the response
	if data["error"] == true {
		return fmt.Errorf("API error: %s", data["msg"])
	}

	population := utils.PopulationInfo{
		Mean: 5000000,
		Values: []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		}{
			{Year: 2000, Value: 4500000},
			{Year: 2010, Value: 4800000},
			{Year: 2020, Value: 5300000},
		},
	}

	// Marshal the response into JSON
	response, err := json.Marshal(population)
	if err != nil {
		return err
	}

	// Send the JSON response with HTTP status OK
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		return err
	}

	// https://countriesnow.space/api/v0.1/countries/population/filter
	// Try using the link above with a query for country and given years

	return nil
}
