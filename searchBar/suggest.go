package search

import (
	"encoding/json"
	"groupie/fetch"
	"net/http"
	"strconv"
	"strings"
)

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("query"))

	type Suggestion struct {
		Value string `json:"value"`
		Type  string `json:"type"`
		ID    int    `json:"id"`
	}

	var suggestions []Suggestion

	artists := []fetch.Artist{}

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, Suggestion{Value: artist.Name, Type: "artist", ID: artist.ID})
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, Suggestion{Value: member, Type: "member", ID: artist.ID})
			}
		}

		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), query) {
				suggestions = append(suggestions, Suggestion{Value: location, Type: "location", ID: artist.ID})
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			suggestions = append(suggestions, Suggestion{Value: artist.FirstAlbum, Type: "firstAlbum", ID: artist.ID})
		}

		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			suggestions = append(suggestions, Suggestion{Value: strconv.Itoa(artist.CreationDate), Type: "creationDate", ID: artist.ID})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
