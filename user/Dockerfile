FROM golang

RUN mkdir -p /go/src//Users/justyntemme/Documents/code/wp

ADD . /go/src//Users/justyntemme/Documents/code/wp

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /Users/justyntemme/Documents/code/wp/user/cmd  -watch /Users/justyntemme/Documents/code/wp/user
