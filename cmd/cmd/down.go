package cmd

import (
	"gomigrate/internal/migrator"
	"gomigrate/internal/services"
	"log"
	"os"

	"github.com/spf13/cobra"
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
	res := migrator.SelectAll()
	for _, m := range res {
		_, down := services.GetFileName(m.Name, m.Id)
		content, err := os.ReadFile(migrator.SchemasPath + "\\" + down)
		if err != nil {
			log.Println(err)
		}
		migrator.ExecuteQuery(string(content))
		migrator.UpdateVersionById(m.Id, m.Version+1)
	}
}
