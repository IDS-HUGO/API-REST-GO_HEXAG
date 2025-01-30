package repositories

import (
	tabase/sql"
	oyecto/domain/entities"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *entities.User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := r.db.Exec(query, user.Name, user.Email)
	return err
}

func (r *MySQLUserRepository) FindByID(id int32) (*entities.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) Update(user *entities.User) error {
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *MySQLUserRepository) Delete(id int32) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
