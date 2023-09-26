include .env

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${SUPER_USER} -e POSTGRES_PASSWORD=${SUPER_PASSWORD} -d postgres:12-alpine

stop-containers:
	@echo "Stop other docker containers"

start-docker:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init 

migrate-up:
	sqlx migrate run --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@${HOST}:5432/${SUPER_USER}?sslmode=disable"

migrate-down:
	sqlx migrate revert --database-url "postgres://${SUPER_USER}:${SUPER_PASSWORD}@${HOST}:5432/${SUPER_USER}?sslmode=disable"

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
	./${BINARY_NAME}
	@echo "api started!"

stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backen

# Not yet ready for use - Dockerfile not completed 
run-server-image:
	docker run shortner

build-server-image:
	docker build --tag shortner .

start-project: run
	cd frontend && npm i && npm run dev
#