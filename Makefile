APP_NAME = server
PG_CONTAINER = shopdevgo-mysql

GOOSE_DRIVER= mysql
GOOSE_DBSTRING= "root:root1234@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

run:
	go run ./cmd/${APP_NAME}

up:
	docker compose up -d

down:
	docker compose down

down-v:
	docker compose dowm -v

connect-mysql:
	docker exec -it ${PG_CONTAINER} mysql -u root -proot1234

goose-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

goose-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

goose-reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

