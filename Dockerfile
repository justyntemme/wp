FROM golang:latest as user-server-build

WORKDIR /go/src/github.com/justyntemme/wp

COPY . ./
WORKDIR /go/src/github.com/justyntemme/wp/user
RUN go mod tidy
RUN go build -o user-server server/main.go
