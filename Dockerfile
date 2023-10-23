FROM golang:latest

WORKDIR /opt/app

ADD ./bin /opt/app/bin
ADD ./README.md /opt/app/README.md

ENTRYPOINT bin/hello-world