FROM golang:1.15-alpine3.13 AS build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories;\
    apk add --no-cache make git bash
COPY . /go/src/github.com/inspursoft/cefco
WORKDIR /go/src/github.com/inspursoft/cefco
RUN go env -w GOPROXY=https://goproxy.cn,direct
ARG VERSION 
ARG GIT_INFO
RUN VERSION=$VERSION gitInfo=$GIT_INFO make filesync-controller noRace=T


FROM alpine:3.13

COPY --from=build /go/src/github.com/inspursoft/cefco/filesync-controller /filesync-controller

CMD ["/filesync-controller"]
