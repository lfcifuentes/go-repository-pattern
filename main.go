package main

import (
	"GoRepositoryPattern/database"
	"GoRepositoryPattern/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// Configura la base de datos
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Configura el enrutador
	router := users.NewRouter(r, db)
	router = users.NewRouter(router, db)

	// Inicia el servidor web
	http.ListenAndServe(":8080", router)
}
