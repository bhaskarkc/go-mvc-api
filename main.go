package main

import (
	"encoding/json"
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
		`SELECT ID, post_title, post_content FROM wp_posts 
		WHERE post_type='post' and post_status='publish' 
		ORDER BY ID DESC LIMIT 6`,
	)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(posts)
}

func singlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	post := Post{}
	err := db.Db.Get(&post,
		`SELECT ID, post_title, post_content FROM wp_posts
		WHERE ID = ?`, vars["id"],
	)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(post)
}

func main() {
	// Setup Database.
	db := db.DbConnect()
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/posts", index)
	router.HandleFunc("/posts/{id:[0-9]+}", singlePost).Methods("GET")

	log.Println("Serving at localhost:7000")
	log.Fatal(http.ListenAndServe(":7000", router))
}
