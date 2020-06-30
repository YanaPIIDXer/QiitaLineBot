FROM golang:latest

RUN mkdir /go/src/github.com/
RUN mkdir /go/src/github.com/YanaPIIDXer
RUN mkdir /go/src/github.com/YanaPIIDXer/QiitaLineBot
WORKDIR /go/src/github.com/YanaPIIDXer/QiitaLineBot
ADD . /go/src/github.com/YanaPIIDXer/QiitaLineBot
