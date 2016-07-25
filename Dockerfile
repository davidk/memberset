# Run lint on anything in current directory
# docker build -t lint .
# docker run lint -v ./:/go/src/github.com/davidk/memberset
FROM golang:alpine
MAINTAINER davidk <kdavid+build@gmail.com>
RUN apk update
RUN apk add git
RUN go get -u github.com/golang/lint/golint
