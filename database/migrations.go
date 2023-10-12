package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"os"
)

// CreateConnectionString create database connection string
func CreateConnectionString() string {
	dbName := os.Getenv("DB_NAME")
	// DB credentials
	dns := CreateConnectionWithoutDbString()

	// Crear la cadena de conexión para PostgreSQL sin especificar la base de datos
	return fmt.Sprintf("%s dbname=%s", dns, dbName)
}

// CreateConnectionWithoutDbString
func CreateConnectionWithoutDbString() string {
	// DB credentials
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Crear la cadena de conexión para PostgreSQL sin especificar la base de datos
	return fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		user, pass, host, port)
}

// CreateConnection Create a new database connection
func CreateConnection() (*sql.DB, error) {
	// Configura la base de datos
	db, err := sql.Open(
		"postgres",
		CreateConnectionString(),
	)

	return db, err
}
