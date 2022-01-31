FROM golang:1.17-alpine3.15 as builder

RUN apk update\
 && apk add --no-cache make\
 && apk add --no-cache ca-certificates\
 && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /query-analyzer-app /app/cmd/application/main.go

FROM alpine:latest
COPY --from=builder /query-analyzer-app /query-analyzer-app
EXPOSE 8080
ENTRYPOINT ["/query-analyzer-app"]
