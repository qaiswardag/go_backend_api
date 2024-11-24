run:
	@echo "Running..."
	@go build -o bin/$(APP_NAME) cmd/api/$(APP_NAME)/main.go

build: 
	@echo "Running..."
	@go run cmd/api/$(APP_NAME)/main.go