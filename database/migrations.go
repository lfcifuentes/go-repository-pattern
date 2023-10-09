package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func CreateConnectionString() string {
	// DB credentials
	godotenv.Load(".env")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"%v:%v@tcp(localhost:3306)/%v",
		user,
		pass,
		database,
	)
}

// CreateConnection Create a new database connection
func CreateConnection() (*sql.DB, error) {
	// Configura la base de datos
	db, err := sql.Open(
		"mysql",
		CreateConnectionString(),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
