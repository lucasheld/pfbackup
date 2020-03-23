FROM golang:alpine AS builder

COPY . /go/src/github.com/lucasheld/pfbackup/
WORKDIR /go/src/github.com/lucasheld/pfbackup/

RUN apk --no-cache add make git
RUN GOOS=linux GOARCH=amd64 make build
RUN ./build/pfbackup --version


FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/lucasheld/pfbackup/build/pfbackup /usr/local/bin/

WORKDIR /data
VOLUME /data

ENTRYPOINT [ "pfbackup" ]
