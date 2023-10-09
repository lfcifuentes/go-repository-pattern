package responses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Data string
}

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

func TestResponseStruct(t *testing.T) {
	response := Response{
		Data:    "Data test",
		Message: "Message test",
	}

	assert.Equal(t, response.Data, "Data test")
	assert.Equal(t, response.Message, "Message test")
}
