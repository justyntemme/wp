FROM golang:latest as builder
WORKDIR /go/src/github.com/justyntemme/wp
COPY ./ ./

FROM builder as user-server-build
WORKDIR /go/src/github.com/justyntemme/wp/user
RUN go mod tidy
RUN go build -o user-server server/main.go

FROM golang:alpine as user-server-run
COPY --from=user-server-build /go/src/github.com/justyntemme/wp/user/user-server /go/bin/user-server
EXPOSE 8080
ENTRYPOINT [ "/go/bin/user-server" ]

FROM builder as club-server-build
WORKDIR /go/src/github.com/justyntemme/wp/club
RUN go mod tidy
RUN go build -o club-server server/main.go

FROM golang:alpine as club-server-run
COPY --from=club-server-build /go/src/github.com/justyntemme/wp/club/club-server /go/bin/club-server
EXPOSE 8080
ENTRYPOINT [ "/go/bin/club-server" ]

FROM builder as vote-server-build
WORKDIR /go/src/github.com/justyntemme/wp/vote
RUN go mod tidy
RUN go build -o vote-server server/main.go

FROM golang:alpine as vote-server-run
COPY --from=club-server-build /go/src/github.com/justyntemme/wp/vote/vote-server /go/bin/vote-server
EXPOSE 8080
ENTRYPOINT [ "/go/bin/club-server" ]