FROM golang:alpine AS build-env

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh

# build directories
ADD . /go/src/quadstingray/speedtest-influxdb
WORKDIR /go/src/quadstingray/speedtest-influxdb

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
RUN go build -o speedtestInfluxDB *.go

# final stage
FROM alpine
WORKDIR /app

MAINTAINER QuadStingray <docker-speedtest@quadstingray.com>

ENV INTERVAL=3600 \
    INFLUXDB_USE="true" \
    INFLUXDB_DB="speedtest" \
    INFLUXDB_URL="http://influxdb:8086" \
    INFLUXDB_USER="DEFAULT" \
    INFLUXDB_PWD="DEFAULT" \
    HOST="local" \
    SPEEDTEST_SERVER="" \
    SPEEDTEST_LIST_SERVERS="false" \
    SHOW_EXTERNAL_IP="false"

RUN apk add ca-certificates
COPY --from=build-env /go/src/quadstingray/speedtest-influxdb/speedtestInfluxDB /app/speedtestInfluxDB
ADD run.sh run.sh
CMD sh run.sh