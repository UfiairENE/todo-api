package main

import (
	"net/http"
	"github.com/UfiairENE/todo-api/controller"


)

func main() {
	mux := controller.Register()
	
	http.ListenAndServe(":3000", mux)
}