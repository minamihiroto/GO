FROM golang

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app