FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/KKitsun/usdc-tracker-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/usdc-tracker-svc /go/src/github.com/KKitsun/usdc-tracker-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/usdc-tracker-svc /usr/local/bin/usdc-tracker-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["usdc-tracker-svc"]
