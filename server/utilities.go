package server

import "groupie/fetch"

type DatesAndLocations struct {
	CreationDateMin   int
	FirstAlbumDateMin string
	AllLocations      []string
}

func getDatesAndLocations(artists []fetch.Artist) DatesAndLocations {
	if len(artists) == 0 {
		return DatesAndLocations{}
	}

	creationDateMinn := artists[0].CreationDate
	firstAlbumDateMinn := artists[0].FirstAlbum
	allLocations := []string{}
	locationSet := make(map[string]bool)
	result := DatesAndLocations{}

	for _, art := range artists {
		if art.CreationDate < creationDateMinn {
			creationDateMinn = art.CreationDate
		}
		if art.FirstAlbum < firstAlbumDateMinn {
			firstAlbumDateMinn = art.FirstAlbum
		}

		for _, loc := range art.Locations {
			if !locationSet[loc.Name] {
				allLocations = append(allLocations, loc.Name)
				locationSet[loc.Name] = true
			}
		}

	}
	if len(firstAlbumDateMinn) >= 4 {
		firstAlbumDateMinn = firstAlbumDateMinn[len(firstAlbumDateMinn)-4:]
	}
	result.CreationDateMin = creationDateMinn
	result.FirstAlbumDateMin = firstAlbumDateMinn
	result.AllLocations = allLocations
	return result
}
