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

func FetchArtists(URL string) {

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

	var wg sync.WaitGroup
	chLoc := make(chan Location, len(artists))
	chDat := make(chan Date, len(artists))
	chRel := make(chan Relation, len(artists))
	updatedArtists := make([]Artist, len(artists))
	// ch := make(chan Artist, len(artists))

	for i, artist := range artists {
		wg.Add(1)
		go func(artist Artist, i int, chLoc chan<- Location, chDat chan<- Date, chRel chan<- Relation, wg *sync.WaitGroup) {
			defer wg.Done()

			wg.Add(3)
			// fmt.Println(URL + "/locations/" + strconv.Itoa(i+1))
			go fetchLocations(URL + "/locations/" + strconv.Itoa(i+1), chLoc, wg)
			// fmt.Println(URL + "/dates/" + strconv.Itoa(i+1))
			go fetchDates(URL + "/dates/" + strconv.Itoa(i+1), chDat, wg)
			// fmt.Println(URL + "/relation/" + strconv.Itoa(i+1))
			go fetchRelation(URL + "/relation/" + strconv.Itoa(i+1), chRel, wg)
			updatedArtists[i] = artist
			
			// ch <- artist
		}(artist, i, chLoc, chDat, chRel, &wg)
	}

	wg.Wait()
	close(chLoc)
	close(chDat)
	close(chRel)

	// for artist := range ch {
	// 	updatedArtists = append(updatedArtists, artist)
	// }

	fmt.Println("Artists:")
	for _, artist := range updatedArtists {
		fmt.Printf("ID: %d\n", artist.ID)
		fmt.Printf("Name: %s\n", artist.Name)
		fmt.Printf("Image: %s\n", artist.Image)
		fmt.Printf("Members: %v\n", artist.Members)
		fmt.Printf("Creation Date: %d\n", artist.CreationDate)
		fmt.Printf("First Album: %s\n", artist.FirstAlbum)
		fmt.Printf("Locations: %s\n", artist.Locations)
		fmt.Printf("Concert Dates: %s\n", artist.ConcertDates)
		fmt.Printf("Relations: %s\n", artist.Relations)
		fmt.Println()
	}
}
