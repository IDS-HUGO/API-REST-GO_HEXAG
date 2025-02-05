package handlers

import (
	"API_REST/src/application"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var productService = application.ProductService{}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var product map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		name, nameOk := product["name"].(string)
		priceFloat, priceOk := product["price"].(float64)

		if !nameOk || !priceOk {
			http.Error(w, "Invalid name or price", http.StatusBadRequest)
			return
		}

		productService.CreateProduct(name, float32(priceFloat))
		w.WriteHeader(http.StatusCreated)
		fmt.Println("Producto creado:", name, priceFloat)

	case http.MethodGet:
		products, err := productService.GetAllProducts()
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)

	case http.MethodPut:
		var product map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		idFloat, idOk := product["id"].(float64)
		name, nameOk := product["name"].(string)
		priceFloat, priceOk := product["price"].(float64)

		if !idOk || !nameOk || !priceOk {
			http.Error(w, "Invalid id, name, or price", http.StatusBadRequest)
			return
		}

		id := int(idFloat)
		if err := productService.UpdateProduct(int32(id), name, float32(priceFloat)); err != nil {
			http.Error(w, "Error updating product", http.StatusInternalServerError)
			return
		}

		response := map[string]string{"message": "Producto actualizado correctamente"}
		json.NewEncoder(w).Encode(response)
		fmt.Println("Producto actualizado:", id, name, priceFloat)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := productService.DeleteProduct(int32(id)); err != nil {
			http.Error(w, "Error deleting product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		fmt.Println("Producto eliminado:", id)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
