#!/bin/bash

./speedtestInfluxDB -interval="$INTERVAL"  \
                    -saveToInfluxDb="$INFLUXDB_USE" \
                    -influxHost="$INFLUXDB_URL"  \
                    -influxDB="$INFLUXDB_DB"  \
                    -influxUser="$INFLUXDB_USER"  \
                    -influxPwd="$INFLUXDB_PWD"  \
                    -host="$HOST"  \
                    -server="$SPEEDTEST_SERVER"  \
                    -list="$SPEEDTEST_LIST_SERVERS"  \
                    -keepProcessRunning="$SPEEDTEST_LIST_KEEP_CONTAINER_RUNNING"  \
                    -distanceUnit="$SPEEDTEST_DISTANCE_UNIT"  \
                    -includeHumanOutput="$INCLUDE_READABLE_OUTPUT"  \
                    -showExternalIp="$SHOW_EXTERNAL_IP"