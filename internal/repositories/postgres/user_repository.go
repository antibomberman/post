package postgres

import (
	"antibomberman/post/internal/models"
)

type UserRepository interface {
	CreateUser(user models.UserCreate) error
	ExistsUserByEmail(email string) (bool, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id string) (models.User, error)
	//UpdateUser(user *models.User) error
	//DeleteUser(id int) error
}

type userRepository struct {
	db *Postgres
}

func NewUserRepository(db *Postgres) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(data models.UserCreate) error {
	_, err := u.db.NamedExec("INSERT INTO users (name, email, password) VALUES (:name, :email, :password)", data)
	if err != nil {
		return err
	}
	return nil

}
func (u *userRepository) ExistsUserByEmail(email string) (bool, error) {

	var count int
	err := u.db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *userRepository) GetUserByEmail(email string) (models.User, error) {

	var user models.User
	err := u.db.Get(&user, "SELECT id, name, email, password, created_at FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (u *userRepository) GetUserById(id string) (models.User, error) {
	var user models.User
	err := u.db.Get(&user, "SELECT id, name, email, password, created_at FROM users WHERE id = $1", id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
