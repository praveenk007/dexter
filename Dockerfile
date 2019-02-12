FROM golang:1.8-onbuild
MAINTAINER kamathpraveen1992@gmail.com
ADD . /go/src/dexter
RUN go install dexter
ENTRYPOINT /go/bin/dexter

EXPOSE 8080