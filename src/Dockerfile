# ビルド用のDockerファイルです。
FROM golang:alpine
COPY ./src /go/src/app
WORKDIR /go/src/app
RUN apk update && apk add git && apk add gcc && apk add --no-cache musl-dev
RUN go mod tidy
