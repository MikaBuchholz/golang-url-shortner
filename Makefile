include .env

postgres: create-network
	docker run --name ${DB_DOCKER_CONTAINER} --network=${DOCKER_NETWORK_NAME} -p 5432:5432 -e POSTGRES_USER=${SUPER_USER} -e POSTGRES_PASSWORD=${SUPER_PASSWORD} -d postgres:12-alpine

stop-containers:
	@echo "Stop other docker containers"

start-docker:
	docker start ${DB_DOCKER_CONTAINER}

create-migrations:
	@if [ ! -d "migrations" ]; then \
		sqlx migrate add -r init;  \
	fi \
	
migrate-up:
	sqlx migrate run --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@${HOST}:5432/${SUPER_USER}?sslmode=disable"

migrate-down:
	sqlx migrate revert --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@${HOST}:5432/${SUPER_USER}?sslmode=disable"

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --user=${SUPER_USER} --owner=${SUPER_USER} urldb

build:
	@echo "Building Server binary"
	go build -o ${BINARY_NAME} cmd/server/*.go
	@echo "Binary done building"

run: build start-docker
	@echo "Starting api"
	./${BINARY_NAME}
	@echo "api started!"

stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backend"

create-network:
	docker network create --driver bridge ${DOCKER_NETWORK_NAME} || true

run-server-image: 
	docker run  --publish 8080:8080 --network=${DOCKER_NETWORK_NAME} shortner

build-server-image: 
	docker build --tag shortner .

bootstrap: 
	- make build-server-image 
	- make postgres 
	- make create-migrations 
	- make createdb 
	make migrate-up 
	- make run-server-image
	make mock-post

mock-post:
	curl -X POST \
		-H "Content-type: application/json" \
		-H "Accept: application/json" \
		-d '{"payload":"google.com"}' \
		"http://localhost:8080/api/v1/url/new"