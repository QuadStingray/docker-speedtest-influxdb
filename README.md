# speedtest-influxdb:1.0.0

- [Introduction](#introduction)
    - [Contributing](#contributing)
    - [Issues](#issues)
- [Getting started](#getting-started)
    - [Installation](#installation)
    - [Quickstart](#quickstart)
    - [Environment Variables](#environment-variables)
    - [Grafana](#grafana)

# Introduction

Git-Repository to build [Docker](https://www.docker.com/) Container Image to run speedtest with [NDT7 Server](https://github.com/m-lab/ndt-server) from [mLabs](https://www.measurementlab.net/tests/ndt/ndt7/) to influxdb. The Implementation is inspired
by https://github.com/frdmn/docker-speedtest

## Contributing

If you find this image helpfull, so you can see here how you can help:

- Create an new branch and send a pull request with your features and bug fixes
- Help users resolve their [issues](https://github.com/QuadStingray/docker-speedtest-influxdb/issues).

## Issues

Before reporting your issue please try updating Docker to the latest version and check if it resolves the issue. Refer to the
Docker [installation guide](https://docs.docker.com/installation) for instructions.

If that recommendations do not help then [report your issue](https://github.com/QuadStingray/docker-speedtest-influxdb/issues/new) along with the following information:

- Output of the `docker version` and `docker info` commands
- The `docker run` command or `docker-compose.yml` used to start the image. Mask out the sensitive bits.

# Getting started

## Installation

Automated builds of the image are available on
[Dockerhub](https://hub.docker.com/r/quadstingray/speedtest-influxdb/)

```bash
docker pull speedtest-influxdb:1.0.0
```

Alternatively you can build the image yourself.

```bash
docker build . --tag 'speedtest-influxdb:dev';
```

## Quickstart

```bash
docker run -e "HOST=local" speedtest-influxdb:1.0.0
```

*Alternatively, you can use the sample [docker-compose.yml](docker-compose.yml) file to start the container using [Docker Compose](https://docs.docker.com/compose/)*

## Environment Variables

| Variable         | Default Value          | Informations                                                                                  |
|:-----------------|:-----------------------|:----------------------------------------------------------------------------------------------|
| INTERVAL         | 3600                   | Seconds between import of statistics                                                          |
| HOST             | local                  | host where the speedtest is running for grafana filter                                        |
| [SPEEDTEST_SERVER](#environment-variable-speedtest_server) | ''                     | ndt 7 server. Empty string, means speedtest return server for test                    |
| INCLUDE_READABLE_OUTPUT | false           | Log Speedtest Output to Console |
| SPEEDTEST_DISTANCE_UNIT | K               | Unit for Distance Calculation K = Kilometers, N = Nautical Miles other Values = Miles |
| SPEEDTEST_LIST_SERVERS | 'false'          | list all available ndt7 servers at the console                  |
| SPEEDTEST_LIST_KEEP_CONTAINER_RUNNING | 'true'          | keep docker container running after listing all ndt7 servers                  |
| SHOW_EXTERNAL_IP | 'false'          | You can activate logging your external Ip to InfluxDb to monitor IP changes.                   |
| INFLUXDB_USE     | 'true'   | You can deactivate save speedtest results to influx                                                             |
| INFLUXDB_URL     | http://influxdb:8086   | Url of your InfluxDb installation                                                             |
| INFLUXDB_DB      | speedtest              | Database at your InfluxDb installation                                                        |
| INFLUXDB_USER    | DEFAULT                | optional user for insert to your InfluxDb                                                     |
| INFLUXDB_PWD     | DEFAULT                | optional password for insert to your InfluxDb                                                 |

### Removed Variables

* SPEEDTEST_ALGO_TYPE

### Environment Variable: SPEEDTEST_SERVER

Per default the server is choosen automatically, but you can set `SPEEDTEST_SERVER` with the id of your favorite server. If your favorite Server doesn't answer a default search server
is choosen. You can get a list of all available servers by set the evironment variable `SPEEDTEST_LIST_SERVERS` to `true`. The list is ordered by country.

```
...
2021/02/02 09:16:09 County: AU | Location: Sydney | ServerId: syd03 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: AU | Location: Sydney | ServerId: syd02 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: BE | Location: Brussels | ServerId: bru01 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: BE | Location: Brussels | ServerId: bru03 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: BE | Location: Brussels | ServerId: bru05 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: BE | Location: Brussels | ServerId: bru04 | UplinkSpeed: 10g | Roundrobin: true
2021/02/02 09:16:09 County: BE | Location: Brussels | ServerId: bru02 | UplinkSpeed: 10g | Roundrobin: true

...
```

## Grafana

There is an sample grafana dashboard at this repository. You can import that to your Grafana installation. [speedtest.json](docker/grafana/provisioning/dashboards/speedtest.json)

![](https://raw.githubusercontent.com/QuadStingray/docker-speedtest-influxdb/master/images/speedtest_dashboard.png)

