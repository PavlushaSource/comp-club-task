# syntax=docker/dockerfile:1

FROM golang:1.22.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /task ./cmd/app

ENTRYPOINT [ "/task" ]
