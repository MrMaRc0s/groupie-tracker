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
		Locs = append(Locs, Location{Name: loc})
	}

	return Locs
}

func FetchCoords(locs []*Location) ([]Location, error) {
	for _, loc := range locs {
		parts := strings.Split(loc.Name, ", ")
		if len(parts) < 2 {
			continue
		}

		city, country := parts[0], parts[1]

		locMutec.Lock()
		if cached, ok := locCache[loc.Name]; ok {
			fmt.Printf("Using cached data for %s\n", loc.Name)
			loc.Lat, loc.Lng = cached.Lat, cached.Lng
			locMutec.Unlock()
			continue
		}
		locMutec.Unlock()

		city, country = apiSpecifics(city, country)
		lat, lng, err := geo.Geocode(city, country)
		if err != nil {
			fmt.Printf("Geocoding failed for %s: %v\n", loc.Name, err)
			continue
		} else {
			loc.Lat, loc.Lng = lat, lng
		}

		fmt.Println("Geocoding completed for", loc.Name)
		locMutec.Lock()
		locCache[loc.Name] = Location{Name: loc.Name, Lat: lat, Lng: lng}
		fmt.Printf("Cached data for %s\n", loc.Name)
		locMutec.Unlock()
	}

	updatedLocs := make([]Location, len(locs))
	for i, loc := range locs {
		updatedLocs[i] = *loc
	}

	return updatedLocs, nil
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
