package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost port=5432 user=admin password=secret dbname=go_users sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Banco inacessível:", err)
	}

	fmt.Println("Conectado ao banco PostgreSQL!")
}
