package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Item is the return value
type Item struct {
	ProductID string `json:"productId"`
	Title     string `json:"title"`
}

// Post is for the posts
type Post struct {
	TotalCount int    `json:"totalCount"`
	Items      []Item `json:"items"`
}

func get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post Post
	var items []Item

	countRes, err := db.Query("SELECT COUNT(product_id) FROM sitoo_test_assignment.product")
	if err != nil {
		log.Fatal(err)
	}
	defer countRes.Close()
	for countRes.Next() {
		if err := countRes.Scan(&post.TotalCount); err != nil {
			log.Fatal(err)
		}
	}

	itemRes, err := db.Query("SELECT product_id, title FROM sitoo_test_assignment.product")
	if err != nil {
		log.Fatal(err)
	}
	defer itemRes.Close()

	for itemRes.Next() {
		var item Item
		err := itemRes.Scan(&item.ProductID, &item.Title)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(post)
		items = append(items, item)
	}

	post.Items = items
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
	// w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"errorText": "The error message"}`))
}

func main() {
	db, err = sql.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/sitoo_test_assignment")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connection Established")
	}
	defer db.Close()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/products").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
