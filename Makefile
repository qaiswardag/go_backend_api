APP_NAME=myapp

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

controller:
	@echo "Creating controller..."
	@mkdir -p internal/controller/$(folder)
	@echo -e "/*\n| Controller Method Naming Convention\n| Controller methods: index, create, store, show, edit, update, destroy.\n| Please aim for consistency by using these method names in all controllers.\n*/" > internal/controller/$(folder)/$(name).go
	@echo "package $(folder)" >> internal/controller/$(folder)/$(name).go
	@echo "func Index() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Create() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Store() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Show() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Edit() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Update() {}" >> internal/controller/$(folder)/$(name).go
	@echo "func Destroy() {}" >> internal/controller/$(folder)/$(name).go
