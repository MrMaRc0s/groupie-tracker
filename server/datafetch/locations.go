package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationData struct {
	ID        int    `json:"id"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
}

func FetchLocations(locationURL string) (LocationData, error) {
	// Send GET request to fetch location data
	resp, err := http.Get(locationURL)
	if err != nil {
		return LocationData{}, fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return LocationData{}, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON into the struct
	var locationData LocationData
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return LocationData{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return locationData, nil
}
