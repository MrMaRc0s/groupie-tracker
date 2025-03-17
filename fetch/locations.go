package fetch

import (
	"groupie/geo"
	"strings"
)

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
		// city = strings.Replace(city, "_", " ", -1)
		// country = strings.Replace(country, "_", " ", -1)
		// fmt.Println(city, country)
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
		lat, lng, err := geo.Geocode(city, country)
		if err != nil {
			// fmt.Printf("Geocoding failed for %s: %v\n", loc, err)
			Locs = append(Locs, Location{Name: loc})
		} else {
			Locs = append(Locs, Location{
				Name: loc,
				Lat:  lat,
				Lng:  lng,
			})
		}
	}

	return Locs
}
