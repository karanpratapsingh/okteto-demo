package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		data := "Hello Okteto!"
		json.NewEncoder(w).Encode(&data)
	})

	log.Println("Server is up!")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
