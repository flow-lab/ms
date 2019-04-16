FROM golang:1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev

WORKDIR /go/src/app
COPY . .

RUN go get -d -t -v ./...
RUN go test -v ./...
RUN go install -v ./...

CMD ["app"]