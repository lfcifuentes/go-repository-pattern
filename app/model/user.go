package model

import (
	"fmt"
	"github.com/lfcifuentes/go-repository-pattern/app/security"
)

type Users struct {
	Data *[]User
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) CreatePasswordHash() string {
	hash, err := security.HashPassword(u.Password)

	fmt.Println(err)

	return hash
}
