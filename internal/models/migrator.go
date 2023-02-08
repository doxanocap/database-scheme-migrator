package models

import (
	"database/sql"
	"fmt"
	"gomigrate/internal/connection"
	"log"
	"time"
)

type Migration struct {
	Id        int    `json:"id"`
	Name      string `json:"Name"`
	Version   int    `json:"Version"`
	CreatedAt int64  `json:"CreatedAt"`
	ChangedAt int64  `json:"ChangedAt"`
}

type MigrationStash struct {
	Name         string `json:"Name"`
	Version      int    `json:"Version"`
	UpFileBody   string `json:"UpFileBody"`
	DownFileBody string `json:"DownFileBody"`
	ChangeTime   int64  `json:"ChangeTime"`
}

func ParseResults(res *sql.Rows) []Migration {
	var list []Migration
	for res.Next() {
		m := Migration{}
		err := res.Scan(&m.Id, &m.Name, &m.Version, &m.CreatedAt, &m.CreatedAt)
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
		(Name, CreatedAt, ChangedAt) 
		VALUES('%s', %d, %d)
		RETURNING *;`,
		name, time.Now().Unix(), time.Now().Unix()))

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

func UpdateMigratorById(id int) {
	_, err := connection.DB.Query(fmt.Sprintf(`
		UPDATE gomigrate_migrators 
		SET 
			version=version + 1,
			changedAt=%d

		WHERE id=%d 
	`, time.Now().Unix(), id))

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

func AddMigrationToStashById(m Migration, up, down string) {
	_, err := connection.DB.Query(fmt.Sprintf(`
		INSERT INTO gomigrate_migrator_stash
		(migrationid, name, version, upfilebody, downfilebody, changedat)
		VALUES(%d, '%s',%d, '%s', '%s', '%d')`, m.Id, m.Name, m.Version, up, down, time.Now().Unix()))

	if err != nil {
		log.Println(err)
	}
}

func SelectLatestMigraton() Migration {
	res, err := connection.DB.Query(fmt.Sprintf(`
		SELECT * FROM gomigrate_migrators
		ORDER BY changedat
	`))

	if err != nil {
		log.Println(err)
	}
	result := ParseResults(res)
	fmt.Println(result)
	if len(result) > 0 {
		return result[0]
	}
	return Migration{}
}
