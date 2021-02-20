package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                         // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:5000", nil)) // http.ListenAndServer will only return a value (an error) when something goes wrong. When it returns. log will print the error and then call os.Exit(1)
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Print(r.URL)
}
