package search

import (
	"fmt"
	"groupie/fetch"
	"strings"
)

type Suggestion struct {
	Display  string // What to show in the dropdown
	Type     string // Entity type (member, band, location, etc.)
	ArtistID int    // Reference to the full artist
}

func GetSuggestions(query string, artists []fetch.Artist) []Suggestion {
	query = strings.ToLower(query)
	var suggestions []Suggestion
	seen := make(map[string]bool) // To prevent duplicates

	for _, artist := range artists {
		// Check name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			key := fmt.Sprintf("name:%s", artist.Name)
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Display:  artist.Name,
					Type:     "artist",
					ArtistID: artist.ID,
				})
				seen[key] = true
			}
		}

		// Check members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				key := fmt.Sprintf("member:%s", member)
				if !seen[key] {
					suggestions = append(suggestions, Suggestion{
						Display:  member,
						Type:     "member",
						ArtistID: artist.ID,
					})
					seen[key] = true
				}
			}
		}

		// Check other fields similarly for locations, albums, etc.
	}
	return suggestions[:10] // Return top 10 suggestions
}
