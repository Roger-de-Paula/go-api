package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-jwt-api/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Hello, World!")
	fmt.Println("Listening on port 8080")

	api.JWTKey = []byte("secret123")

	r.HandleFunc("/login", api.HandleLogin).Methods("POST")
	r.HandleFunc("/protected", api.HandleProtectedRoute).Methods("GET")

	http.ListenAndServe(":8000", r)
}
