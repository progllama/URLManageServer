# 2020/10/14最新versionを取得
FROM golang:1.15.2-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres
  # go get -u github.com/golang/mock/gomock && \
  # go get -u github.com/golang/mock/mockgen && \
  # go get -u github.com/google/wire/cmd/wire 

# ENV PATH $PATH:$HOME/go/bin

RUN go get -u github.com/oxequa/realize 
CMD ["realize", "start"]