FROM golang:latest

WORKDIR /opt/app

ADD ./bin /opt/app/bin

ENTRYPOINT bin/hello-world