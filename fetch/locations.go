package fetch

import (
	"encoding/json"
	"fmt"
	"groupie/geo"
	"net/http"
	"strings"
)

type Location struct {
	Locations []string `json:"locations"`
}

type ConcertLocation struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

func fetchLocations(locationURL string) ([]ConcertLocation, error) {
	resp, err := http.Get(locationURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var locData Location
	if err := json.NewDecoder(resp.Body).Decode(&locData); err != nil {
		return nil, fmt.Errorf("failed to parse locations JSON: %w", err)
	}

	for i := range locData.Locations {
		locData.Locations[i] = strings.Replace(locData.Locations[i], "_", " ", -1)
		locData.Locations[i] = strings.Replace(locData.Locations[i], "-", ", ", -1)
	}

	var concertLocs []ConcertLocation
	for _, loc := range locData.Locations {
		parts := strings.Split(loc, ", ")
		if len(parts) < 2 {
			// Invalid format: save without coordinates
			concertLocs = append(concertLocs, ConcertLocation{Name: loc})
			continue
		}

		city, country := parts[0], parts[1]
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
			fmt.Printf("Geocoding failed for %s: %v\n", loc, err)
			concertLocs = append(concertLocs, ConcertLocation{Name: loc})
		} else {
			concertLocs = append(concertLocs, ConcertLocation{
				Name: loc,
				Lat:  lat,
				Lng:  lng,
			})
		}
	}

	return concertLocs, nil
}

func LocationParse(relations map[string][]string) []ConcertLocation {
	var locs []string
	var concertLocs []ConcertLocation
	for relation := range relations {
		locs = append(locs, relation)
	}

	for _, loc := range locs {
		parts := strings.Split(loc, ", ")
		if len(parts) < 2 {
			// Invalid format: save without coordinates
			concertLocs = append(concertLocs, ConcertLocation{Name: loc})
			continue
		}

		city, country := parts[0], parts[1]
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
			fmt.Printf("Geocoding failed for %s: %v\n", loc, err)
			concertLocs = append(concertLocs, ConcertLocation{Name: loc})
		} else {
			concertLocs = append(concertLocs, ConcertLocation{
				Name: loc,
				Lat:  lat,
				Lng:  lng,
			})
		}
	}

	return concertLocs
}
