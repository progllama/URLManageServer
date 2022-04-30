# # 2020/10/14最新versionを取得
# FROM golang:alpine
# # アップデートとgitのインストール！！
# RUN apk update && apk add git && apk add gcc && apk add --no-cache musl-dev
# # appディレクトリの作成
# RUN mkdir /go/src/app
# # ワーキングディレクトリの設定
# WORKDIR /go/src/app
# # ホストのファイルをコンテナの作業ディレクトリに移行
# ADD . /go/src/app

# RUN go get -u github.com/gin-gonic/gin && \
#   go get github.com/jinzhu/gorm && \
#   go get github.com/jinzhu/gorm/dialects/postgres

# RUN go get -u github.com/oxequa/realize 
# CMD ["realize", "start"]

FROM golang:alpine AS Builder
COPY ./src /go/src/url_manager
WORKDIR /go/src/url_manager
RUN go build ./main.go

FROM alpine:latest AS Product
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=Builder /go/src/url_manager/main ./
COPY --from=Builder /go/src/url_manager/favicon.ico ./favicon.ico
COPY --from=Builder /go/src/url_manager/app/templates ./app/templates
COPY --from=Builder /go/src/url_manager/app/assets ./app/assets