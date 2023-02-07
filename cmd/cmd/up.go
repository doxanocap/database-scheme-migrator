package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"gomigrate/internal/migrator"
	"gomigrate/internal/services"
	"log"
	"os"
	"strings"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate up schemas",
	Long:  `Long: Migrate up schemas`,
	Run:   MigrateUpAllSchemas,
}

type F struct {
	rw *os.File
}

func init() {
	rootCmd.AddCommand(upCmd)
}

func MigrateUpAllSchemas(cmd *cobra.Command, args []string) {
	res := migrator.SelectAll()
	for _, m := range res {
		fn := services.GetFileName(m.Name, m.Id)
		rFile, wFile := F{}, F{}
		var parsingTypes = false
		var err error

		rFile.rw, err = os.Open(migrator.SchemasPath + "\\" + fn + "up.sql")
		if err != nil {
			log.Println(err)
		}

		wFile.rw, err = os.OpenFile(migrator.SchemasPath+"\\"+fn+"types.go", os.O_RDWR, 0777)
		if err != nil {
			log.Println(err)
		}

		wFile.Write("package schema \n\n")

		scanner := bufio.NewScanner(rFile.rw)
		scanner.Split(bufio.ScanLines)
		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		for i, line := range lines {
			arr := services.GetWordsFromLine(line)
			migrator.ExecuteQuery(line)

			if len(arr) == 1 && arr[0] == ")" {
				parsingTypes = false
				wFile.Write(fmt.Sprintf("} \n\n"))
				continue
			}

			if parsingTypes && len(arr) > 1 {
				temp := strings.ToLower(arr[1])

				wFile.Write(fmt.Sprintf("    %s", arr[0]))

				if len(temp) > 8 && temp == "timestamp" {
					wFile.Write(fmt.Sprintf("    int64\n"))
				}
				if temp == "serial" || temp == "int" || temp == "bigint" {
					wFile.Write(fmt.Sprintf("    int\n"))
				}

				if (len(temp) > 6 && temp[:7] == "varchar") || temp == "text" {
					wFile.Write(fmt.Sprintf("    string\n"))
				}

				if temp == "boolean" {
					wFile.Write(fmt.Sprintf("    bool\n"))
				}
			}

			if len(arr) > 1 && strings.ToLower(arr[0]) == "create" && strings.ToLower(arr[1]) == "table" {
				parsingTypes = true
				wFile.Write(fmt.Sprintf("type %s struct { \n", arr[2]))
			}
		}

		rFile.Close()
		wFile.Close()
		migrator.UpdateVersionById(m.Id, m.Version+1)
	}
}

func (f *F) Write(msg string) {
	_, err := f.rw.Write([]byte(msg))
	if err != nil {
		log.Println("Writing to file error -> ", err)
		os.Exit(0)
	}
}

func (f *F) Close() {
	err := f.rw.Close()
	if err != nil {
		log.Println(err)
	}
}
