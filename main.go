package main 

import (
	"log"
	"net/http"
	// "fmt"
)

func main () {
	fs := http.FileServer( http.Dir("/Users/cj/Documents/Go backend/Go_serve"))

	
	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", fs))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
// }

