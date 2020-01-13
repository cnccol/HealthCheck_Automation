FROM golang:latest
LABEL maintainer="Juan Sosa <jsosa@cnccol.com>"

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN apt-get update && apt-get install -y net-tools vim

WORKDIR /app