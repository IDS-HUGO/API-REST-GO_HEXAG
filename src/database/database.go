package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/api_rest")
	if err != nil {
		panic(err)
	}
	fmt.Println("Base de datos conectada correctamente.")
}
