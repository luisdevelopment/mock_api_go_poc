package ProductModel

type Product struct {
	Id int `json:id`
	Name string `json:name`
	Valor float64 `json:valor`
}


func NewProduct(id int, name string, valor float64) *Product {
	p := &Product{
		Id : id,
		Name : name,
		Valor : valor,
	}
	return p
  }

