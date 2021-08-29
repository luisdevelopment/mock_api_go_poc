package ProductController

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"ProductModel"
	"ProductService"
)

type Product = ProductModel.Product

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var product Product
	var product2 *Product
	json.NewDecoder(r.Body).Decode(&product)

	product2 = ProductService.CreateProduct(product)

	json.NewEncoder(w).Encode(product2)
}

func ListProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {fmt.Println(err)}

	json.NewEncoder(w).Encode(ProductService.ListProduct(id, 1, 10000))
}

func ListProducts(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	page_size, err := strconv.Atoi(vars["pageSize"])

	if err != nil {fmt.Println(err)}

	json.NewEncoder(w).Encode(ProductService.ListProducts(page, page_size))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	var product Product
	var rowsAffected int
	json.NewDecoder(r.Body).Decode(&product)

	rowsAffected = ProductService.UpdateProduct(product)

	json.NewEncoder(w).Encode(rowsAffected)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {fmt.Println(err)}

	json.NewEncoder(w).Encode(ProductService.DeleteProduct(id))

}