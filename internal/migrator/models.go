package migrator

import (
	"database/sql"
	"fmt"
	"gomigrate/internal/connection"
	"log"
	"time"
)

func ParseResults(res *sql.Rows) []Migration {
	var list []Migration
	for res.Next() {
		m := Migration{}
		err := res.Scan(&m.Id, &m.Name, &m.Version, &m.CreatedAt)
		if err != nil {
			log.Println(err)
		}
		list = append(list, m)
	}
	return list
}

func Insert(name string) Migration {
	res, err := connection.DB.Query(fmt.Sprintf(`
		INSERT INTO
		gomigrate_migrators 
		(Name, CreatedAt) 
		VALUES('%s', %d)
		RETURNING *;`,
		name, time.Now().Unix()))

	if err != nil {
		log.Println(err)
	}
	return ParseResults(res)[0]
}

func ExecuteQuery(body string) {
	_, err := connection.DB.Query(body)
	if err != nil {
		log.Println(err)
	}
}

func UpdateVersionById(id, version int) {
	_, err := connection.DB.Query(fmt.Sprintf(`
		UPDATE gomigrate_migrators 
		SET version=%d
		WHERE id=%d 
	`, version, id))

	if err != nil {
		log.Println(err)
	}
}

func SelectAll() []Migration {
	res, err := connection.DB.Query(fmt.Sprintf(`
		SELECT * FROM gomigrate_migrators`))

	if err != nil {
		log.Println(err)
	}
	return ParseResults(res)
}
