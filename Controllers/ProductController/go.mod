module ProductController

go 1.16

replace ProductModel => ../../Models/ProductModel

require (
	ProductModel v0.0.0-00010101000000-000000000000
	ProductService v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace ProductService => ../../Services/ProductService
