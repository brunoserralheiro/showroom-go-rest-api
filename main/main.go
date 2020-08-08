package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner *User  `json:"user"`
	price string `json:"price"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Rank  string `json:"rank"` //master, admin, manager, seller, buyer
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/v1/upload", getProducts).Methods("POST")
	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	r.HandleFunc("/api/v1/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", postProduct).Methods("POST")
	r.HandleFunc("/api/v1/upload/{id}", upload).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

} // end main

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HomeHandler")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Print("getProducts")
	// vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	prods := mockProducts()
	json.NewEncoder(w).Encode(prods)

	// fmt.Fprintf(w, "Products: %v\n", prods)

}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}
func postProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}

//@todo organize packets and files...
func upload(w http.ResponseWriter, r *http.Request) {
	log.Print("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		log.Println(err)
		return
	}
	defer file.Close()

	log.Printf("uploaded file:  %+v", handler.Filename)
	log.Printf("file size:  %+v", handler.Size)
	log.Printf("MIME header:  %+v", handler.Header)

	//write temp file in server
	tempFile, err := ioutil.TempFile("../.frontend/images", "")
	check(err)
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	check(err)
	tempFile.Write(fileBytes)
	// return whether this is successful
}

// ===============
// mock data @todo implement DB
func mockProducts() *[]Product {

	user0 := &User{
		ID:    "0101",
		Name:  "name",
		Email: "email",
		Rank:  "seller", //master, admin, manager, seller, buyer
	}
	user2 := &User{
		ID:    "0102",
		Name:  "name2",
		Email: "email2",
		Rank:  "admin", //master, admin, manager, seller, buyer
	}

	productz := []Product{
		Product{
			ID:    "001",
			Name:  "name",
			Owner: user0,
			price: "3000",
		},
		Product{
			ID:    "002",
			Name:  "name2",
			Owner: user2,
			price: "price",
		},
	}
	return &productz
}
