package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gomigrate/internal/connection"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init",
	Long:  `Long: Init`,
	Run:   InitGoMigrator,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func InitGoMigrator(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Invalid postgres connection line: ")
		os.Exit(1)
	}

	connection.InitDatabaseSchemas()
}
