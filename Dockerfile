FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go mod download
RUN go mod verify
RUN go test -v ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
