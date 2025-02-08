package server

import (
	"fmt"
	"groupie/fetch"
	"html/template"
	"net/http"
)

var artists []fetch.Artist

func Fetching(apiURL string) {
	fetch.FetchArtists(apiURL, func(fetchedArtists []fetch.Artist) { // Pass data
		dataMutex.Lock()
		artists = fetchedArtists // Store fetched artists
		dataLoaded = true
		dataMutex.Unlock()
		fmt.Println("Initial data loading completed!")
	})
}

func DataRender(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	tmpl, err := template.ParseFiles("static/mainpage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, artist := range artists {
		tmpl.Execute(w, artist)
	}
}
