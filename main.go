package main

import (
	"net/http"

)

func main() {
	mux := controller.Register()
	
	http.ListenAndServe(":3000", mux)
}