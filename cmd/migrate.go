package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lfcifuentes/go-repository-pattern/database"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"log"
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
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	switch os.Getenv("DB_CONNECTION") {
	case "psql":
		migratePsql()
		break
	case "mysql":
		migrateMysql()
		break
	default:
		log.Fatal("DB CONNECTION NOT SUPPORTED")
	}
}

func migratePsql() {
	dbName := os.Getenv("DB_NAME")

	// Crear la cadena de conexión para PostgreSQL sin especificar la base de datos
	dsn := database.CreateConnectionWithoutDbString()

	// Intenta conectar a PostgreSQL sin seleccionar una base de datos específica
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal("Error pining to PostgreSQL:", err)
	}

	defer db.Close()

	// Ejecuta la sentencia SQL para crear la base de datos si no existe
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		log.Fatal("Error creating database:", err)
	}
	fmt.Println("Database created or already exists.")

	// Conectar a la base de datos recién creada
	dsn = fmt.Sprintf("%s dbname=%s sslmode=disable", dsn, dbName)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the created database:", err)
	}
	defer db.Close()

	// Ejecuta las sentencias SQL desde schema.sql para crear tablas u otras estructuras de base de datos
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		name VARCHAR(255),
		password VARCHAR(255),
		email VARCHAR(255)
	);`)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
	fmt.Println("Tables created successfully.")
}

func migrateMysql() {
	// DB credentials
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
