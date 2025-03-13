package repositories

import (
	"cadastro-usuarios-go/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user models.User) (int, error) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	var id int
	err := r.DB.QueryRow(query, user.Name, user.Email).Scan(&id)
	return id, err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user models.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
