FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/pvms-client
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/pvms-client /go/src/pvms-client


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/pvms-client /usr/local/bin/pvms-client
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["pvms-client"]
