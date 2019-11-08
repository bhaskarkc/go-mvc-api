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
	Id      uint64 `db:"ID" json:"ID"`
	Title   string `db:"post_title" json:"Title"`
	Content string `db:"post_content" json:"Content"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts := []Post{}

	err := db.Db.Select(&posts,
		"Select ID, post_title, post_content from wp_posts WHERE post_type='post' and post_status='publish' order by ID DESC LIMIT 6",
	)

	if err != nil {
		panic(err)
	}

	// log.Println(posts)
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
