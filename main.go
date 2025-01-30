package main

import (
	g"
	t/http"
	oyecto/application"
	oyecto/database"
	oyecto/infrastructure/handlers"
	oyecto/infrastructure/repositories"

	thub.com/gorilla/mux"
)

func main() {
	// Conectar a la base de datos
	db, err := database.NewMySQLDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Inicializar repositorios
	productRepo := repositories.NewMySQLProductRepository(db)
	userRepo := repositories.NewMySQLUserRepository(db)

	// Inicializar servicios
	productService := application.NewProductService(productRepo)
	userService := application.NewUserService(userRepo)

	// Inicializar handlers
	productHandler := handlers.NewProductHandler(productService)
	userHandler := handlers.NewUserHandler(userService)

	// Configurar el enrutador
	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
