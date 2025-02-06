package main

import (
	"fmt"
	fetch "groupie/server"
	"net/http"
	"sync"
	"time"
)

var (
	dataLoaded bool
	dataMutex  sync.RWMutex
)

func main() {
	startNow := time.Now()

	apiURL := "https://groupietrackers.herokuapp.com/api"

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

	// Start data fetching in background
	go func() {
		fetch.FetchArtists(apiURL, func() {
			dataMutex.Lock()
			dataLoaded = true
			dataMutex.Unlock()
			fmt.Println("Initial data loading completed!")
		})
	}()

	fmt.Println("Time Taken to Start:", time.Since(startNow))
	fmt.Println("Server is running on http://localhost:8080")

	// Start server without blocking
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
