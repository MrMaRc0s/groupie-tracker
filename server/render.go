package server

import (
	"encoding/json"
	"fmt"
	"groupie/fetch"
	"groupie/filters"
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
	Artists     []fetch.Artist
	ArtistsJSON template.JS
	Query       string
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
	searchedArtists := search.SearchArtists(query, artists)
	filteredArtists := filters.FilterArtists(searchedArtists, r.URL.Query())

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		tmpl, err := template.ParseFiles("static/searched.html")
		if err != nil {
			ErrorPage(w, http.StatusInternalServerError)
			return
		}

		pageData := PageData{
			Artists: filteredArtists,
			Query:   query,
			// Tracks:  TopTracks,
		}

		tmpl.ExecuteTemplate(w, "searched", pageData)
		return
	}

	tmpl, err := template.ParseFiles("static/homepage.html", "static/searched.html", "static/header.html")
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
	tmpl, err := template.ParseFiles("static/info.html", "static/header.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	id--
	if id < 0 || id >= len(artists) {
		ErrorPage(w, http.StatusBadRequest)
		return
	}

	var locPtrs []*fetch.Location
	for i := range artists[id].Locations {
		if artists[id].Locations[i].Lat == 0 && artists[id].Locations[i].Lng == 0 {
			locPtrs = append(locPtrs, artists[id].Locations[i])
		} else {
			// fmt.Println("Using cached data from file for", artists[id].Locations[i].Name)
		}
	}

	if len(locPtrs) > 0 {
		fmt.Println()
		fetch.FetchCoords(locPtrs)
		saveData("data.json", artists)
		fmt.Println("Cache coordinates for locations:")
		for i := range locPtrs {
			fmt.Println(locPtrs[i].Name, locPtrs[i].Lat, locPtrs[i].Lng)
		}
		fmt.Println()
	}

	pageData := PageData{
		Artists: []fetch.Artist{artists[id]},
		// Tracks:  TopTracks,
	}

	//JSON marshaling
	artistsJSON, err := json.Marshal(pageData.Artists)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	pageData.ArtistsJSON = template.JS(artistsJSON)
	// fmt.Println("Sending JSON:", template.JS(artistsJSON))
	tmpl.Execute(w, pageData)
}
