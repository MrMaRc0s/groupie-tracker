package main

import (
	"fmt"
	fetch "godot/server/datafetch"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Band Info Server!")
	})
	apiURL := "https://groupietrackers.herokuapp.com/api"
	artist, err := fetch.FetchArtists(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Artists:")
	for _, a := range artist {
		fmt.Printf("ID: %d\n", a.ID)
		fmt.Printf("Name: %s\n", a.Name)
		fmt.Printf("Image: %s\n", a.Image)
		fmt.Printf("Members: %v\n", a.Members)
		fmt.Printf("Creation Date: %d\n", a.CreationDate)
		fmt.Printf("First Album: %s\n", a.FirstAlbum)
		fmt.Printf("Locations: %s\n", a.Locations)
		fmt.Printf("Concert Dates: %s\n", a.ConcertDates)
		fmt.Printf("Relations: %s\n", a.Relations)
		fmt.Println()
	}

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
