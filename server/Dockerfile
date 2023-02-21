FROM golang:1.20.1-bullseye AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

ENV PORT 8000
EXPOSE ${PORT}

ENTRYPOINT [ "/usr/local/bin/app" ]
