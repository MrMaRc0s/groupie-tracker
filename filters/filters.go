package filters

import (
	"groupie/fetch"
	"net/url"
	"strconv"
	"strings"
)

func FilterArtists(artists []fetch.Artist, queryParams url.Values) []fetch.Artist {
	var filteredArtists []fetch.Artist

	// Parse query parameters safely
	creationDateMin, errMin := strconv.Atoi(queryParams.Get("creationDateMin"))
	creationDateMax, errMax := strconv.Atoi(queryParams.Get("creationDateMax"))
	if errMin != nil {
		creationDateMin = 0 // Default value
	}
	if errMax != nil {
		creationDateMax = 9999 // Default max year
	}

	firstAlbumDateMin, errMinAlbum := strconv.Atoi(queryParams.Get("firstAlbumDateMin"))
	firstAlbumDateMax, errMaxAlbum := strconv.Atoi(queryParams.Get("firstAlbumDateMax"))
	if errMinAlbum != nil {
		firstAlbumDateMin = 0
	}
	if errMaxAlbum != nil {
		firstAlbumDateMax = 9999
	}

	numberMembers := queryParams["Member"] // Match HTML checkboxes
	var albumYear int

	for _, artist := range artists {
		// Correct date filtering
		if artist.CreationDate < creationDateMin || artist.CreationDate > creationDateMax {
			continue
		}

		album := strings.Split(artist.FirstAlbum, "-")
		if len(album) > 0 {
			albumYear, _ = strconv.Atoi(album[len(album)-1])
		}

		if albumYear < firstAlbumDateMin || albumYear > firstAlbumDateMax {
			continue
		}

		// Filter by number of members
		if len(numberMembers) > 0 {
			membersMatch := false
			for _, member := range numberMembers {
				memberCount, err := strconv.Atoi(member)
				if err == nil && len(artist.Members) == memberCount {
					membersMatch = true
					break
				}
			}
			if !membersMatch {
				continue
			}
		}

		filteredArtists = append(filteredArtists, artist)
	}

	return filteredArtists
}
