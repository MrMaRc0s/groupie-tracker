package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func FetchArtists(apiURL string) ([]Artist, error) {
	// Send GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// for i, artist := range artists {
	// 	locations, dates, err := fetchLocations(artist.LocURL)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	// 	}
	// 	relation, err := FetchRelation(artist.RelURL)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to fetch relation: %v", err)
	// 	}
	// 	artists[i].Locations = locations
	// 	artists[i].ConcertDates = dates
	// 	artists[i].Relations = relation
	// }

	return artists, nil
}
