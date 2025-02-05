package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Location struct {
	Locations []string `json:"locations"`
}

func fetchLocations(locationURL string, ch chan<- Location, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(locationURL)
	if err != nil {
		fmt.Printf("failed to fetch locations: %v", err)
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

	var locations map[string][]Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return
	}

	for _, location := range locations["index"] {
		ch <- location
	}
}
