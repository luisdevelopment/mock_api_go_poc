module main

go 1.16

require (
	ProductController v0.0.0-00010101000000-000000000000
	ProductModel v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace ProductController => ./Controllers/ProductController

replace ProductModel => ./Models/ProductModel

replace ProductService => ./Services/ProductService
