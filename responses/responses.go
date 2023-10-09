package responses

import (
	"github.com/go-chi/render"
	"net/http"
)

func ResponseOk(data interface{}, message string) Response {
	return Response{
		data,
		message,
	}
}

func baseResponseError(w http.ResponseWriter, r *http.Request, data interface{}, message string, status int) {
	response := Response{
		data,
		message,
	}
	render.Status(r, status)
	render.JSON(w, r, response)
}

func ResponseError(w http.ResponseWriter, r *http.Request, data interface{}, message string) {
	baseResponseError(w, r, data, message, http.StatusUnprocessableEntity)
}

func ResponseUnauthorized(w http.ResponseWriter, r *http.Request) {
	baseResponseError(w, r, nil, "Unauthorized", http.StatusUnauthorized)
}
func ResponseUnprocessableEntity(w http.ResponseWriter, r *http.Request) {
	baseResponseError(w, r, nil, "Unauthorized", http.StatusUnprocessableEntity)
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}
