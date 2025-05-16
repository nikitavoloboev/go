default:
	@just --list

build:
	go build -o dist/cli-go

run:
	./dist/cli-go
