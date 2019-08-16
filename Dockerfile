FROM golang:1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go test -v ./...
RUN go build -o $GOPATH/bin/app

CMD ["app"]