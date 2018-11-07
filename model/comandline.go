package model

import (
	"flag"
	"log"
)

func Parser() Settings {
	var interval int
	var server string
	var host string
	var influxHost string
	var influxDB string
	var influxPwd string
	var influxUser string
	var list bool
	var showExternalIp bool
	var saveToInfluxDb bool

	flag.IntVar(&interval, "interval", 3600, "seconds between statistics import")

	flag.StringVar(&server, "server", "", "speedtest.net server")
	flag.StringVar(&host, "host", "", "host where the speedetest is running")
	flag.StringVar(&influxHost, "influxHost", "http://influxdb:8086", "host of your influxdb instance")
	flag.StringVar(&influxDB, "influxDB", "rspamd", "influxdb database")
	flag.StringVar(&influxUser, "influxUser", "DEFAULT", "influxdb Username")
	flag.StringVar(&influxPwd, "influxPwd", "DEFAULT", "influxdb Password")

	flag.BoolVar(&saveToInfluxDb, "saveToInfluxDb", false, "save to influxdb")
	flag.BoolVar(&list, "list", false, "list servers")
	flag.BoolVar(&showExternalIp, "showExternalIp", true, "save and show external Ip of docker host")

	flag.Parse()

	log.Println("**************************************************************")
	log.Println("******** Parser started with following commands **************")
	log.Printf("**  interval %v", interval)
	log.Println("**  server " + server)
	log.Println("**  host " + host)

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
	return Settings{interval, host, server, list, showExternalIp, InfluxDbSettings{saveToInfluxDb, influxHost, influxUser, influxPwd, influxDB}}
}
