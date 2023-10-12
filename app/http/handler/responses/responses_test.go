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

// getHttpMuk returns a fake HTTP response recorder and request.
func getHttpMuk() (*httptest.ResponseRecorder, *http.Request) {
	// Create a fake HTTP request and a fake ResponseWriter.
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	return w, req
}

// TestResponseStruct tests the Response struct.
func TestResponseStruct(t *testing.T) {
	response := Response{
		Data:    "Data test",
		Message: "Message test",
	}

	assert.Equal(t, response.Data, "Data test")
	assert.Equal(t, response.Message, "Message test")
}

// TestResponseOk tests the ResponseOk function.
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

// TestResponseBaseError tests the baseResponseError function.
func TestResponseBaseError(t *testing.T) {
	w, req := getHttpMuk()
	// Call the baseResponseError function with test data
	data := map[string]interface{}{"key": "value"}
	message := "Error message"
	status := http.StatusBadRequest
	baseResponseError(w, req, data, message, status)

	assert.Contains(t, w.Body.String(), "value")
	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseError tests the ResponseError function.
func TestResponseError(t *testing.T) {
	w, req := getHttpMuk()

	// Call the ResponseError function with test data
	data := map[string]interface{}{"key": "value"}
	message := "Error message"
	status := http.StatusUnprocessableEntity
	ResponseError(w, req, data, message)

	assert.Contains(t, w.Body.String(), "value")
	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseStatusUnauthorizedError tests the ResponseUnauthorized function.
func TestResponseStatusUnauthorizedError(t *testing.T) {
	w, req := getHttpMuk()

	status := http.StatusUnauthorized
	message := "Unauthorized"
	ResponseUnauthorized(w, req)

	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}

// TestResponseStatusUnprocessableEntityError tests the ResponseUnprocessableEntity function.
func TestResponseStatusUnprocessableEntityError(t *testing.T) {
	w, req := getHttpMuk()

	status := http.StatusUnprocessableEntity
	message := "Unauthorized"
	ResponseUnprocessableEntity(w, req)

	assert.Contains(t, w.Body.String(), message)
	assert.Equal(t, w.Code, status)
}
