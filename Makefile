#!make
-include .env
export

run:
	go run cmd/main.go

up:
	docker-compose up -d

down:
	docker-compose down -v --remove-orphans

.PHONY: lint
lint: 
	test -z $$(gofmt -l . | grep -v vendor/) || (echo "Formatting issues found in:" $$(gofmt -l . | grep -v vendor/) && exit 1)

.PHONY: all
all: build

.PHONY: build-cli
build: dep## Build the cmd binary
	go build -trimpath -o bin/display cmd/main.go

.PHONY: dep
dep: ## Download app dependencies
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	rm -f ./bin/*
