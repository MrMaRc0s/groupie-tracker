package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func fetchRelation(relationURL string, ch chan<- Relation, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(relationURL)
	if err != nil {
		fmt.Printf("failed to fetch relation: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return
	}

	var relation map[string][]Relation
	err = json.Unmarshal(body, &relation)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return
	}

	for _, relation := range relation["index"] {
		ch <- relation
	}
}
