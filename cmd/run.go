package cmd

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lfcifuentes/go-repository-pattern/database"
	"github.com/lfcifuentes/go-repository-pattern/users"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var Run = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Run: func(cmd *cobra.Command, args []string) {
		// Lógica para ejecutar la aplicación aquí
		fmt.Println("Running the application server...")
		runServer()
	},
}

func runServer() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
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

	// Inicia el servidor web
	http.ListenAndServe(":8080", router)
}
