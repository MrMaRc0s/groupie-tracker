package server

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	dataLoaded bool
	dataMutex  sync.RWMutex
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	if !dataLoaded {
		w.WriteHeader(http.StatusProcessing)
		fmt.Fprintln(w, "Server is starting up... Please refresh in a moment!")
		return
	}
	http.ServeFile(w, r, "static/mainpage.html")
}
