package main

import(

	"net/http"
	"github.com/gorilla/mux"
	"ProductModel"
	"ProductController"
)

type Product = ProductModel.Product


func main(){
	r := mux.NewRouter()
	productRoutes := r.PathPrefix("/products").Subrouter()

	productRoutes.HandleFunc("/page/{page}/pagesize/{pageSize}", ProductController.ListProducts).Methods("GET").Host("localhost").Schemes("http")
	productRoutes.HandleFunc("/", ProductController.CreateProduct).Methods("POST").Host("localhost").Schemes("http")
	productRoutes.HandleFunc("/{id}", ProductController.ListProduct).Methods("GET").Host("localhost").Schemes("http")
	productRoutes.HandleFunc("/", ProductController.UpdateProduct).Methods("PUT").Host("localhost").Schemes("http")
	productRoutes.HandleFunc("/{id}", ProductController.DeleteProduct).Methods("DELETE").Host("localhost").Schemes("http")
	
	http.ListenAndServe(":5000", r)
}

