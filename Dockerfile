FROM golang:bullseye
COPY ./src /go/src/app
COPY ./env/.env /go/src/app
WORKDIR /go/src/app
RUN go build cmd/main.go
ENTRYPOINT ./main