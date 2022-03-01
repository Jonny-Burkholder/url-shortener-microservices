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

func handleGetShorty(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome!")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json;utf=8")
	data := "get.shorty/1"
	b, err := json.Marshal(map[string]interface{}{
		"api": map[string]interface{}{
			"name":    "get shorty",
			"version": 0.1,
			"data":    data,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodPost:
		//do stuff
		fmt.Println(r.Body)
	default:
		http.Error(w, "Unable to process request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "%s", b)
}
