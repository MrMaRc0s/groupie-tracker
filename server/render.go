package server

import (
	"fmt"
	"groupie/fetch"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var (
	dataLoaded bool
	dataMutex  sync.RWMutex
	artists    []fetch.Artist
	// TopTracks  []fetch.Track
)

type PageData struct {
	Loaded  bool
	Artists []fetch.Artist
	// Tracks  []fetch.Track
}

func Fetching(apiURL string) {
	go func() { // Run in goroutine
		fetch.FetchArtists(apiURL, func(fetchedArtists []fetch.Artist) {
			dataMutex.Lock()
			defer dataMutex.Unlock()
			artists = fetchedArtists
			// for _, artist := range artists {
			// 	TopTracks = append(TopTracks, fetch.FetchTracks(artist.Name))
			// }
			dataLoaded = true
			fmt.Println("Initial data loading completed!")
		})
	}()
}

func HomeRender(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("static/welcomepage.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DataRender(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	if r.URL.Path != "/homepage" {
		ErrorPage(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("static/homepage.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	pageData := PageData{
		Loaded:  dataLoaded,
		Artists: artists,
		// Tracks:  TopTracks,
	}

	tmpl.Execute(w, pageData)
}

func DetailsRender(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/details" {
		ErrorPage(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("static/info.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	id -= 1
	if id < 0 || id >= len(artists) {
		ErrorPage(w, http.StatusBadRequest)
		return
	}
	pageData := PageData{
		Loaded:  dataLoaded,
		Artists: []fetch.Artist{artists[id]},
		// Tracks:  TopTracks,
	}
	tmpl.Execute(w, pageData)
}
