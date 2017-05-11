FROM golang

ADD . /go/src/github.com/anistark/go-boilerplate
RUN go install github.com/anistark/go-boilerplate
ENTRYPOINT /go/bin/go-boilerplate

EXPOSE 8080
