package handlers

import (
	"API_REST/src/application"
	"encoding/json"
	"net/http"
	"strconv"
)

var productService = application.ProductService{}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var product map[string]interface{}
		json.NewDecoder(r.Body).Decode(&product)
		productService.CreateProduct(product["name"].(string), float32(product["price"].(float64)))
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		products, _ := productService.GetAllProducts()
		json.NewEncoder(w).Encode(products)
	case http.MethodPut:
		var product map[string]interface{}
		json.NewDecoder(r.Body).Decode(&product)
		id, _ := strconv.Atoi(product["id"].(string))
		productService.UpdateProduct(int32(id), product["name"].(string), float32(product["price"].(float64)))
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		productService.DeleteProduct(int32(id))
		w.WriteHeader(http.StatusNoContent)
	}
}
