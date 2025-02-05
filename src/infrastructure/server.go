package infrastructure

import (
	"API_REST/src/infrastructure/handlers"
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/products", handlers.ProductHandler)
	http.HandleFunc("/users", handlers.UserHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
