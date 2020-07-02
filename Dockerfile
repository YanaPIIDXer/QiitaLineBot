FROM golang:latest

RUN mkdir /go/src/github.com/
RUN mkdir /go/src/github.com/YanaPIIDXer
RUN mkdir /go/src/github.com/YanaPIIDXer/QiitaLineBot
WORKDIR /go/src/github.com/YanaPIIDXer/QiitaLineBot
RUN go get -v github.com/gin-gonic/gin
RUN go get -v -u github.com/oxequa/realize
RUN go get -v github.com/line/line-bot-sdk-go/linebot
ADD . /go/src/github.com/YanaPIIDXer/QiitaLineBot
RUN go build -o server.out .
