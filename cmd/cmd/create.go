package cmd

import (
	"github.com/spf13/cobra"
	"gomigrate/internal/connection"
	"gomigrate/internal/migrator"
	"gomigrate/internal/services"
	"log"
	"os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creating migrations",
	Long:  `Long: Creating migrations`,
	Run:   CreateMigrations,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func CreateMigrations(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Printf("Invalid schema name \n")
		os.Exit(1)
	}

	schemaName := args[0]
	newMigrator := migrator.Insert(schemaName)

	fn := services.GetFileName(schemaName, newMigrator.Id)
	_, err := os.Create(migrator.SchemasPath + "\\" + fn + "up.sql")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	_, err = os.Create(migrator.SchemasPath + "\\" + fn + "down.sql")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	_, err = os.Create(migrator.SchemasPath + "\\" + fn + "types.go")
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	log.Println("Successfully created")
	connection.DB.Close()
}
