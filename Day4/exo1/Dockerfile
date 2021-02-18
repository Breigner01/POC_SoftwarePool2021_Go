FROM golang:alpine

WORKDIR /usr/app

ENV HELLOW_MESSAGE=docker

COPY . /usr/app

RUN go get github.com/gin-gonic/gin
RUN go build

CMD ./SoftwareGoDay2
EXPOSE 8080
