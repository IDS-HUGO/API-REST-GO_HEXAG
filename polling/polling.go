package polling

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

// Short Polling para nuevos usuarios
func ShortPollingUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lastChecked := time.Now().Add(-5 * time.Second)
		query := "SELECT COUNT(*) FROM users WHERE created_at > ?"
		var newUsers int

		err := db.QueryRow(query, lastChecked).Scan(&newUsers)
		if err != nil {
			http.Error(w, "Error en la consulta de usuarios", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Nuevos usuarios registrados: %d", newUsers)
	}
}

// Short Polling para productos actualizados
func ShortPollingProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lastChecked := time.Now().Add(-5 * time.Second)
		query := "SELECT COUNT(*) FROM products WHERE updated_at > ?"
		var updatedProducts int

		err := db.QueryRow(query, lastChecked).Scan(&updatedProducts)
		if err != nil {
			http.Error(w, "Error en la consulta de productos", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Productos actualizados: %d", updatedProducts)
	}
}

// Long Polling para detectar nuevos productos
func LongPollingNewProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lastChecked := time.Now()

		for {
			query := "SELECT id, name FROM products WHERE created_at > ? LIMIT 1"
			var id int
			var name string

			err := db.QueryRow(query, lastChecked).Scan(&id, &name)
			if err != nil && err != sql.ErrNoRows {
				http.Error(w, "Error en la consulta de productos", http.StatusInternalServerError)
				return
			}

			if err == nil { // Nuevo producto encontrado
				fmt.Fprintf(w, "Nuevo producto agregado: %s (ID: %d)", name, id)
				return
			}

			time.Sleep(2 * time.Second) // Intervalo antes de la siguiente consulta
		}
	}
}
