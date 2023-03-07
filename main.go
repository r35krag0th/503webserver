package main

import (
	"net/http"
)

func notReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
}

func main() {
	http.HandleFunc("/health", notReadyHandler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
