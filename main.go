package main

import (
	"fmt"
	fetch "godot/server/datafetch"
	"net/http"
	"time"
)

func main() {
	startNow := time.Now()

	apiURL := "https://groupietrackers.herokuapp.com/api"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Band Info Server!")
	})

	fetch.FetchArtists(apiURL)

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
