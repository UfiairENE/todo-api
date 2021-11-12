package controller

import (
	"encoding/json"
	"net/http"

	"github.com/UfiairENE/todo-api/views"
)
func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body:"pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}