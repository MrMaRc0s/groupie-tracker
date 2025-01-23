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
	locationData, err := fetch.FetchLocations(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d\n", locationData.ID)
	fmt.Printf("Locations: %v\n", locationData.Locations)
	fmt.Printf("Dates URL: %s\n", locationData.Dates)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
