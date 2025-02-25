APP_NAME=myapp

air: 
	@echo "Live reloading..."
	@air -c .air.toml

build: 
	@echo "Building..."
	@go build -o bin/$(APP_NAME) cmd/api/main.go

run:
	@echo "Running..."
	@go build -o bin/$(APP_NAME) cmd/api/main.go
	@./bin/$(APP_NAME)

up:
	@echo "Up..."
	@go run cmd/migration/main.go up

down:
	@echo "Down..."
	@go run cmd/migration/main.go down

drop:
	@echo "Drop..."
	@go run cmd/drop/main.go

migrate:
	@echo "Seeding database..."
	@go run cmd/migrate/main.go
	
seed:
	@echo "Seeding database..."
	@go run cmd/seed/main.go