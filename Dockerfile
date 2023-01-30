# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

# copying all source files
COPY ./albums/ ./albums
COPY ./utils/ ./utils
COPY ./*.go ./
COPY ./.air.toml ./

RUN go build -o /app

CMD ["air", "-c", ".air.toml"]