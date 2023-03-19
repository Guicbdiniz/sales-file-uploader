# syntax=docker/dockerfile:1

FROM golang:1.20.2-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /go-api ./cmd/web

CMD [ "/go-api" ]
