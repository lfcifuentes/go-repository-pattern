package users

import (
	"GoRepositoryPattern/responses"
	"GoRepositoryPattern/users/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type UserController struct {
	UserRepository UserRepository
}

func NewUserController(repo UserRepository) *UserController {
	return &UserController{
		UserRepository: repo,
	}
}

func (uc *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserRepository.GetAll()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	response := responses.ResponseOk(users, "pong")
	render.JSON(w, r, response)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	newUser := &domain.User{
		Name:  "NombreUsuario",
		Email: "correo@example.com",
	}

	err := uc.UserRepository.Create(newUser)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := responses.ResponseOk(nil, "pong")
	render.JSON(w, r, response)
}

func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := uc.UserRepository.GetID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := responses.ResponseOk(user, "pong")

	render.JSON(w, r, response)
}
