package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Date struct {
	Dates []string `json:"dates"`
}

func fetchDates(dateURL string) []string {

	resp, err := http.Get(dateURL)
	if err != nil {
		fmt.Printf("failed to fetch dates: %v", err)
		return nil
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return nil
	}

	var dates Date
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return nil
	}

	return dates.Dates
}
