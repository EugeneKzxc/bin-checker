.PHONY: all up db_init wait_for_db_init server

all: up wait_for_db_init server

up:
	docker-compose up -d postgres

wait_for_db_init:
	@echo "Waiting for postgres container to be healthy..."
	@until [ "`docker inspect -f {{.State.Health.Status}} pg_container`" == "healthy" ]; do \
		sleep 1; \
	done
	@echo "Initializing the database..."
	go run ./utils/CSVparser/main.go
	@echo "Database initialization completed."

server:
	docker-compose up -d go-server
