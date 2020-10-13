package main

import (
	"log"
	"net/http"
	// "fmt"
)

// func main () {
// 	// fs := http.FileServer( http.Dir("/Users/cj/Documents/Go backend/Go_serve"))

// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" && r.Method == "GET" {
// 		http.ServeFile(w, r, "index.html")
// 	}
// }

func main() {

	// Create a mux for routing incoming requests
	m := http.NewServeMux()

	// All URLs will be handled by this function ?
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Create a server listening on port listed
	s := &http.Server{
		Addr:    ":9001",
		Handler: m,
	}

	// Continue to process new requests until an error occurs then log
	log.Fatal(s.ListenAndServe())
}
