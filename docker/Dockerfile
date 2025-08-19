# Stage 1: Base image and downloading dependencies
FROM golang:1.21 AS base
MAINTAINER Van Le <vanlt@phx-smartschool.com>

WORKDIR /go/src/backend

COPY go.mod .
COPY go.sum .

RUN go mod download

# Stage 2: Build the go app
FROM base AS builder

COPY . /go/src/backend
WORKDIR /go/src/backend/cmd/http

# Sử dụng ARG để cấu hình GOPROXY
ARG GOPROXY=https://goproxy.cn

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.4 && \
    swag init && \
    mkdir -p bin && \
    GOPROXY=$GOPROXY go build -o ./bin/ ./...

# Stage 3: Create the final image
FROM debian:stable-slim

MAINTAINER Van Le <vanlt@phx-smartschool.com>

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates \
        netbase && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /go/src/backend/cmd/http/bin ./app
COPY --from=builder /go/src/backend/configs ./app/configs

WORKDIR ./app
ENTRYPOINT ["./http"]