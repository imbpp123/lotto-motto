#!make
-include .env
export

.PHONY: lint
lint: 
	test -z $$(gofmt -l . | grep -v vendor/) || (echo "Formatting issues found in:" $$(gofmt -l . | grep -v vendor/) && exit 1)

.PHONY: all
all: clean build

.PHONY: build-cli
build: clean dep
	go build -trimpath -o bin/lotto_6aus49 cmd/lotto6aus49/main.go
	go build -trimpath -o bin/lotto_eurojackpot cmd/lottoEuroJackpot/main.go
	chmod +x bin/lotto_6aus49
	chmod +x bin/lotto_eurojackpot

.PHONY: dep
dep: 
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	rm -f ./bin/*

euro:
	go run cmd/eurojackpot/main.go -rows=5 file=https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip

6aus49:
	go run cmd/6aus49/main.go -rows=5 file=https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_lotto.zip
