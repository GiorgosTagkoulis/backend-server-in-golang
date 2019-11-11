package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// DB the variable to have the db global
var DB *sql.DB

func get(w http.ResponseWriter, r *http.Request) {
	var (
		title string
		price float64
	)
	rows, err := DB.Query("select title, price from sitoo_test_assignment.product where product_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&title, &price)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(title, price)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
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
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	db, err := sql.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/sitoo_test_assignment")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connection Established")
	}
	defer db.Close()

	DB = db

	r := mux.NewRouter()
	api := r.PathPrefix("/api/products").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
