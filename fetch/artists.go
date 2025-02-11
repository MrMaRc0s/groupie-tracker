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

func FetchArtists(URL string, completion func([]Artist)) { // Pass artists data
	artistURL := URL + "/artists"
	resp, err := http.Get(artistURL)
	if err != nil {
		fmt.Printf("Failed to fetch artists: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code %d", resp.StatusCode)
		return
	}

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Printf("Failed to parse artists JSON: %v", err)
		return
	}

	var wg sync.WaitGroup
	updatedArtists := make([]Artist, len(artists))

	for i, artist := range artists {
		wg.Add(1)
		go func(a Artist, index int) {
			defer wg.Done()

			// Fetch extra details
			locCh := make(chan []string, 1)
			dateCh := make(chan []string, 1)
			relCh := make(chan map[string][]string, 1)

			locationURL := URL + "/locations/" + strconv.Itoa(a.ID)
			datesURL := URL + "/dates/" + strconv.Itoa(a.ID)
			relationURL := URL + "/relation/" + strconv.Itoa(a.ID)

			go func() {
				locations, err := fetchLocations(locationURL)
				if err != nil {
					fmt.Printf("Error fetching locations for artist %d: %v\n", a.ID, err)
					locCh <- nil
					return
				}
				locCh <- locations
			}()

			go func() {
				dates, err := fetchDates(datesURL)
				if err != nil {
					fmt.Printf("Error fetching dates for artist %d: %v\n", a.ID, err)
					dateCh <- nil
					return
				}
				dateCh <- dates
			}()

			go func() {
				relations, err := fetchRelation(relationURL)
				if err != nil {
					fmt.Printf("Error fetching relations for artist %d: %v\n", a.ID, err)
					relCh <- nil
					return
				}
				relCh <- relations
			}()

			a.Locations = <-locCh
			a.ConcertDates = <-dateCh
			a.Relations = <-relCh

			updatedArtists[index] = a
		}(artist, i)
	}

	wg.Wait()

	fmt.Println("Artists fetched successfully!")

	// Pass updated artist data to completion function
	if completion != nil {
		completion(updatedArtists)
	}
}
