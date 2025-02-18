package search

import (
	"encoding/json"
	"groupie/fetch"
	"net/http"
	"strconv"
	"strings"
)

var artists []fetch.Artist

type PageData struct {
	Artists []fetch.Artist
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("query"))
	filteredArtists := []fetch.Artist{}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) || strings.Contains(strings.ToLower(artist.FirstAlbum), query) || strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			filteredArtists = append(filteredArtists, artist)
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}

		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), query) {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredArtists)
}
