package server

import (
	"groupie/fetch"
	search "groupie/searchBar"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var (
	dataMutex sync.RWMutex
	// TopTracks  []fetch.Track
)

type PageData struct {
	Artists []fetch.Artist
	Query   string
	// Tracks  []fetch.Track
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

	query := r.URL.Query().Get("query")
	filteredArtists := search.SearchArtists(query, artists)

	tmpl, err := template.ParseFiles("static/homepage.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	pageData := PageData{
		Artists: filteredArtists,
		Query:   query,
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
		Artists: []fetch.Artist{artists[id]},
		// Tracks:  TopTracks,
	}
	tmpl.Execute(w, pageData)
}
