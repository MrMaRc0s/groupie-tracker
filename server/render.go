package server

import (
	"fmt"
	"groupie/fetch"
	"html/template"
	"net/http"
	"sync"
)

var (
	dataLoaded bool
	dataMutex  sync.RWMutex
	artists    []fetch.Artist
)

type PageData struct {
	Loaded  bool
	Artists []fetch.Artist
}

func Fetching(apiURL string) {
	go func() { // Run in goroutine
		fetch.FetchArtists(apiURL, func(fetchedArtists []fetch.Artist) {
			dataMutex.Lock()
			defer dataMutex.Unlock()
			artists = fetchedArtists
			dataLoaded = true
			fmt.Println("Initial data loading completed!")
		})
	}()
}

func DataRender(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	tmpl, err := template.ParseFiles("static/mainpage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageData := PageData{
		Loaded:  dataLoaded,
		Artists: artists,
	}

	// Show loading template if data isn't ready
	if !pageData.Loaded {
		w.WriteHeader(http.StatusProcessing)
		tmpl.ExecuteTemplate(w, "loading", nil)
		return
	}

	tmpl.Execute(w, pageData)
}
