package cmd

import (
	"github.com/spf13/cobra"
	"gomigrate/internal/services"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate down all schemas",
	Long:  `Long: Migrate down all schemas`,
	Run:   MigrateDownAllSchemas,
}

func init() {
	rootCmd.AddCommand(downCmd)
}

func MigrateDownAllSchemas(cmd *cobra.Command, args []string) {
	services.MigrateDownAllSchemas()
}
