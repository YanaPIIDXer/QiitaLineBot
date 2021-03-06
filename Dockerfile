FROM golang:latest as builder

# ディレクトリ生成
RUN mkdir /go/src/github.com/
RUN mkdir /go/src/github.com/YanaPIIDXer
RUN mkdir /go/src/github.com/YanaPIIDXer/QiitaLineBot
RUN mkdir /go/src/github.com/YanaPIIDXer/QiitaLineBot/bin
WORKDIR /go/src/github.com/YanaPIIDXer/QiitaLineBot

# 各種モジュールのインストール
RUN go get -v github.com/gin-gonic/gin
RUN go get -v -u github.com/oxequa/realize
RUN go get -v github.com/line/line-bot-sdk-go/linebot
RUN go get -v github.com/joho/godotenv

ADD . /go/src/github.com/YanaPIIDXer/QiitaLineBot

ENV PORT=${PORT}

# ビルド
RUN go build -o bin/server.out .

#ENTRYPOINT ["bin/server.out"]
CMD bin/server.out
