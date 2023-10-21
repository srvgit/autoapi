.PHONY: generate-gqlgen run-generate go-mod-tidy all logs

go-mod-tidy:
	@echo "Tidying up go.mod and go.sum"
	@go mod tidy

generate-gqlgen:
	@echo "Generating GraphQL code with gqlgen"
	@go run github.com/99designs/gqlgen

run-generate:
	@echo "Running go generate"
	@go generate ./...

logs:
	@echo "Completed tidy, generate, and go generate."

all: generate-gqlgen run-generate go-mod-tidy logs
