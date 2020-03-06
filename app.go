package main

import (
	"github.com/dchest/uniuri"
	"github.com/kylegrantlucas/speedtest"
	"github.com/kylegrantlucas/speedtest/coords"
	"github.com/kylegrantlucas/speedtest/http"
	"log"
	"os"
	"quadstingray/speedtest-influxdb/model"
	"sort"
	"time"
)

func main() {
	settings := model.Parser()

	if settings.ListServers {
		listServers()
		if settings.KeepProcessRunning {
			for true {
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
		os.Exit(0)
	}

	for true {
		log.Printf("speed test started")
		stats, err := runTest(settings)

		if err != nil {
			time.Sleep(time.Duration(5) * time.Minute)
		} else {
			if settings.InfluxDbSettings.Use_Influx {
				go model.SaveToInfluxDb(stats, settings)
			}
			if settings.ShowMyIp {
				log.Printf("Ping: %3.2f ms | Download: %3.2f Mbps | Upload: %3.2f Mbps | External_Ip: %s", stats.Ping, stats.Down_Mbs, stats.Up_Mbs, stats.Client.ExternalIp)
			} else {
				log.Printf("Ping: %3.2f ms | Download: %3.2f Mbps | Upload: %3.2f Mbps", stats.Ping, stats.Down_Mbs, stats.Up_Mbs)
			}
			log.Printf("sleep for %v seconds", settings.Interval)
			time.Sleep(time.Duration(settings.Interval) * time.Second)
		}

	}
}

func listServers() {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		log.Printf("error creating client: %v", err)
	}

	allServers, err := client.HTTPClient.GetServers()
	if err != nil {
		log.Printf("error creating client: %v", err)
	}

	sort.Slice(allServers, func(i, j int) bool {
		return allServers[i].Country < allServers[j].Country
	})

	for _, v := range allServers {
		log.Printf("County: %v | Location: %v | ServerId: %v | Sponsor: %v", v.Country, v.Name, v.ID, v.Sponsor)
	}
}

func speedTestClient(settings model.Settings) (*speedtest.Client, error) {
	config := &http.SpeedtestConfig{
		ConfigURL:       "http://c.speedtest.net/speedtest-config.php?x=" + uniuri.New(),
		ServersURL:      "http://c.speedtest.net/speedtest-servers-static.php?x=" + uniuri.New(),
		AlgoType:        settings.AlgoType,
		NumClosest:      3,
		NumLatencyTests: 3,
		UserAgent:       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.21 Safari/537.36",
	}
	timeOut := time.Hour
	return speedtest.NewClient(config, speedtest.DefaultDLSizes, speedtest.DefaultULSizes, timeOut)
}

func runTest(settings model.Settings) (model.SpeedTestStatistics, error) {
	client, err := speedTestClient(settings)
	if err != nil {
		log.Printf("error creating client: %v", err)
	}

	// Pass an empty string to select the fastest server
	server, err := client.GetServer(settings.Server)
	if err != nil {
		log.Printf("error getting server: %v", err)
	}

	dmbps, err := client.Download(server)
	if err != nil {
		log.Printf("error getting download: %v", err)
	}

	umbps, err := client.Upload(server)
	if err != nil {
		log.Printf("error getting upload: %v", err)
	}

	clientInformations := model.ClientInformations{client.HTTPClient.Config.IP, client.HTTPClient.Config.Isp, coords.Coordinate{client.HTTPClient.Config.Lat, client.HTTPClient.Config.Lon}}

	result := model.SpeedTestStatistics{clientInformations, server, server.Latency, dmbps, umbps}
	return result, nil
}
