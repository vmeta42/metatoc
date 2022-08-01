FROM golang:1.17 AS build
COPY . /go/src/metatoc-client
WORKDIR /go/src/metatoc-client
ENV GOPROXY https://goproxy.cn,direct
RUN make metatoc-client.docker

FROM scratch
COPY --from=build /go/src/metatoc-client/metatoc-client.docker /opt/
ENTRYPOINT ["/opt/metatoc-client.docker"]