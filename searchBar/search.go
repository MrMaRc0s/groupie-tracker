package search

import (
	"groupie/fetch"
	"strconv"
	"strings"
)

func SearchArtists(query string, artists []fetch.Artist) []fetch.Artist {
	if query == "" {
		return artists
	}

	var filteredArtists []fetch.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) || strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) || strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(query)) {
			filteredArtists = append(filteredArtists, artist)
		} else {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
					filteredArtists = append(filteredArtists, artist)
					goto NextArtist
				}
			}
			for _, location := range artist.Locations {
				if strings.Contains(strings.ToLower(location.Name), strings.ToLower(query)) {
					filteredArtists = append(filteredArtists, artist)
					break
				}
			}
		}
	NextArtist:
	}
	return filteredArtists
}
