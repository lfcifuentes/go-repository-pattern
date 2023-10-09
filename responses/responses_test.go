package responses

import (
	"log"
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
	if responseOk.Data != testStruct {
		log.Fatal("Response data is wrong")
	}
	if responseOk.Message != "message" {
		log.Fatal("Response message is wrong")
	}
}

func TestResponseStruct(t *testing.T) {
	response := Response{
		Data:    "Data test",
		Message: "Message test",
	}

	if response.Data != "Data test" {
		log.Fatal("Response data is wrong")
	}
	if response.Message != "Message test" {
		log.Fatal("Response message is wrong")
	}
	/*
		fmt.Println(response)
		// Output:
		// [Data: "Data test 2" Message: "Message test"]
	*/
}
