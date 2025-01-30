package repositories

import (
	"database/sql"
	"proyecto/domain/entities"
)

type MySQLProductRepository struct {
	db *sql.DB
}

func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) Save(product *entities.Product) error {
	query := `INSERT INTO products (name, price) VALUES (?, ?)`
	_, err := r.db.Exec(query, product.Name, product.Price)
	return err
}

func (r *MySQLProductRepository) FindByID(id int32) (*entities.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = ?`
	row := r.db.QueryRow(query, id)
	product := &entities.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *MySQLProductRepository) Update(product *entities.Product) error {
	query := `UPDATE products SET name = ?, price = ? WHERE id = ?`
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (r *MySQLProductRepository) Delete(id int32) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
