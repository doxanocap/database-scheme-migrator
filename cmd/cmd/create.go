package cmd

import (
	"github.com/spf13/cobra"
	"gomigrate/internal/services"
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
	services.CreateMigrations(args)
}
