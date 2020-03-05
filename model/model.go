package model

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
	ListServers        bool
	KeepProcessRunning bool
	ShowMyIp           bool
	InfluxDbSettings   InfluxDbSettings
}

type SpeedTestStatistics struct {
	ServerId    string
	Location    string
	Ping        float64
	Down_Mbs    float64
	Up_Mbs      float64
	Distance    float64
	External_Ip string
}
