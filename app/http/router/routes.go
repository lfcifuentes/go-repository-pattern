package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/lfcifuentes/go-repository-pattern/app/http/handler"
	"github.com/lfcifuentes/go-repository-pattern/app/repository"
)

func NewRouter(r *chi.Mux, db *sql.DB) *chi.Mux {

	userRepository := repository.NewSQLUserRepository(db)
	userController := handler.NewUserController(userRepository)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAll)
		r.Post("/", userController.Create)
	})

	return r
}
