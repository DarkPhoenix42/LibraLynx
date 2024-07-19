include .env

DB_URL := mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

DEFAULT: build run

migration_up: 
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose up

migration_down: 
	yes | migrate -path db/migrations/ -database "$(DB_URL)" -verbose down

migration_restart:
	make migration_down
	make migration_up

build:
	go mod tidy
	go build -o bin/libralynx cmd/main.go
	
run:
	./bin/libralynx