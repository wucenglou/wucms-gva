FROM golang:alpine as builder

WORKDIR /go/src/github.com/wucenglou/wucms-gva/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="wucenglou@qq.com"

WORKDIR /go/src/github.com/wucenglou/wucms-gva/server

COPY --from=0 /go/src/github.com/wucenglou/wucms-gva/server/server ./
COPY --from=0 /go/src/github.com/wucenglou/wucms-gva/server/resource ./resource/
COPY --from=0 /go/src/github.com/wucenglou/wucms-gva/server/config.docker.yaml ./
COPY --from=0 /go/src/github.com/wucenglou/wucms-gva/server/uploads ./uploads/

EXPOSE 8828
ENTRYPOINT ./server -c config.docker.yaml