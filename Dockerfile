FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev make

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN make test verify build-app

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]