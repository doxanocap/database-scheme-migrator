package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gomigrate/internal/connection"
	"gomigrate/internal/migrator"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of all your migrations",
	Long:  `long: List of all your migrations`,
	Run:   ListAllMigrations,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func ListAllMigrations(cmd *cobra.Command, args []string) {
	res := migrator.SelectAll()
	fmt.Println("_____________________________________________________________________")
	for _, m := range res {
		fmt.Printf("%d  | Name: %s  -> Version: %d -> Created At: %d \n", m.Id, m.Name, m.Version, m.CreatedAt)
	}
	fmt.Println("_____________________________________________________________________")
	connection.DB.Close()
}
