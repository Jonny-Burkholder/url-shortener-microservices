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

func handleGetShorty(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json;utf=8")
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			fn(w, r)
		case http.MethodPost:
			//do stuff
		default:
			log.Println("Error!")
			http.Error(w, "Unable to process request", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%s", b)
	}
}
