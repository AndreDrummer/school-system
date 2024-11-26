.PHONY: default run build test docs clean

APP_NAME=school-system

default: run

server:
	@go run cmd/server/main.go

app:
	@go run cmd/app/main.go