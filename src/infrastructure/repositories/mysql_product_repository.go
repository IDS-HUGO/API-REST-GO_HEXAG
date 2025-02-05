package repositories

import (
	"API_REST/src/database"
	"API_REST/src/domain/entities"
)

type ProductRepository struct{}

func (r *ProductRepository) Create(product *entities.Product) error {
	_, err := database.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	return err
}

func (r *ProductRepository) GetAll() ([]entities.Product, error) {
	rows, err := database.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) Update(product *entities.Product) error {
	_, err := database.DB.Exec("UPDATE products SET name=?, price=? WHERE id=?", product.Name, product.Price, product.ID)
	return err
}

func (r *ProductRepository) Delete(id int32) error {
	_, err := database.DB.Exec("DELETE FROM products WHERE id=?", id)
	return err
}
