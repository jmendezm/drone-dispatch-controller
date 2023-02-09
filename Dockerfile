FROM golang:1.19-bullseye AS builder
WORKDIR /drone-dispatch-controller
COPY . .
RUN go mod vendor
RUN go build -ldflags "-linkmode external -extldflags -static"

#### deploy #####
FROM alpine:3.17
WORKDIR /opt/
COPY --from=builder /drone-dispatch-controller/drone-dispatch-controller /opt/drone-dispatch-controller
COPY ./config/*.json /var/config/
RUN chmod +x /opt/drone-dispatch-controller

ENV GIN_MODE=release \
    CONFIG_FILE=/var/config/config.json

EXPOSE 8080

CMD /opt/drone-dispatch-controller --config=$CONFIG_FILE
