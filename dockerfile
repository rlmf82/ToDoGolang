# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

ADD . /app

RUN go mod download

RUN go env -w GO111MODULE=on

RUN go build -o /toDoGolang

CMD [ "/toDoGolang" ]