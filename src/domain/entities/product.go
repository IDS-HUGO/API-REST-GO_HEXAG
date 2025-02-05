package entities

type Product struct {
	ID    int32
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{ID: 1, Name: name, Price: price}
}
