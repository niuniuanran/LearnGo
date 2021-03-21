package main

import (
	"flag"
	"fmt"
	"net/http"
)

const defaultPortNumber = 8080

func main() {
	portNumber := flag.Int("port", defaultPortNumber, "Port number to serve the api")
	flag.Parse()
	http.HandleFunc("/api", handleRequest)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), nil)
}
