#!/bin/bash

./speedtestInfluxDB -interval="$INTERVAL" -host="$HOST" -server="$SPEEDTEST_SERVER" -influxHost="$INFLUXDB_URL" -influxDB="$INFLUXDB_DB" -influxUser="$INFLUXDB_USER" -influxPwd="$INFLUXDB_PWD"