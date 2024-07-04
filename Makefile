include .env

DB_URL := mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

DEFAULT: run

migration_up: 
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose up

migration_down: 
	yes | migrate -path db/migrations/ -database "$(DB_URL)" -verbose down

migration_restart:
	make migration_down
	make migration_up
	
run:
	make migration_restart
	go run main.go