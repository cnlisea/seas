# Version: 1.0
# Author: lisea
# Date: 2026-09-10

# base image
FROM golang:1.18-alpine AS builder

# Env Arg
ARG ASSETSPATH=""
ARG GOPROXY="https://goproxy.cn"

# Env
ENV GOPROXY=$GOPROXY
ENV GO111MODULE=on

# go mod
WORKDIR /go/cache
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add files
WORKDIR /go/release
ADD . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" -o cmd/app cmd/main.go && \
    chmod +x tools/upx && \
    tools/upx cmd/app

# pack app and config
RUN if [ -f $ASSETSPATH ];then \
        tar zcvf /app.tar.gz cmd/app $ASSETSPATH; \
    else \
        tar zcvf /app.tar.gz cmd/app; \
    fi

# base image for two build
FROM alpine:latest

# Maintainer
LABEL maintainer.name="lisea" \
      maintainer.email="cnlisea@126.com"

# update apk mirror and timezone
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
        apk update && apk add ca-certificates tzdata && \
        ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
        echo "Asia/Shanghai" > /etc/timezone

# app build package for builder image
COPY --from=builder /app.tar.gz /

# unpack
RUN tar xf /app.tar.gz && mkdir /logs && rm -f /app.tar.gz

WORKDIR cmd

CMD ["./app"]
