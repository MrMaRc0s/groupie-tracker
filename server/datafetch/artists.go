package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func fetchArtists(artistURL string, ch chan<- Artist, wg *sync.WaitGroup) {
	defer wg.Done()

	// Send GET request to the API
	resp, err := http.Get(artistURL)
	if err != nil {
		fmt.Printf("failed to fetch artists: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check for non-200 status code
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

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Printf("failed to parse JSON: %v", err)
		return
	}

	for _, artist := range artists {
		ch <- artist
	}
}
