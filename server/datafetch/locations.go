package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Locations []string `json:"locations"`
}

func fetchLocations(locationURL string) ([]Location, error) { // Send GET request to fetch location data

	resp, err := http.Get(locationURL)
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

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// locations.Dates, err = fetchDates(locations.DatesURL)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to fetch dates: %v", err)
	// }

	return locations, nil
}
