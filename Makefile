SHELL := /bin/bash

deps:
	go mod download

deps-reset:
	git checkout -- go.mod

tidy:
	go mod tidy

verify:
	go mod verify

test:
	go test -v ./...

build-docker:
	docker build -t flowlab/ms .

build-app:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

run-local:
	docker run -it --rm flowlab/ms