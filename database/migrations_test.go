package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// TestCreateConnectionPsqlString
func TestCreateConnectionPsqlString(t *testing.T) {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env.psql-test")
	if err != nil {
		t.Fatalf("Error loading .env.ci file: %v", err)
	}

	expectedUser := os.Getenv("DB_USER")
	expectedPass := os.Getenv("DB_PASSWORD")
	expectedDB := os.Getenv("DB_NAME")
	expectedPort := os.Getenv("DB_PORT")
	expectedHost := os.Getenv("DB_HOST")

	expectedConnectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s sslmode=disable dbname=%s",
		expectedUser, expectedPass, expectedHost, expectedPort, expectedDB,
	)
	fmt.Println(expectedConnectionString)
	// Llamar a la función para crear la cadena de conexión
	actualConnectionString := CreateConnectionString()

	// Verificar que la cadena de conexión generada coincide con la esperada
	assert.Equal(t, expectedConnectionString, actualConnectionString)
}

// TestCreateConnectionPsqlWithoutDbString
func TestCreateConnectionPsqlWithoutDbString(t *testing.T) {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env.psql-test")
	if err != nil {
		t.Fatalf("Error loading .env.ci file: %v", err)
	}

	expectedUser := os.Getenv("DB_USER")
	expectedPass := os.Getenv("DB_PASSWORD")
	expectedPort := os.Getenv("DB_PORT")
	expectedHost := os.Getenv("DB_HOST")

	expectedConnectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s sslmode=disable",
		expectedUser, expectedPass, expectedHost, expectedPort,
	)
	// Llamar a la función para crear la cadena de conexión
	actualConnectionString := CreateConnectionWithoutDbString()

	// Verificar que la cadena de conexión generada coincide con la esperada
	assert.Equal(t, expectedConnectionString, actualConnectionString)
}

// TestCreatePsqlConnection
func TestCreatePsqlConnection(t *testing.T) {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env.psql-test")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Crear la cadena de conexión para PostgreSQL sin especificar la base de datos
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		user, pass, host, port)

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
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", database))
	if err != nil {
		log.Fatal("Error deleting database:", err)
	}
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", database))
	if err != nil {
		log.Fatal("Error creating database:", err)
	}
	fmt.Println("Database created or already exists.")

	fmt.Println(CreateConnectionString())
	// Llamar a la función para crear la conexión
	db, err = CreateConnection()
	defer db.Close() // Asegúrate de cerrar la conexión

	// Verificar que no haya errores al crear la conexión
	assert.NoError(t, err)

	// Verificar que la conexión esté abierta
	assert.NotNil(t, db)

	// Intentar ejecutar una consulta simple para verificar la conexión a la base de datos
	err = db.Ping()
	assert.NoError(t, err)
}
