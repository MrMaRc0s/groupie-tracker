package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DateData struct {
	Dates []string `json:"dates"`
}

func fetchDates(dateURL string) ([]string, error) {

	resp, err := http.Get(dateURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dates: %v", err)
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

	var dates DateData
	err = json.Unmarshal(body, &dates)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return dates.Dates, nil
}
