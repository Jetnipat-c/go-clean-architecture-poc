FROM golang:1.22.1-alpine

ARG GITHUB_TOKEN

WORKDIR /app

RUN apk add git mercurial

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN apk add git mercurial

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go mod download

CMD ["air", "-c", ".air.toml"]