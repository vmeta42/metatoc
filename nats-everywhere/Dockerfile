FROM golang:1.17 AS build
COPY . /go/src/nats-everywhere
WORKDIR /go/src/nats-everywhere
ENV GOPROXY https://goproxy.cn,direct
RUN make nats-everywhere.docker

FROM scratch
ARG NATS_URL
ARG NATS_SUBJECT
ARG NATS_DURABLE
ENV NATS_URL ${NATS_URL}
ENV NATS_SUBJECT ${NATS_SUBJECT}
ENV NATS_DURABLE ${NATS_DURABLE}
COPY --from=build /go/src/nats-everywhere/nats-everywhere.docker /opt/
COPY --from=build /go/src/nats-everywhere/logs /opt/
ENTRYPOINT ["/opt/nats-everywhere.docker"]