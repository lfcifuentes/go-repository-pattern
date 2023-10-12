package model

import (
	"encoding/json"
	"fmt"
	"github.com/lfcifuentes/go-repository-pattern/app/security"
)

type Users struct {
	Data *[]User
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

func UnmarshalUser(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}

func (u *User) CreatePasswordHash() string {
	hash, err := security.HashPassword(u.Password)

	fmt.Println(err)

	return hash
}
