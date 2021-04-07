.PHONY: all
all: build
FORCE: ;

ALBUM_ENV ?= dev

.PHONY: build

clean:
	rm -rf bin/*

build: dependencies build-api

dependencies:
	go mod download

build-api:
	go build -tags $(ALBUM_ENV) -o ./bin/api api/main.go