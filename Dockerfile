# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

# copying all source files
COPY ./albums/ ./albums
COPY ./utils/ ./utils
COPY ./*.go ./

RUN go build -o /app

CMD [ "/app" ]