.PHONY: default run build test docs clean go_with_swagger 

#Variables
APP_NAME=gopportunities-api

#tasks
default: go_with_swagger 

run: 
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
go_with_swagger:
	@swag init
	@go run main.go