FROM golang:1.19.3-alpine3.15 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY https://goproxy.cn,direct
RUN go install github.com/beego/bee/v2@latest
RUN go mod tidy
RUN bee update
RUN bee generate routers
CMD ["bee", "run"]
