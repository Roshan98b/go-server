package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func openFile(filename string) ESG {
	file, err := os.ReadFile(filename)
	if (err != nil) {
		log.Fatalln(err)
	}

	esg := ESG {}
	err = json.Unmarshal(file, &esg)
	if (err != nil) {
		log.Fatalln(err)
	}

	return esg
}

func mainLink(w http.ResponseWriter, r *http.Request) {
	log.Println("/api is called")
	fmt.Fprintf(w, "Go API")
}

func uploadLink(w http.ResponseWriter, r *http.Request) {
	filename := "response.json"
	log.Println("/api is called")

	w.Header().Set("Content-Type", "application/json")
	esg := openFile(filename)

	log.Println("sending back response : ")
	json.NewEncoder(w).Encode(esg)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	log.Println("Started service")
	router.HandleFunc("/api", mainLink).Methods("GET")
	router.HandleFunc("/api/upload", uploadLink).Methods("GET")

	log.Fatalln(http.ListenAndServe(":8080", router))
}