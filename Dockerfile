FROM golang:latest as user-server-build

WORKDIR /go/src/github.com/justyntemme/wp

COPY . ./
WORKDIR /go/src/github.com/justyntemme/wp/user
RUN go mod tidy
RUN go build -o user-server server/main.go
FROM golang:alpine as user-server-run
COPY --from=user-server-build /go/src/github.com/justyntemme/wp/user/user-server /go/bin/user-server
EXPOSE 8080
ENTRYPOINT [ "/go/bin/user-erver" ]
