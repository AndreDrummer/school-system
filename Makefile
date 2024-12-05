.PHONY: default run build test docs clean

APP_NAME=school-system

default: run-local

run-local:
	@go run cmd/main.go

run:
	docker-compose up -d; go run cmd/main.go