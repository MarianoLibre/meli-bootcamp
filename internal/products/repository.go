package products

import "fmt"

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Colour    string  `json:"colour"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"created-at"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	LastID() (int, error)
	Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	p := Product{id, name, colour, price, stock, code, published, createdAt}
	ps = append(ps, p)
	lastID = p.Id
	fmt.Println("REPOSITORY>>> ", id, name, colour, code, createdAt)
	return p, nil
}

func (r *repository) Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	p := Product{Name: name, Colour: colour, Code: code, CreatedAt: createdAt, Stock: stock, Price: price, Published: published}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}
