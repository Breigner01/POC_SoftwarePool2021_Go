FROM golang:alpine

WORKDIR /go/src/app

ENV HELLOW_MESSAGE=docker

COPY . /go/src/app

RUN go get github.com/gin-gonic/gin
RUN go build

CMD ./SoftwareGoDay2
EXPOSE 8080
