package main

import (
	"fmt"
	fetch "godot/server/datafetch"
	"net/http"
	"time"
)

func main() {
	startNow := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Band Info Server!")
	})
	apiURL := "https://groupietrackers.herokuapp.com/api"
	artists, locations, err := fetch.FetchAll(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Artists:")
	for _, artist := range artists {
		fmt.Printf("ID: %d\n", artist.ID)
		fmt.Printf("Name: %s\n", artist.Name)
		fmt.Printf("Image: %s\n", artist.Image)
		fmt.Printf("Members: %v\n", artist.Members)
		fmt.Printf("Creation Date: %d\n", artist.CreationDate)
		fmt.Printf("First Album: %s\n", artist.FirstAlbum)
		// fmt.Printf("Locations: %s\n", artist.Locations)
		// fmt.Printf("Concert Dates: %s\n", artist.ConcertDates)
		// fmt.Printf("Relations: %s\n", artist.Relations)
		fmt.Println()
	}

	fmt.Println("Locations:")

	fmt.Printf("Location: %s\n", locations)

	fmt.Println("Time Taken: ", time.Since(startNow))

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		fmt.Println("Server is running on http://localhost:8081")

		if err := http.ListenAndServe(":8081", nil); err != nil {
			fmt.Println("Port 8081 is already in use", err)
			return
		}
	}
}
