deps:
	go mod download

verify:
	go mod verify

test:
	go test -v ./...

build-docker:
	docker build -t flowlab/ms .

run-local:
	docker run -it --rm flowlab/ms

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app