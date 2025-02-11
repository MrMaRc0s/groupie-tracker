package main

import (
	"fmt"
	"groupie/server"
	"net/http"
	"time"
)

func main() {
	startNow := time.Now()
	apiURL := "https://groupietrackers.herokuapp.com/api"

	// Setup routes first
	http.HandleFunc("/", server.HomeRender)
	http.HandleFunc("/homepage", server.DataRender)
	http.HandleFunc("/details", server.DetailsRender)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	// Start data fetching in background
	server.Fetching(apiURL) // Already wrapped in goroutine

	fmt.Println("Time Taken to Start:", time.Since(startNow))
	fmt.Println("Server is running on http://localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
