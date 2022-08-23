FROM golang:1.17.13-alpine3.15 as builder

WORKDIR /src

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server

FROM alpine:3.15

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/server .

ENTRYPOINT /app/server
