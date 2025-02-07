package main

import (
	"encoding/json"
	"fmt"
	fetch "groupie/server"
	"net/http"
	"sync"
	"time"
)

var (
	dataLoaded bool
	dataMutex  sync.RWMutex
	artists    []fetch.Artist // Store fetched artist data
)

func main() {
	startNow := time.Now()

	apiURL := "https://groupietrackers.herokuapp.com/api"

	// Serve the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dataMutex.RLock()
		defer dataMutex.RUnlock()

		if !dataLoaded {
			w.WriteHeader(http.StatusProcessing)
			fmt.Fprintln(w, "Server is starting up... Please refresh in a moment!")
			return
		}
		http.ServeFile(w, r, "static/mainpage.html")
	})

	// ✅ Add API Route for artists
	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		dataMutex.RLock()
		defer dataMutex.RUnlock()

		if !dataLoaded {
			w.WriteHeader(http.StatusProcessing)
			fmt.Fprintln(w, "Data is still loading. Try again shortly.")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists) // Return artist data
	})

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	// Start data fetching in the background
	go func() {
		fetch.FetchArtists(apiURL, func(fetchedArtists []fetch.Artist) { // Pass data
			dataMutex.Lock()
			artists = fetchedArtists // Store fetched artists
			dataLoaded = true
			dataMutex.Unlock()
			fmt.Println("Initial data loading completed!")
		})
	}()

	fmt.Println("Time Taken to Start:", time.Since(startNow))
	fmt.Println("Server is running on http://localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
