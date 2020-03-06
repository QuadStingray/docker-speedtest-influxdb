package model

import (
	"github.com/kylegrantlucas/speedtest/coords"
	"github.com/kylegrantlucas/speedtest/http"
)

type InfluxDbSettings struct {
	Use_Influx bool
	Db_Url     string
	Username   string
	Password   string
	Db         string
}

type Settings struct {
	Interval           int
	Host               string
	Server             string
	AlgoType           string
	ListServers        bool
	KeepProcessRunning bool
	ShowMyIp           bool
	InfluxDbSettings   InfluxDbSettings
}

type ClientInformations struct {
	ExternalIp string
	Provider   string
	Coordinate coords.Coordinate
}

type SpeedTestStatistics struct {
	Client   ClientInformations
	Server   http.Server
	Ping     float64
	Down_Mbs float64
	Up_Mbs   float64
}
