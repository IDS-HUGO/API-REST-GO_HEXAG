package entities

type product struct {
	id    int32
	name  string
	price float32
}

func NewProduct(name string, price float32) *product {
	return &product{
		name:  name,
		price: price,
	}
}

func (p *product) GetID() int32 {
	return p.id
}

func (p *product) SetID(id int32) {
	p.id = id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) SetName(name string) {
	p.name = name
}

func (p *product) GetPrice() float32 {
	return p.price
}

func (p *product) SetPrice(price float32) {
	p.price = price
}
