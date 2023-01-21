FROM golang:latest as user-server-build
WORKDIR /go/src/github.com/justyntemme/wp
COPY ./user ./
RUN go mod download 
RUN go build -o user-server server/main.go
