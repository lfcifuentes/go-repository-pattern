package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lfcifuentes/go-repository-pattern/app/model"
)

type UserRepository interface {
	GetAll() ([]*model.User, error)
	GetID(id int) (*model.User, error)
	Create(user *model.User) error
	ChangePassword(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// GetAll get all users from database
func (r *SQLUserRepository) GetAll() ([]*model.User, error) {
	query := "SELECT id, name, email FROM users ORDER BY name "

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *SQLUserRepository) GetID(id int) (*model.User, error) {

	query := "SELECT id, name, email FROM users WHERE id = ?"

	var user model.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *SQLUserRepository) Create(user *model.User) error {

	query := "INSERT INTO public.users (name, password, email) VALUES ($1, $2, $3);"

	_, err := r.db.Exec(query, user.Name, user.CreatePasswordHash(), user.Email)

	if err != nil {
		fmt.Println(err)
	}

	return errors.New("ERROR")
}

func (r *SQLUserRepository) ChangePassword(user *model.User) error {
	// Implementa la lógica para crear un usuario en la base de datos.
	return errors.New("ERROR")
}

func (r *SQLUserRepository) Update(user *model.User) error {
	// Implementa la lógica para actualizar un usuario en la base de datos.
	return errors.New("ERROR")
}

func (r *SQLUserRepository) Delete(id int) error {
	// Implementa la lógica para eliminar un usuario de la base de datos.
	return errors.New("ERROR")
}
