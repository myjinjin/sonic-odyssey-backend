.PHONY: dbup dbdown run clean

dbup:
	docker-compose -f infrastructure/database/docker-compose.yml up -d

dbdown:
	docker-compose -f infrastructure/database/docker-compose.yml down

run:
	go run cmd/app/main.go

clean:
	docker-compose -f infrastructure/database/docker-compose.yml down -v
	rm -rf infrastructure/database/data