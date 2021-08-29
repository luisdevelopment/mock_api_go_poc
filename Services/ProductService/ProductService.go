package ProductService

import(
	"ProductModel"
	"strconv"
	"fmt"
)

type Product = ProductModel.Product

func CreateProduct(p Product) *Product {
	p2 := &Product{
		Id : p.Id,
		Name : p.Name,
		Valor : p.Valor,
	}
	return p2
}

func ListProduct(id int, page int, page_size int) Product{
	var products_list []Product

	for i:=0; i<page_size; i++ {
		var product Product

		product.Id = i
		product.Name = "Luis"+strconv.Itoa(i)
		product.Valor = float64(i)

		products_list = append(products_list, product)
	}

	return products_list[id]
}

func ListProducts(page int, page_size int) []Product{
	var products_list []Product

	for i:=0; i<page_size; i++ {
		var product Product

		product.Id = i
		product.Name = "Luis"+strconv.Itoa(i)
		product.Valor = float64(i)

		products_list = append(products_list, product)
	}

	return products_list

}

func UpdateProduct(p Product) int {
	p2 := &Product{
		Id : p.Id,
		Name : p.Name,
		Valor : p.Valor,
	}
	
	fmt.Println(p2.Name)

	return 1
}

func DeleteProduct(id int) int {
	return id
}