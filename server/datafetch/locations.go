package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Locations []string `json:"locations"`
}

func fetchLocations(locationURL string) []string {

	resp, err := http.Get(locationURL)
	if err != nil {
		fmt.Printf("failed to fetch locations: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return nil
	}

	var locations Location
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return nil
	}

	return locations.Locations
}
