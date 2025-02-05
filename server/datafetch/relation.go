package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func fetchRelation(relationURL string) map[string][]string {

	resp, err := http.Get(relationURL)
	if err != nil {
		fmt.Printf("failed to fetch relation: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return nil
	}

	var relation Relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return nil
	}

	return relation.DatesLocations
}
