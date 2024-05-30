.PHONY: db-up db-down run clean test test-db-setup test-db-teardown run-tests generate swag

db-up:
	docker-compose -f infrastructure/database/docker-compose.yml up -d

db-down:
	docker-compose -f infrastructure/database/docker-compose.yml down

run:
	go run cmd/app/main.go

clean:
	docker-compose -f infrastructure/database/docker-compose.yml down -v
	rm -rf infrastructure/database/data

test: test-db-setup run-tests test-db-teardown

test-db-setup:
	@echo "Starting test database..."
	docker-compose -f infrastructure/repository_impls/postgresql/testhelper/db/docker-compose.yml up -d
	@echo "Waiting for test database to be ready..."
	chmod +x infrastructure/repository_impls/postgresql/testhelper/db/wait-for-db.sh
	./infrastructure/repository_impls/postgresql/testhelper/db/wait-for-db.sh localhost 54320 echo "Test database is ready"

test-db-teardown:
	@echo "Stopping test database..."
	docker-compose -f infrastructure/repository_impls/postgresql/testhelper/db/docker-compose.yml down -v

run-tests:
	@echo "Running tests..."
	go test -v ./...

generate:
	go generate ./...

swag:
	swag init -g ./cmd/app/main.go