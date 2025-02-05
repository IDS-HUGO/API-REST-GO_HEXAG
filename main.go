package main

import (
	"API_REST/src/database"
	"API_REST/src/infrastructure/handlers"
	"fmt"
	"net/http"
)

func main() {
	database.InitDB()

	http.HandleFunc("/products", handlers.ProductHandler)
	http.HandleFunc("/users", handlers.UserHandler)

	port := ":8080"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	http.ListenAndServe(port, nil)
}
