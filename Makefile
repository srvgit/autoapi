.PHONY: generate-gqlgen run-generate

generate-gqlgen:
	go run github.com/99designs/gqlgen

run-generate:
	go generate ./...

all: generate-gqlgen run-generate
