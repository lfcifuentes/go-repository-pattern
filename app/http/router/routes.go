package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/lfcifuentes/go-repository-pattern/app/http/handler"
	"github.com/lfcifuentes/go-repository-pattern/app/repository"
	"os"
	"strings"
	"time"
)

func NewRouter(db *sql.DB) *chi.Mux {

	userRepository := repository.NewSQLUserRepository(db)
	userController := handler.NewUserController(userRepository)

	r := chi.NewRouter()

	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: strings.Split(os.Getenv("ACCEPTED_DOMAINS"), ","),
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAll)
		r.Post("/", userController.Create)
	})

	return r
}
