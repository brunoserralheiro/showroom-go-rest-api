package main

import (
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
	Owner User   `json:"user"`
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
	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	r.HandleFunc("/api/v1/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", postProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/upload/{id}", upload).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
	// http.Handle("/", r)

} // end main

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	prods := mockProducts()
	fmt.Fprintf(w, "Products: %v\n", prods)

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

//TODO
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

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
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	check(err)
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	check(err)
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

// ===============

func mockProducts() *[]Product {

	user0 := User{
		ID:    "0101",
		Name:  "name",
		Email: "email",
		Rank:  "seller", //master, admin, manager, seller, buyer
	}
	user2 := User{
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
			price: "price",
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
