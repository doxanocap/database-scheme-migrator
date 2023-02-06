package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gomigrate/internal/connection"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: InitGoMigrator,
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
