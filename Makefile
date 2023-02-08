init:
	go run cmd/main.go init postgres://postgres:eldoseldos@localhost:5432/gomigrate
create:
	go run cmd/main.go create test
list:
	go run cmd/main.go list
up:
	go run cmd/main.go up
down:
	go run cmd/main.go down
dbversion:
	go run cmd/main.go dbversion