package filters

import (
	"net/url"
	"strconv"
)

// Artist represents the structure of an artist.
type Artist struct {
	ID           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   int
}

// FilterArtists filters the list of artists based on the query parameters.
func FilterArtists(artists []Artist, queryParams url.Values) []Artist {
	var filteredArtists []Artist

	// Parse query parameters
	minCreationDate, _ := strconv.Atoi(queryParams.Get("min-date"))
	minAlbumDate, _ := strconv.Atoi(queryParams.Get("min-album-date-range"))
	numberMembers := queryParams["number-members"]

	for _, artist := range artists {
		// Filter by minimum creation date
		if artist.CreationDate < minCreationDate {
			continue
		}

		// Filter by minimum first album date
		if artist.FirstAlbum < minAlbumDate {
			continue
		}

		// Filter by number of members
		if len(numberMembers) > 0 {
			membersMatch := false
			for _, member := range numberMembers {
				memberCount, _ := strconv.Atoi(member)
				if len(artist.Members) == memberCount {
					membersMatch = true
					break
				}
			}
			if !membersMatch {
				continue
			}
		}

		// Add artist to the filtered list
		filteredArtists = append(filteredArtists, artist)
	}

	return filteredArtists
}
