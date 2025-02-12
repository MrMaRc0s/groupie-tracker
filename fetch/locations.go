package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Location struct {
	Locations []string `json:"locations"`
}

func fetchLocations(locationURL string) ([]string, error) {
	resp, err := http.Get(locationURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var locData Location
	if err := json.NewDecoder(resp.Body).Decode(&locData); err != nil {
		return nil, fmt.Errorf("failed to parse locations JSON: %w", err)
	}

	for i := range locData.Locations {
		locData.Locations[i] = strings.Replace(locData.Locations[i], "_", " ", -1)
		locData.Locations[i] = strings.Replace(locData.Locations[i], "-", ", ", -1)
	}

	return locData.Locations, nil
}
