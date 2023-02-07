package services

import (
	"strconv"
	"strings"
)

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
