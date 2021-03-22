package main

import (
	"github/niuniuanran/Day15/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.ServeHome)
	http.HandleFunc("/about", handlers.ServeAbout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
