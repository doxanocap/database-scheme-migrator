package services

import "strconv"

func GetFileName(name string, id int) (up, down string) {
	strId := strconv.Itoa(id)
	filename := ""
	for i := 0; i < 5-len(strId); i++ {
		filename += "0"
	}
	filename += strId + "_" + name + "_"
	up = filename + "up.sql"
	down = filename + "down.sql"
	return up, down
}
