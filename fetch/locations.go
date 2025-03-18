package fetch

import (
	"fmt"
	"groupie/geo"
	"strings"
	"sync"
)

var locCache = make(map[string]Location)
var locMutec sync.Mutex

type Location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

func LocationParse(relations map[string][]string) []Location {
	var locs []string
	var Locs []Location
	for relation := range relations {
		locs = append(locs, relation)
	}

	for _, loc := range locs {
		parts := strings.Split(loc, ", ")
		if len(parts) < 2 {
			// Invalid format: save without coordinates
			Locs = append(Locs, Location{Name: loc})
			continue
		}

		city, country := parts[0], parts[1]

		locMutec.Lock()
		if cached, ok := locCache[loc]; ok {
			Locs = append(Locs, cached)
			locMutec.Unlock()
			continue
		}
		locMutec.Unlock()

		city, country = apiSpecifics(city, country)
		lat, lng, err := geo.Geocode(city, country)
		if err != nil {
			fmt.Printf("Geocoding failed for %s: %v\n", loc, err)
			Locs = append(Locs, Location{Name: loc})
		} else {
			Locs = append(Locs, Location{
				Name: loc,
				Lat:  lat,
				Lng:  lng,
			})
		}

		locMutec.Lock()
		locCache[loc] = Locs[len(Locs)-1]
		locMutec.Unlock()
	}

	return Locs
}

func apiSpecifics(city, country string) (string, string) {
	switch country {
	case "uk":
		country = "gb"
	case "french polynesia":
		country = "fp"
	case "new caledonia":
		country = "fr"
	}
	switch city {
	case "north carolina":
		city = "charlotte"
	case "south carolina":
		city = "columbia"
	case "queensland":
		city = "brisbane"
	}
	return city, country
}
