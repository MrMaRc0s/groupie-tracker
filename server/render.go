package server

import (
	"fmt"
	"groupie/fetch"
	"html/template"
	"net/http"
	"sync"
)

var (
	// dataLoaded bool
	dataMutex sync.RWMutex
)

var artists []fetch.Artist

type PageData struct {
	Artists []fetch.Artist
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	dataMutex.RLock()
// 	defer dataMutex.RUnlock()

// 	if !dataLoaded {
// 		w.WriteHeader(http.StatusProcessing)
// 		fmt.Fprintln(w, "Server is starting up... Please refresh in a moment!")
// 		return
// 	}
// 	http.ServeFile(w, r, "static/mainpage.html")
// }

func Fetching(apiURL string) {
	fetch.FetchArtists(apiURL, func(fetchedArtists []fetch.Artist) { // Pass data
		dataMutex.Lock()
		artists = fetchedArtists // Store fetched artists
		// dataLoaded = true
		dataMutex.Unlock()
		fmt.Println("Initial data loading completed!")
	})
}

func DataRender(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()
	pageData := PageData{
		Artists: artists,
	}
	tmpl, err := template.ParseFiles("static/mainpage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, pageData)
}
