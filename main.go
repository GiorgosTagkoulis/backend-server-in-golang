package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

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

// GetAllProducts is for all products
type GetAllProducts struct {
	TotalCount int    `json:"totalCount"`
	Items      []Item `json:"items"`
}

// Attribute is for a specific attribute
type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetProduct is for one product only
type GetProduct struct {
	ProductID   int         `json:"productId"`
	Title       string      `json:"title"`
	Sku         string      `json:"sku"`
	Barcodes    []string    `json:"barcodes"`
	Description string      `json:"description"`
	Attributes  []Attribute `json:"attributes"`
	Price       float64     `json:"price"`
	Created     time.Time   `json:"created"`
	LastUpdated time.Time   `json:"lastUpdated"`
}

// HasElem is a funciton to test deep equality of structs
func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if reflect.DeepEqual(arrV.Index(i).Interface(), elem) {
				return true
			}
		}
	}
	return false
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	start := v.Get("start")
	if len(start) == 0 {
		start = "0"
	}
	num := v.Get("num")
	if len(num) == 0 {
		num = "10"
	}
	sku := v.Get("sku")
	if len(sku) == 0 {
		sku = "%"
	}
	barcode := v.Get("barcode")
	if len(barcode) == 0 {
		barcode = "%"
	}

	var getAllProducts GetAllProducts
	var items []Item

	countRes, err := db.Query("SELECT COUNT(product_id) FROM sitoo_test_assignment.product")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		// log.Fatal(err)
	}
	defer countRes.Close()
	for countRes.Next() {
		if err := countRes.Scan(&getAllProducts.TotalCount); err != nil {
			// log.Fatal(err)
			errorHandler(w, r, http.StatusInternalServerError)
		}
	}

	itemRes, err := db.Query("SELECT p.product_id, p.title FROM sitoo_test_assignment.product p LEFT JOIN sitoo_test_assignment.product_barcode b ON p.product_id = b.product_id WHERE p.sku LIKE ? AND b.barcode LIKE ? LIMIT ? OFFSET ?", sku, barcode, num, start)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		// log.Fatal(err)
	}
	defer itemRes.Close()

	for itemRes.Next() {
		var item Item
		err := itemRes.Scan(&item.ProductID, &item.Title)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			// log.Fatal(err)
		}
		// fmt.Println(item)
		items = append(items, item)
	}

	getAllProducts.Items = items
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllProducts)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT p.product_id, p.title, p.sku, b.barcode, p.description, a.name, a.value, p.price, p.created, p.last_updated FROM sitoo_test_assignment.product p LEFT JOIN sitoo_test_assignment.product_barcode b ON p.product_id = b.product_id LEFT JOIN sitoo_test_assignment.product_attribute a ON p.product_id = a.product_id WHERE p.product_id = ? ", params["id"])
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	defer result.Close()
	var getProduct GetProduct
	var barcodes []string
	var barcode string
	var desc sql.NullString
	var attributes []Attribute
	var attribute Attribute
	i := 0
	for result.Next() {
		err := result.Scan(&getProduct.ProductID, &getProduct.Title, &getProduct.Sku, &barcode, &desc, &attribute.Name, &attribute.Value, &getProduct.Price, &getProduct.Created, &getProduct.LastUpdated)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		if desc.Valid {
			getProduct.Description = desc.String
		} else {
			getProduct.Description = "null"
		}
		if !HasElem(attributes, attribute) {
			attributes = append(attributes, attribute)
		}
		if !HasElem(barcodes, barcode) {
			barcodes = append(barcodes, barcode)
		}
		i++
	}
	if i == 0 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	getProduct.Barcodes = barcodes
	getProduct.Attributes = attributes
	json.NewEncoder(w).Encode(getProduct)
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

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		w.Write([]byte(`{"404": "Page Not Found"}`))
		break
	case http.StatusInternalServerError:
		w.Write([]byte(`"500": "Internal Server Error !"`))
		break
	default:
		w.Write([]byte(`"500": "Internal Server Error !"`))
	}
}

func main() {
	db, err = sql.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/sitoo_test_assignment?parseTime=true&")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connection Established")
	}
	defer db.Close()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/products").Subrouter()
	api.HandleFunc("", getProducts).Methods(http.MethodGet)
	api.HandleFunc("/{id}", getProduct).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
