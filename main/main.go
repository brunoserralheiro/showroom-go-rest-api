package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mock "github.com/brunoserralheiro/showroom-go-rest-api/main/mock"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/v1/Upload", GetProducts).Methods("POST")
	r.HandleFunc("/api/v1/products", GetProducts).Methods("GET")
	r.HandleFunc("/api/v1/product/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", NewProduct).Methods("POST")
	r.HandleFunc("/api/v1/upload/{id}", Upload).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

} // end main

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HomeHandler")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Print("GetProducts")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	prods := mock.MockProducts()
	json.NewEncoder(w).Encode(prods)

	// fmt.Fprintf(w, "Products: %v\n", prods)

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	log.Print("BEGIN GetProduct handler function")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for k, v := range params {

		log.Printf("Param  %v: %+v", k, v)

	}

	prods := mock.MockProducts()
	for _, item := range *prods {

		if item.ID == params["id"] {
			log.Printf("Found product: %+v", item.ID)
			json.NewEncoder(w).Encode(item)
		}
	}
	log.Println(w, "Product: %v", params["product"])
	log.Print(" END GetProduct handler function")
}
func NewProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}

//@todo organize packets and files...
func Upload(w http.ResponseWriter, r *http.Request) {
	log.Print("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// Upload of 10 MB files.
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
