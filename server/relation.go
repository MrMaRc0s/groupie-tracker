package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func fetchRelation(relationURL string) (map[string][]string, error) {
	resp, err := http.Get(relationURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var relData Relation
	if err := json.NewDecoder(resp.Body).Decode(&relData); err != nil {
		return nil, fmt.Errorf("failed to parse relation JSON: %w", err)
	}

	return relData.DatesLocations, nil
}
