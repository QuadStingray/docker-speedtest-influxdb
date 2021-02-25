package model

import (
	"flag"
	"log"
)

func Parser() Settings {
	var interval int
	var retryInterval int

	var server string
	var host string
	var influxHost string
	var influxDB string
	var influxPwd string
	var influxUser string

	var list bool
	var keepProcessRunning bool
	var showExternalIp bool
	var saveToInfluxDb bool
	var distanceUnit string
	var includeHumanOutput bool
	var retryZeroValue bool

	flag.IntVar(&interval, "interval", 3600, "seconds between statistics import")
	flag.IntVar(&retryInterval, "retryInterval", 3600, "seconds between statistics retry")

	flag.StringVar(&host, "host", "", "host where the speedetest is running")
	flag.StringVar(&influxHost, "influxHost", "http://influxdb:8086", "host of your influxdb instance")
	flag.StringVar(&influxDB, "influxDB", "speetest", "influxdb database")
	flag.StringVar(&influxUser, "influxUser", "DEFAULT", "influxdb Username")
	flag.StringVar(&influxPwd, "influxPwd", "DEFAULT", "influxdb Password")
	flag.StringVar(&distanceUnit, "distanceUnit", "K", "Distance Unit between GeoPoints possible Values K|M|N")

	flag.BoolVar(&includeHumanOutput, "includeHumanOutput", true, "Log HumanReadableOutput to Console")
	flag.BoolVar(&saveToInfluxDb, "saveToInfluxDb", false, "save to influxdb")
	flag.BoolVar(&keepProcessRunning, "keepProcessRunning", false, "keep process running")
	flag.BoolVar(&showExternalIp, "showExternalIp", true, "save and show external Ip of docker host")
	flag.BoolVar(&retryZeroValue, "retryZeroValue", false, "retry speedtest at zero values returned")

	flag.StringVar(&server, "server", "", "ndt7 server")
	flag.BoolVar(&list, "list", false, "list servers")

	flag.Parse()

	log.Println("**************************************************************")
	log.Println("******** Parser started with following commands **************")
	log.Printf("**  interval %v", interval)
	log.Println("**  Distance Unit " + distanceUnit)
	log.Println("**  Host " + host)

	if showExternalIp {
		log.Println("**  showExternalIp: true")
	} else {

		log.Println("**  showExternalIp: false")
	}

	if saveToInfluxDb {
		log.Println("**  influxHost " + influxHost)
		log.Println("**  influxDB " + influxDB)
		log.Println("**  influxUser " + influxUser)

		if influxPwd == "DEFAULT" {
			log.Println("**  influxPwd DEFAULT")
		} else {
			log.Println("**  influxPwd is not Default")
		}
	}

	log.Println("**************************************************************")
	log.Println("**************************************************************")
	return Settings{interval, host, server, distanceUnit, list, keepProcessRunning, showExternalIp, includeHumanOutput, InfluxDbSettings{saveToInfluxDb, influxHost, influxUser, influxPwd, influxDB}, retryZeroValue, retryInterval}
}
