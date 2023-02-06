package connection

import (
	"database/sql"
	"fmt"
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
	res, err := DB.Query(fmt.Sprintf(`
		CREATE TABLE gomigrate_migrators (
			id SERIAL PRIMARY KEY,
			Name VARCHAR(255) NOT NULL,
			Version INT NOT NULL DEFAULT 1,
			CreatedAt BIGINT NOT NULL
		);
	`))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	res.Close()

	res, err = DB.Query(fmt.Sprintf(`
		CREATE TABLE gomigrate_migrator_stash (
			id SERIAL PRIMARY KEY,
			Name VARCHAR(255) NOT NULL,
			Version INT NOT NULL,
			upFileBody TEXT,
			downFileBody TEXT,
			ChangedAt BIGINT NOT NULL
		);
	`))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	res.Close()
}
