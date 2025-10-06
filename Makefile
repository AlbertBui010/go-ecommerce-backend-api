APP_NAME = server
PG_CONTAINER = shopdevgo-mysql

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