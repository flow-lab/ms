FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev make

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN make test
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -ldflags="-w -s -X 'main.version=$(git describe --tags || echo "unknown")' -X 'main.commit=$(git rev-parse HEAD)' -X 'main.date=$(date)'" -o /go/bin/app ./cmd/srv/

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]