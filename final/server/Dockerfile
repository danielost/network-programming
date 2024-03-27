FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/pvms-final
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/pvms-final /go/src/pvms-final


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/pvms-final /usr/local/bin/pvms-final
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["pvms-final"]
