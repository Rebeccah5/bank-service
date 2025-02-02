# Set the database connection URL
DB_URL=postgres://bankuser:bankpass@localhost:5432/bankdb?sslmode=disable

# Set migration folder
MIGRATION_DIR=internal/migrations

# Build and run the service
run:
	go run cmd/main.go

# Generate a new migration file (Usage: make migration name=create_accounts)
migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

# Apply database migrations
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATION_DIR) up

# Rollback the last migration
migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATION_DIR) down 1

# Reset database by rolling back all migrations and reapplying them
migrate-reset:
	migrate -database "$(DB_URL)" -path $(MIGRATION_DIR) down
	migrate -database "$(DB_URL)" -path $(MIGRATION_DIR) up

# Install dependencies
deps:
	go mod tidy

# Build the project
build:
	go build -o bin/bank-service cmd/main.go

# Run tests
test:
	go test ./... -v

# Docker commands
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
