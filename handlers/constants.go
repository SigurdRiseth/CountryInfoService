package handler

import (
	"fmt"
	"net/http"
)

// Define the paths for the different endpoints
const (
	INFO_PATH              = "/country/v1/info"
	POPULATION_PATH        = "/country/v1/population"
	STATUS_PATH            = "/country/v1/status"
	API_VERSION            = "1.0"
	COUNTRIES_NOW_API_URL  = "http://129.241.150.113:3500/api/v0.1/"
	REST_COUNTRIES_API_URL = "http://129.241.150.113:8080/v3.1/alpha/"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Country Info API</title>
</head>
<body>
    <h1>Welcome to the Country Info API</h1>
    <p>This API provides information about different countries.</p>

    <h2>Endpoints</h2>

    <ul>
        <li>
            <code>/country/v1/info</code> - List country information
            <ul>
                <li><strong>Method:</strong> GET</li>
                <li><strong>Query Parameters:</strong></li>
                <ul>
                    <li><code>two_letter_country_code</code> (string, required) - The 2-letter-ISO of the country (e.g., <code>no</code>)</li>
                    <li><code>limit</code> (int, optional) - The number of cities listen in the response (e.g., <code>?limit=10</code>)</li>
                </ul>
                <li><strong>Example Request:</strong></li>
                <pre><code>GET /country/v1/info/no?limit=10</code></pre>
            </ul>
        </li>

        <li>
            <code>/country/v1/population</code> - Get details about a country's population over time
            <ul>
                <li><strong>Method:</strong> GET</li>
                <li><strong>Query Parameters:</strong></li>
                <ul>
                    <li><code>two_letter_country_code</code> (string, required) - The name of the country (e.g., <code>jp</code>)</li>
                    <li><code>?limit={startYear-endYear}</code> (string, optional) - Specific year interval for population data (e.g., <code>?limit=2020-2024</code>)</li>
                </ul>
                <li><strong>Example Request:</strong></li>
                <pre><code>GET /country/v1/population/jp?limit=2020-2024</code></pre>
            </ul>
        </li>

        <li>
            <code>/country/v1/status</code> - Get API status
            <ul>
                <li><strong>Method:</strong> GET</li>
                <li><strong>Response:</strong> API status details (e.g., uptime, version)</li>
                <li><strong>Example Request:</strong></li>
                <pre><code>GET /country/v1/status</code></pre>
            </ul>
        </li>
    </ul>
</body>
</html>
		`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
