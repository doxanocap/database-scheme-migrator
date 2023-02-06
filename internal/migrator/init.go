package migrator

import (
	"errors"
	"log"
	"os"
)

func InitMigrator() {
	osGotPath, err := os.Getwd()
	if err != nil {
		log.Printf("Init path error - %v", err)
		os.Exit(1)
	}

	WorkdirPath = osGotPath
	SchemasPath = WorkdirPath + "/schema"

	if _, err := os.Stat(SchemasPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(SchemasPath, 0777)
		if err != nil {
			log.Println(err)
		}
	}
}
