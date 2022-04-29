# 2020/10/14最新versionを取得
FROM golang:alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git && apk add gcc && apk add --no-cache musl-dev
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres

RUN go get -u github.com/oxequa/realize 
CMD ["realize", "start"]