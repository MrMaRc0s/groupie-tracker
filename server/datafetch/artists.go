package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"-"`
	ConcertDates []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
}

func fetchArtists(URL string, ch chan<- Artist, wg *sync.WaitGroup) {
	defer wg.Done()

	artistURL := URL + "/artists"
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

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return
	}

	for i, artist := range artists {
		wg.Add(1)
		// fmt.Println(URL + "/locations/" + strconv.Itoa(i+1))
		artist.Locations = fetchLocations(URL + "/locations/" + strconv.Itoa(i+1))
		// fmt.Println(URL + "/dates/" + strconv.Itoa(i+1))
		artist.ConcertDates = fetchDates(URL + "/dates/" + strconv.Itoa(i+1))
		// fmt.Println(URL + "/relation/" + strconv.Itoa(i+1))
		artist.Relations = fetchRelation(URL + "/relation/" + strconv.Itoa(i+1))
		ch <- artist
	}

}
