#!/bin/bash

./speedtestInfluxDB -saveToInfluxDb="$INFLUXDB_USE" -interval="$INTERVAL" -host="$HOST" -server="$SPEEDTEST_SERVER" -influxHost="$INFLUXDB_URL" -influxDB="$INFLUXDB_DB" -influxUser="$INFLUXDB_USER" -influxPwd="$INFLUXDB_PWD" -list="$SPEEDTEST_LIST_SERVERS" -showExternalIp="$SHOW_EXTERNAL_IP"