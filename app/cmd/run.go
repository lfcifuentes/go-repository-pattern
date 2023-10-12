package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lfcifuentes/go-repository-pattern/app/http/router"
	"github.com/lfcifuentes/go-repository-pattern/database"
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

	// Configura el enrutador
	router := router.NewRouter(db)

	fmt.Println(fmt.Sprintf("Server is running in port %v", "8080"))
	// Inicia el servidor web
	http.ListenAndServe(":8080", router)
}
