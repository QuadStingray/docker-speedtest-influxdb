package model

import (
	"github.com/kylegrantlucas/speedtest/coords"
)

type InfluxDbSettings struct {
	Use_Influx bool
	Db_Url     string
	Username   string
	Password   string
	Db         string
}

type Settings struct {
	Interval             int
	Host                 string
	Server               string
	DistanceUnit         string
	ListServers          bool
	KeepProcessRunning   bool
	ShowMyIp             bool
	IncludeHumanReadable bool
	InfluxDbSettings     InfluxDbSettings
}

type ClientInformations struct {
	ExternalIp string
	Provider   string
	Coordinate coords.Coordinate
}

type SpeedTestStatistics struct {
	Client             ClientInformations
	Server             Server
	Ping               float64
	Down_Mbs           float64
	Up_Mbs             float64
	DownRetransPercent float64
}

type Server struct {
	URL      string
	Lat      float64
	Lon      float64
	Name     string
	Country  string
	City     string
	Distance float64
	Latency  float64
}
