package main

import (
	"gomigrate/cmd/cmd"
	"gomigrate/internal/connection"
	"gomigrate/internal/migrator"
)

func main() {
	connection.InitConnection("postgres://postgres:eldoseldos@localhost:5432/gomigrate?sslmode=disable")
	migrator.InitMigrator()
	cmd.Execute()
}
