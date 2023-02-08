package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitConnection(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	DB = db
}

func InitDatabaseSchemas() {
	content, err := os.ReadFile(".\\internal\\connection\\init.sql")
	if err != nil {
		log.Fatal(err)
	}
	res, err := DB.Query(string(content))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	res.Close()
}
