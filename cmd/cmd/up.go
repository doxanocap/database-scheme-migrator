package cmd

import (
	"github.com/spf13/cobra"
	"gomigrate/internal/services"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate up schemas",
	Long:  `Long: Migrate up schemas`,
	Run:   MigrateUpAllSchemas,
}

func init() {
	rootCmd.AddCommand(upCmd)
}

func MigrateUpAllSchemas(cmd *cobra.Command, args []string) {
	services.MigrateUpAllSchemas()
}
