FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod tidy