BINARY_NAME=netflux-backend

generate:
	@echo "Generating..."
	@go generate ./...
	@echo "Done"
test:
	@echo "Testing..."
	@go test ./...
	@echo "Done"
build: test
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v
	@echo "Done"
