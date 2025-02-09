# CountryInfoService

**CountryInfoService** is a simple REST API that provides country-related information, including the status of external country APIs. The service returns details such as the uptime of the service, API status, and version.

## Features
- **Country Information**: Get the name, capital, population, cities of a given country.
  - Optional: Limit the amount of cities returned.
- **Population Information**: Get the population history of a given country.
   - Optional: Get the population history for a given time interval.
- **Status Information**: Get the status of the service and external APIs.

## Requirements

- Go 1.18+
- External APIs: CountriesNow API and RestCountries API
- Go modules (use `go mod` for managing dependencies)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/SigurdRiseth/CountryInfoService.git
   cd CountryInfoService
    ```
   
2.	Install dependencies (if you haven’t already):

    ```bash
    go mod tidy
    ```

3.	Create a .env file for environment variables (if necessary). Here’s an example of what you might need to include:

    ```bash
    PORT=8080
    ```

4.	Run the service:

    ```bash
    go run main.go
    ```

5.	The service should now be running at http://localhost:8080/.

## API Endpoints

### GET /countryinfo/v1/info/

Returns the name, capital, population, and area of a country. The country name should be passed as a query parameter.

Example: http://localhost:8080/country/v1/info/no?limit=3

Response:
```json
{
  "name": "Norway",
  "continents": [
    "Europe"
  ],
  "population": 5379475,
  "languages": {
    "nno": "Norwegian Nynorsk",
    "nob": "Norwegian Bokmål",
    "smi": "Sami"
  },
  "borders": [
    "FIN",
    "SWE",
    "RUS"
  ],
  "flag": "🇳🇴",
  "capital": "Oslo",
  "cities": [
    "Abelvaer",
    "Adalsbruk",
    "Adland"
  ]
}
```

### GET /countryinfo/v1/population/

Returns the population of a country. The country name should be passed as a query parameter.

Example: http://localhost:8080/country/v1/population/no?limit=2002-2008

### GET /countryinfo/v1/status/

Returns the uptime of the service, API version, and status of the external APIs.

Example: http://localhost:8080/country/v1/status

Response:
```json
{
  "countriesnowapi": 404,
  "restcountriesapi": 400,
  "version": "1.0",
  "uptime": 13
}
```