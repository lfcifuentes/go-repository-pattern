package security

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadPepperPannic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "Variable de entorno APP_SECRET no definida", r)
		}
	}()
	LoadPepper()
	assert.Fail(t, "La función LoadPepper no entró en pánico como se esperaba")
}

func TestLoadPepper(t *testing.T) {
	pepper := "my_secure_pepper"
	_ = os.Setenv("APP_SECRET", pepper)

	assert.Equal(t, pepper, LoadPepper())
}

func TestHashPassword(t *testing.T) {
	_ = os.Setenv("APP_SECRET", "my_secure_pepper")

	password := "my_password"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCheckPasswordHash(t *testing.T) {
	_ = os.Setenv("APP_SECRET", "my_secure_pepper")

	password := "my_password"
	hashedPassword, _ := HashPassword(password)

	assetHash := CheckPasswordHash(password, hashedPassword)

	assert.Equal(t, true, assetHash)
}
