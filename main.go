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

	http.HandleFunc("/", server.Handler)
	http.HandleFunc("/artists", server.DataRender)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	server.Fetching(apiURL)

	fmt.Println("Time Taken to Start:", time.Since(startNow))
	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
