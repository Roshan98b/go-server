package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

func receiveLink(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("validationToken")
	if (token != "") {
		log.Println("token validation suceeded " + token)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	} else {
		body, err := io.ReadAll(r.Body)		
		if err != nil {
			panic(err)
		}		
		log.Println(string(body))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	log.Println("Started service")

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	corsOrigins := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})

	router.HandleFunc("/api", mainLink).Methods("GET")
	router.HandleFunc("/api/upload", uploadLink).Methods("GET")
	router.HandleFunc("/api/receive", receiveLink).Methods("POST")

	log.Fatalln(http.ListenAndServe(":8080", handlers.CORS(corsOrigins)(router)))
}