package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Model struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Models []Model

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/models", GetModels)
	log.Fatal(http.ListenAndServe(":5200", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	// enabling CORS to work with Angular on the same machine
	w.Header().Set("Access-Control-Allow-Origin", "*")

	models := Models{
		Model{ID: 1, Name: "Fibonacci"},
		Model{ID: 2, Name: "Couchy"},
		Model{ID: 3, Name: "Laplace"},
		Model{ID: 5, Name: "Fourier"},
		Model{ID: 8, Name: "Lagrange"},
		Model{ID: 13, Name: "Talete"},
		Model{ID: 21, Name: "Leibnitz"},
		Model{ID: 34, Name: "Dini"},
		Model{ID: 55, Name: "Rouche"},
		Model{ID: 89, Name: "Capelli"},
	}

	json.NewEncoder(w).Encode(models)

}
