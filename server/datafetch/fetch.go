package fetch

import "sync"

func FetchAll(URL string) ([]Artist, error) {
	var wg sync.WaitGroup

	ch := make(chan Artist)

	go fetchArtists(URL, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var artists []Artist
	for artist := range ch {
		artists = append(artists, artist)
	}

	return artists, nil
}
