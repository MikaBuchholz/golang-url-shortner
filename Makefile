PORT=8080
DB_DOCKER_CONTAINER=url_shortner
BINARY_NAME=urlshortner
SUPER_USER=postgres
SUPER_PASSWORD=secret
DSN="host=localhost port=5432 user=${SUPER_USER} password=${SUPER_PASSWORD} dbname=${SUPER_USER} sslmode=disable timezone=UTC connect_timeout=5"

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${SUPER_USER} -e POSTGRES_PASSWORD=${SUPER_PASSWORD} -d postgres:12-alpine

stop-containers:
	@echo "Stop other docker containers"

start-docker:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init 

migrate-up:
	sqlx migrate run --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@localhost:5432/${SUPER_USER}?sslmode=disable"

migrate-down:
	sqlx migrate revert --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@localhost:5432/${SUPER_USER}?sslmode=disable"

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --user=${SUPER_USER} --owner=${SUPER_USER} urldb


stop_containers:
	@echo "Stopping other docker containers"
	@if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no active containers found..."; \
	fi

build:
	@echo "Building Server binary"
	go build -o ${BINARY_NAME} cmd/server/*.go
	@echo "Binary done building"

run: build stop_containers start-docker
	@echo "Starting api"
	@env PORT=${PORT} DSN=${DSN} ./${BINARY_NAME} &
	@echo "api started!"

stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backen