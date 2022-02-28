package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Now serving on port 4040")

	http.HandleFunc("/get-shorty/", handleGetShorty)

	http.ListenAndServe(":4040", nil)

}

func handleGetShorty(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//do stuff
		fmt.Println(r.Body)
	default:
		http.Error(w, "Unable to process request", http.StatusBadRequest)
	}
}
