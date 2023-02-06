package cmd

import (
	"github.com/spf13/cobra"
	"gomigrate/internal/migrator"
	"gomigrate/internal/services"
	"log"
	"os"
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
	res := migrator.SelectAll()
	for _, m := range res {
		up, _ := services.GetFileName(m.Name, m.Id)
		content, err := os.ReadFile(migrator.SchemasPath + "\\" + up)
		if err != nil {
			log.Println(err)
		}
		migrator.ExecuteQuery(string(content))
		migrator.UpdateVersionById(m.Id, m.Version+1)
	}
}
