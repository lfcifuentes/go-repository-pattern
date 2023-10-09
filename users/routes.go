package users

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
)

func NewRouter(r *chi.Mux, db *sql.DB) *chi.Mux {

	userRepository := NewSQLUserRepository(db)
	userController := NewUserController(userRepository)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAll)
		r.Post("/", userController.Create)
	})

	return r
}
