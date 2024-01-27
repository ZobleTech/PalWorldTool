FROM golang:alpine AS build

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

ADD . /src

RUN cd /src && go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=build /src/main /app/

EXPOSE 8080

ENTRYPOINT ./main