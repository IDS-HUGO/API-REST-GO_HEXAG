package polling

import (
	"database/sql"
	"log"
	"net/http"
)

// Inicia el servidor para el polling
func StartPollingServer(db *sql.DB) {
	http.HandleFunc("/polling/short/users", ShortPollingUsers(db))
	http.HandleFunc("/polling/short/products", ShortPollingProducts(db))
	http.HandleFunc("/polling/long/products", LongPollingNewProduct(db))

	log.Println("Servidor de polling ejecutándose en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
