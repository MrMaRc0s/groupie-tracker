package main

import (
	"fmt"
	fetch "groupie/server"
	"net/http"
	"time"
)

func main() {
	startNow := time.Now()

	apiURL := "https://groupietrackers.herokuapp.com/api"

	// Serve static files (like HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/mainpage.html")
	})

	// Fetch data
	fetch.FetchArtists(apiURL)

	fmt.Println("Time Taken: ", time.Since(startNow))
	fmt.Println("Server is running on http://localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		fmt.Println("Trying port 8081...")

		if err := http.ListenAndServe(":8081", nil); err != nil {
			fmt.Println("Port 8081 is already in use", err)
			return
		}
	}
}
