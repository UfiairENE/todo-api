package controller

import (
	"encoding/json"
	"net/http"
)


func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc ("/ping", ping())
	return mux
}