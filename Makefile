.DEFAULT_GOAL := run

run: run-server
build: build-server

build-server:
	docker-compose build server

run-server: build-server
	docker-compose up -d server
	docker-compose logs -f server

down:
	docker-compose down

.PHONY: build build-server run run-server
