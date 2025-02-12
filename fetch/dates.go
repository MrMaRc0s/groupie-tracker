package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Date struct {
	Dates []string `json:"dates"`
}

func fetchDates(dateURL string) ([]string, error) {
	resp, err := http.Get(dateURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var dateData Date
	if err := json.NewDecoder(resp.Body).Decode(&dateData); err != nil {
		return nil, fmt.Errorf("failed to parse dates JSON: %w", err)
	}

	for i := range dateData.Dates {
		dateData.Dates[i] = strings.Replace(dateData.Dates[i], "*", "", -1)
	}

	return dateData.Dates, nil
}
