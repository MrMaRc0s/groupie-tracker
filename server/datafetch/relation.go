package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RelationData struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchRelation(relationURL string) (map[string][]string, error) {

	resp, err := http.Get(relationURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relation: %v", err)
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

	var relation RelationData
	err = json.Unmarshal(body, &relation)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return relation.DatesLocations, nil
}
