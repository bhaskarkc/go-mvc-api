package main

import (
	"encoding/json"
	"fmt"
	"github.com/bhaskarkc/go-api/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Post struct {
	ID      string `json:ID`
	Title   string `json:post_title`
	Content string `json:post_content`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	json.NewEncoder(w).Encode(posts)
}

func singleForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Requested ID: %s\n", vars["id"])
	fmt.Fprintf(w, "Single Form Page.")
}

func formSubmissions(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "Downloading Submissions for ID:%s", vars["id"])
}

func main() {
	// Setup Database.
	db := db.DbConnect()
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc("/forms/{id:[0-9]+}", singleForm).Methods("GET")
	router.HandleFunc("/forms/{id:[0-9]+}/submissions", formSubmissions).Methods("GET")

	log.Fatal(http.ListenAndServe(":7000", router))
}
