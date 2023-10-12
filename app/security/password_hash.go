package security

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

// LoadPepper loads the pepper from an environment variable.
func LoadPepper() string {
	pepper := os.Getenv("APP_SECRET")
	if pepper == "" {
		panic("Variable de entorno APP_SECRET no definida")
	}
	return pepper
}

// HashPassword hashes the provided password using the pepper.
func HashPassword(password string) (string, error) {
	passwordWithPepper := password + LoadPepper()

	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithPepper), 14)

	return string(bytes), err
}

// CheckPasswordHash verifies a user's password against a hashed password with the pepper.
func CheckPasswordHash(userPassword, hash string) bool {
	userPasswordWithPepper := userPassword + LoadPepper()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userPasswordWithPepper))
	return err == nil
}
