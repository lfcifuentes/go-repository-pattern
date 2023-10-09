package users

import (
	"GoRepositoryPattern/users/domain"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetID(id int) (*domain.User, error)
	Create(user *domain.User) error
	ChangePassword(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int) error
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// GetAll get all users from database
func (r *SQLUserRepository) GetAll() ([]*domain.User, error) {
	query := "SELECT id, name, email FROM users ORDER BY name "

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		user := &domain.User{}
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

func (r *SQLUserRepository) GetID(id int) (*domain.User, error) {

	query := "SELECT id, name, email FROM users WHERE id = ?"

	var user domain.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *SQLUserRepository) Create(user *domain.User) error {

	query := "INSERT INTO repositorydb.users (id, name, email) values ( ?, ?, ?);"

	row, err := r.db.Query(query, 1, user.Name, user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User not found")
		}
		return err
	}

	fmt.Println(row.Next())

	return errors.New("ERROR")
}

func (r *SQLUserRepository) ChangePassword(user *domain.User) error {
	// Implementa la lógica para crear un usuario en la base de datos.
	return errors.New("ERROR")
}

func (r *SQLUserRepository) Update(user *domain.User) error {
	// Implementa la lógica para actualizar un usuario en la base de datos.
	return errors.New("ERROR")
}

func (r *SQLUserRepository) Delete(id int) error {
	// Implementa la lógica para eliminar un usuario de la base de datos.
	return errors.New("ERROR")
}
