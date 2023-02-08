package services

import (
	"errors"
	"fmt"
	"gomigrate/internal/connection"
	"gomigrate/internal/models"
	"log"
	"os"
)

func InitMigrator() {
	osGotPath, err := os.Getwd()
	if err != nil {
		log.Printf("Init path error - %v", err)
		os.Exit(1)
	}

	models.WorkdirPath = osGotPath
	models.SchemasPath = models.WorkdirPath + "/schema"

	if _, err := os.Stat(models.SchemasPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(models.SchemasPath, 0777)
		if err != nil {
			log.Println(err)
		}
	}
}

func MigrateDownAllSchemas() {
	res := models.SelectAll()
	for _, m := range res {
		fn := GetFileName(m.Name, m.Id)
		content, err := os.ReadFile(models.SchemasPath + "\\" + fn + "down.sql")
		if err != nil {
			log.Println(err)
		}
		models.ExecuteQuery(string(content))
	}
}

func MigrateUpAllSchemas() {
	res := models.SelectAll()
	for _, m := range res {
		fn := GetFileName(m.Name, m.Id)
		rFile, wFile := F{}, F{}
		var err error

		rFile.rw, err = os.Open(models.SchemasPath + "\\" + fn + "up.sql")
		if err != nil {
			log.Println(err)
		}

		wFile.rw, err = os.OpenFile(models.SchemasPath+"\\"+fn+"types.go", os.O_RDWR, 0777)
		if err != nil {
			log.Println(err)
		}

		content1 := ParseStructFromSchema(wFile, rFile)
		content2, err := os.ReadFile(models.SchemasPath + "\\" + fn + "down.sql")
		if err != nil {
			log.Fatal(err)
		}

		models.AddMigrationToStashById(m, content1, string(content2))
		models.ExecuteQuery(content1)
		models.UpdateMigratorById(m.Id)
		wFile.Close()
		rFile.Close()
	}
}

func CreateMigrations(args []string) {
	if len(args) == 0 {
		log.Printf("Invalid schema name \n")
		os.Exit(1)
	}

	schemaName := args[0]
	newMigrator := models.Insert(schemaName)

	fn := GetFileName(schemaName, newMigrator.Id)
	_, err := os.Create(models.SchemasPath + "\\" + fn + "up.sql")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	_, err = os.Create(models.SchemasPath + "\\" + fn + "down.sql")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	_, err = os.Create(models.SchemasPath + "\\" + fn + "types.go")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}
	log.Println("Successfully created")
	connection.DB.Close()
}

func ExecuteLatestMigration() {
	res := models.SelectLatestMigraton()
	fmt.Println(res)
}
