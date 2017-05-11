FROM golang

ADD . /go/src/github.com/anistark/magnolia
RUN go install github.com/anistark/magnolia
ENTRYPOINT /go/bin/magnolia

EXPOSE 8080
