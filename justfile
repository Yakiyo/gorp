alias b := build
alias r := run

default:
	just --list

run:
	go run .

build:
	go build -o gorp.exe .

test:
	go test ./...

fmt:
	go fmt ./...

update:
	go get -u
	go mod tidy