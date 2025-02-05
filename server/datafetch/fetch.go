package fetch

import "sync"

func FetchAll(URL string) ([]Artist, []Location, []Date, []Relation, error) {
	var wg sync.WaitGroup

	ch := make(chan Artist, 52)
	ch1 := make(chan Location, 52)
	ch2 := make(chan Date, 52)
	ch3 := make(chan Relation, 52)
	wg.Add(4)

	go fetchArtists(URL+"/artists", ch, &wg)
	go fetchLocations(URL+"/locations", ch1, &wg)
	go fetchDates(URL+"/dates", ch2, &wg)
	go fetchRelation(URL+"/relation", ch3, &wg)

	go func() {
		wg.Wait()
		close(ch)
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	var artists []Artist
	for artist := range ch {
		artists = append(artists, artist)
	}

	var locations []Location
	for location := range ch1 {
		locations = append(locations, location)
	}

	var dates []Date
	for date := range ch2 {
		dates = append(dates, date)
	}

	var relations []Relation
	for relation := range ch3 {
		relations = append(relations, relation)
	}

	return artists, locations, dates, relations, nil
}
