package model

type InfluxDbSettings struct {
	dbUrl    string
	username string
	password string
	db       string
}

type Settings struct {
	Interval         int
	Host             string
	Server           string
	ListServers      bool
	InfluxDbSettings InfluxDbSettings
}

type SpeedTestStatistics struct {
	ServerId string
	Location string
	Ping     float64
	Down_Mbs float64
	Up_Mbs   float64
	Distance float64
}
