package repositories

import (
	"API_REST/src/database"
	"API_REST/src/domain/entities"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *entities.User) error {
	_, err := database.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	return err
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	rows, err := database.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user *entities.User) error {
	_, err := database.DB.Exec("UPDATE users SET username=?, email=? WHERE id=?", user.Username, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id int32) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
