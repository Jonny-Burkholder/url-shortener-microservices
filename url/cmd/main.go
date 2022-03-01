package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Now serving on port 4040")

	http.HandleFunc("/get-shorty", handleGetShorty)

	http.ListenAndServe(":4040", nil)

}

func setupCORSResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handleGetShorty(w http.ResponseWriter, r *http.Request) {
	setupCORSResponse(&w)
	data := "get.shorty/1"
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP Method:")
	fmt.Println(r.Method)
	fmt.Println("printing url body:")
	fmt.Println(r.Body)
	switch r.Method {
	case http.MethodOptions:
		return
	case http.MethodPost:
		fmt.Fprintf(w, "%s", b)
		return
	default:
		log.Println("Error!")
		http.Error(w, "Unable to process request", http.StatusBadRequest)
		return
	}
}
