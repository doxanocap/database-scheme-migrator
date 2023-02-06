package cmd

import (
	"fmt"
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CreateMigrations(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Printf("Invalid schema name \n")
		os.Exit(1)
	}

	schemaName := args[0]
	newMigrator := migrator.Insert(schemaName)

	up, down := services.GetFileName(schemaName, newMigrator.Id)
	fmt.Println(up, down)
	_, err := os.Create(migrator.SchemasPath + "\\" + up)
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	_, err = os.Create(migrator.SchemasPath + "\\" + down)
	if err != nil {
		log.Println("Error in the creating new schema file", err)
		os.Exit(1)
	}

	log.Println("Successfully created")
	connection.DB.Close()
}
