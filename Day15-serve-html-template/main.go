package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/about", serveAbout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
