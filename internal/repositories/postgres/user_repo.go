package postgres

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Insert(user *models.User) error {
	query := `INSERT INTO users (id, email, password, user_type) VALUES ($1, $2, $3, $4)`
	_, err := ur.db.Exec(query, user.UserId, user.Email, user.Password, user.UserType)
	return err
}

func (ur *userRepository) GetById(id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, user_type FROM users WHERE id = $1`
	if err := ur.db.QueryRow(query, id).Scan(&user.UserId, &user.Email, &user.Password, &user.UserType); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Get(id, password string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, user_type FROM users WHERE id = $1 AND password = $2`
	if err := ur.db.QueryRow(query, id, password).Scan(&user.UserId, &user.Email, &user.Password, &user.UserType); err != nil {
		return nil, err
	}
	return user, nil
}
