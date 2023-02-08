package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type F struct {
	rw *os.File
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

func GetFileName(name string, id int) string {
	strId := strconv.Itoa(id)
	filename := ""
	for i := 0; i < 5-len(strId); i++ {
		filename += "0"
	}
	filename += strId + "_" + name + "_"
	return filename
}

func GetWordsFromLine(line string) []string {
	temp := strings.Split(line, " ")
	words := []string{}
	for i, word := range temp {
		if word != "" {
			if i == len(temp)-1 {
				word = word[:len(word)-1]
			}
			words = append(words, word)
		}
	}
	return words
}

func ParseStructFromSchema(wFile, rFile F) string {
	var parsingTypes = false
	var lines string
	wFile.Write("package schema \n\n")

	scanner := bufio.NewScanner(rFile.rw)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lines += line + "\n"

		arr := GetWordsFromLine(line)

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
	return lines
}
