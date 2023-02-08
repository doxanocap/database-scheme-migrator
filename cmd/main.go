package main

import (
	"gomigrate/cmd/cmd"
	"gomigrate/internal/connection"
	"gomigrate/internal/services"
)

func main() {
	connection.InitConnection("postgres://postgres:eldoseldos@localhost:5432/gomigrate?sslmode=disable")
	services.InitMigrator()
	cmd.Execute()
}
