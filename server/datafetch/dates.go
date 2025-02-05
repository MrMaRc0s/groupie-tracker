package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Date struct {
	Dates []string `json:"dates"`
}

func fetchDates(dateURL string, ch chan<- Date, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(dateURL)
	if err != nil {
		fmt.Printf("failed to fetch dates: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return
	}

	var dates map[string][]Date
	err = json.Unmarshal(body, &dates)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return
	}

	for _, date := range dates["index"] {
		ch <- date
	}
}
