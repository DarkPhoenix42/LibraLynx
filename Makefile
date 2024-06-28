include .env

DB_URL := mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)
migration_up: 
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose up
migration_down: 
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose down
migration_fix:
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose force 1