package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func fetchLocations(locationURL string, artistID int) ([]string, error) { // Send GET request to fetch location data

	resp, err := http.Get(locationURL + "/locations")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var locationData struct {
		Index []struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
			Dates     string   `json:"dates"`
		} `json:"index"`
	}

	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, data := range locationData.Index {
		if data.ID == artistID {
			return data.Locations, nil
		}
	}

	return nil, fmt.Errorf("no locations found for artist ID %d", artistID)
}
