package responses

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Data string
}

// getHttpMuk
func getHttpMuk() (*httptest.ResponseRecorder, *http.Request) {
	// Crea una solicitud HTTP falsa y un ResponseWriter falso
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	return w, req
}

// TestResponseStruct
func TestResponseStruct(t *testing.T) {
	response := Response{
		Data:    "Data test",
		Message: "Message test",
	}

	assert.Equal(t, response.Data, "Data test")
	assert.Equal(t, response.Message, "Message test")
}

// TestResponseOk
func TestResponseOk(t *testing.T) {
	testStruct := TestStruct{
		Data: "Data",
	}
	responseOk := ResponseOk(
		testStruct,
		"message",
	)
	assert.Equal(t, responseOk.Data, testStruct)
	assert.Equal(t, responseOk.Message, "message")
}

// TestResponseBaseError
func TestResponseBaseError(t *testing.T) {
	w, req := getHttpMuk()
	// Llama a la función baseResponseError con datos de prueba
	data := map[string]interface{}{"key": "value"}
	message := "Error message"
	status := http.StatusBadRequest
	baseResponseError(w, req, data, message, status)

	assert.Contains(t, w.Body.String(), "value")
	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseError
func TestResponseError(t *testing.T) {
	w, req := getHttpMuk()

	// Llama a la función baseResponseError con datos de prueba
	data := map[string]interface{}{"key": "value"}
	message := "Error message"
	status := http.StatusUnprocessableEntity
	ResponseError(w, req, data, message)

	assert.Contains(t, w.Body.String(), "value")
	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseStatusUnauthorizedError
func TestResponseStatusUnauthorizedError(t *testing.T) {
	w, req := getHttpMuk()

	status := http.StatusUnauthorized
	message := "Unauthorized"
	ResponseUnauthorized(w, req)

	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseStatusUnprocessableEntityError
func TestResponseStatusUnprocessableEntityError(t *testing.T) {
	w, req := getHttpMuk()

	status := http.StatusUnprocessableEntity
	message := "Unauthorized"
	ResponseUnprocessableEntity(w, req)

	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}
