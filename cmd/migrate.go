package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run the application migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Lógica para ejecutar la aplicación aquí
		fmt.Println("Running migrations...")
		migrate()
	},
}

func migrate() {
	// DB credentials
	godotenv.Load(".env")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_ROOT_PASSWORD")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(localhost:3306)/", user, pass)

	// Intenta conectar a MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error al conectar a MySQL:", err)
		return
	}
	defer db.Close()

	// Ejecuta las sentencias SQL desde schema.sql
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v;", database))
	if err != nil {
		fmt.Println("Error al ejecutar las sentencias SQL:", err)
		return
	}
	// Ejecuta las sentencias SQL desde schema.sql
	_, err = db.Exec(fmt.Sprintf("USE %v;", database))
	if err != nil {
		fmt.Println("Error al ejecutar las sentencias SQL:", err)
		return
	}

	fmt.Println("Database created successfully")

	// Ejecuta las sentencias SQL desde schema.sql
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), password VARCHAR(255), email VARCHAR(255));")
	if err != nil {
		fmt.Println("Error al ejecutar las sentencias SQL:", err)
		return
	}
	fmt.Println("Tables created successfully.")
}
