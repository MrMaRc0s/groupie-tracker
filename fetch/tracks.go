// DEN TREXEI POUTHENA !!!!!!!!

package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Track struct {
	// Name string `json:"name"`
	URL string `json:"url"`
}

type TopTracksResponse struct {
	TopTracks struct {
		Track []Track `json:"track"`
	} `json:"toptracks"`
}

func FetchTracks(name string) Track {
	apiKey := "777c4a8284a8eb36cc81d659242ae210"
	baseURL := "https://ws.audioscrobbler.com/2.0/"
	params := url.Values{}
	params.Set("method", "artist.gettoptracks")
	params.Set("artist", name)
	params.Set("api_key", apiKey)
	params.Set("format", "json")
	params.Set("limit", "1")

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(requestURL)
	if err != nil {
		return Track{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Track{}
	}

	var topTracksResponse TopTracksResponse
	if err := json.NewDecoder(resp.Body).Decode(&topTracksResponse); err != nil {
		return Track{}
	}

	fmt.Println("Top tracks for", name, ":", topTracksResponse.TopTracks.Track)

	return topTracksResponse.TopTracks.Track[0]
}
