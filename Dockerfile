FROM golang:latest

RUN mkdir /go/src/github.com/
RUN mkdir /go/src/github.com/YanaPIIDXer
RUN mkdir /go/src/github.com/YanaPIIDXer/QiitaLineBot
WORKDIR /go/src/github.com/YanaPIIDXer/QiitaLineBot
RUN go get github.com/gin-gonic/gin
ADD . /go/src/github.com/YanaPIIDXer/QiitaLineBot
