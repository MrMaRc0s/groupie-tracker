package fetch

func FetchAll(URL string) ([]Artist, []Location, error) {
	// var wg sync.WaitGroup

	artists, err := FetchArtists(URL + "/artists")
	if err != nil {
		return nil, nil, err
	}
	locations, err := fetchLocations(URL + "/locations")
	if err != nil {
		return nil, nil, err
	}
	fetchDates(URL + "/dates")
	FetchRelation(URL + "/relation")

	return artists, locations, nil
}
