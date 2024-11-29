.PHONY: default run build test docs clean

APP_NAME=school-system

default: run

run:
	@go run cmd/main.go