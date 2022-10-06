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
	go build -trimpath -o bin/display_eurojackpot cmd/eurojackpot/main.go
	chmod +x bin/display_eurojackpot

.PHONY: dep
dep: ## Download app dependencies
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	rm -f ./bin/*

euro: 
	docker-compose run app bin/display_eurojackpot 10 'https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip'