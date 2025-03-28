package filters

import (
	"groupie/fetch"
	"net/url"
	"strconv"
	"strings"
)

func FilterArtists(artists []fetch.Artist, queryParams url.Values) []fetch.Artist {
	var filteredArtists []fetch.Artist

	// Parse query parameters
	creationDateMin, _ := strconv.Atoi(queryParams.Get("creationDateMin"))
	creationDateMax, _ := strconv.Atoi(queryParams.Get("creationDateMax"))

	firstAlbumDateMin, _ := strconv.Atoi(queryParams.Get("firstAlbumDateMin"))
	firstAlbumDateMax, _ := strconv.Atoi(queryParams.Get("firstAlbumDateMax"))
	numberMembers := queryParams["number-members"]
	var album []string
	var albumYear int

	for _, artist := range artists {
		// Filter by minimum creation date
		if artist.CreationDate < creationDateMin && artist.CreationDate > creationDateMax {
			continue
		}

		album = strings.Split(artist.FirstAlbum, "-")
		albumYear, _ = strconv.Atoi(album[len(album)-1])
		// Filter by minimum first album date
		if albumYear < firstAlbumDateMin && albumYear > firstAlbumDateMax {
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

		filteredArtists = append(filteredArtists, artist)
	}

	return filteredArtists
}
