.PHONY: db-up db-down run clean testdb testdb-up testdb-down test

db-up:
	docker-compose -f infrastructure/database/docker-compose.yml up -d

db-down:
	docker-compose -f infrastructure/database/docker-compose.yml down

run:
	go run cmd/app/main.go

clean:
	docker-compose -f infrastructure/database/docker-compose.yml down -v
	rm -rf infrastructure/database/data

testdb: testdb-up test testdb-down

testdb-up:
	@echo "Starting test database..."
	docker-compose -f infrastructure/repository_impls/postgresql/testhelper/db/docker-compose.yml up -d
	@echo "Waiting for test database to be ready..."
	sleep 5

testdb-down:
	@echo "Stopping test database..."
	docker-compose -f infrastructure/repository_impls/postgresql/testhelper/db/docker-compose.yml down

test:
	@echo "Running tests..."
	go test -v ./infrastructure/repository_impls/postgresql/tests/...