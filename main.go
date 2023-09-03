package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequest() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
    log.Print("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
