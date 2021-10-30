package main

import (
	"encoding/json"
	"net/http"

	"github.com/UfiairENE/todo API/struct/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc ("/beep", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body:"beep",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe("localhost:8000", mux)
}