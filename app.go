package main

import (
	"context"
	"errors"
	"github.com/m-lab/ndt7-client-go"
	"github.com/m-lab/ndt7-client-go/spec"
	"log"
	"net"
	"os"
	"quadstingray/speedtest-influxdb/model"
	"quadstingray/speedtest-influxdb/model/speedtest"
	"sort"
	"time"
)

const (
	clientName     = "speedtest-influxdb"
	clientVersion  = "1.1.0"
	defaultTimeout = 60 * time.Second
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

		if settings.IncludeHumanReadable {
			log.Printf("speed test started")
		}

		stats, err := runTest(settings)

		if err != nil {
			time.Sleep(time.Duration(settings.RetryInterval) * time.Second)
		} else {
			if !settings.RetryZeroValue || (stats.Down_Mbs != 0 || stats.Up_Mbs != 0) {
				if settings.InfluxDbSettings.Use_Influx {
					go model.SaveToInfluxDb(stats, settings)
				}
				if settings.IncludeHumanReadable {
					log.Printf("sleep for %v seconds", settings.Interval)
				}
				time.Sleep(time.Duration(settings.Interval) * time.Second)
			} else {
				time.Sleep(time.Duration(settings.RetryInterval) * time.Second)
			}
		}

	}
}

func listServers() {

	allServers, err := speedtest.ListServer()
	if err != nil {
		log.Printf("error creating client: %v", err)
	}

	sort.Slice(allServers, func(i, j int) bool {
		return allServers[i].Country < allServers[j].Country
	})

	for _, v := range allServers {
		log.Printf("County: %v | Location: %v | ServerId: %v | UplinkSpeed: %v | Roundrobin: %v", v.Country, v.City, v.Site, v.UplinkSpeed, v.Roundrobin)
	}
}

func runTest(settings model.Settings) (model.SpeedTestStatistics, error) {
	//geoClient2, _ := model.LocateUser()
	//log.Printf("%v", geoClient2)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var output speedtest.OutputType

	if settings.IncludeHumanReadable {
		output = speedtest.NewHumanReadable()
	} else {
		output = speedtest.SilentOutput{}
	}

	var r = speedtest.TestRunner{
		ndt7.NewClient(clientName, clientVersion),
		output,
	}

	if settings.Server != "" {
		r.Client.Server = "ndt-mlab3-" + settings.Server + ".mlab-oti.measurement-lab.org"
	}

	var code int
	code += r.RunDownload(ctx)
	code += r.RunUpload(ctx)

	if code != 0 {
		code = 0
		log.Printf("No Connection to Server %v restart with search new NDT7 Sever", r.Client.Server)
		r.Client.Server = ""
		code += r.RunDownload(ctx)
		code += r.RunUpload(ctx)
		if code != 0 {
			return model.SpeedTestStatistics{
				Client:             model.ClientInformations{},
				Server:             model.Server{},
				Ping:               0,
				Down_Mbs:           0,
				Up_Mbs:             0,
				DownRetransPercent: 0,
			}, errors.New("server not reachable")
		}
	}

	s := makeSummary(r.Client.FQDN, r.Client.Results())
	r.Output.OnSummary(s)

	geoClient, _ := model.CheckIpLocation(s.ClientIP)
	time.Sleep(time.Duration(1) * time.Second)
	geoSever, _ := speedtest.FindServerByFQDN(s.ServerFQDN)

	var distance float64
	if geoSever.Lat == 0 && geoSever.Lon == 0 || geoClient.Lat == 0 && geoClient.Lon == 0 {
		distance = 0
	} else {
		distance = model.Distance(geoSever.Lat, geoSever.Lon, geoClient.Lat, geoClient.Lon, settings.DistanceUnit)
	}

	return model.SpeedTestStatistics{
		model.ClientInformations{
			ExternalIp: s.ClientIP,
			Provider:   geoClient.Org,
			Coordinate: model.Coordinate{
				geoClient.Lat,
				geoClient.Lon,
			},
		},
		model.Server{
			URL:      s.ServerFQDN,
			Lat:      geoSever.Lat,
			Lon:      geoSever.Lon,
			Name:     s.ServerFQDN,
			Country:  geoSever.Country,
			City:     geoSever.City,
			Distance: distance,
			Latency:  0,
		},
		s.MinRTT.Value,
		s.Download.Value,
		s.Upload.Value,
		s.DownloadRetrans.Value,
	}, nil
}

func makeSummary(FQDN string, results map[spec.TestKind]*ndt7.LatestMeasurements) *speedtest.Summary {

	s := speedtest.NewSummary(FQDN)

	if results[spec.TestDownload] != nil &&
		results[spec.TestDownload].ConnectionInfo != nil {
		// Get UUID, ClientIP and ServerIP from ConnectionInfo.
		s.DownloadUUID = results[spec.TestDownload].ConnectionInfo.UUID

		clientIP, _, err := net.SplitHostPort(results[spec.TestDownload].ConnectionInfo.Client)
		if err == nil {
			s.ClientIP = clientIP
		}

		serverIP, _, err := net.SplitHostPort(results[spec.TestDownload].ConnectionInfo.Server)
		if err == nil {
			s.ServerIP = serverIP
		}
	}

	if dl, ok := results[spec.TestDownload]; ok {
		if dl.Client.AppInfo != nil && dl.Client.AppInfo.ElapsedTime > 0 {
			elapsed := float64(dl.Client.AppInfo.ElapsedTime) / 1e06
			s.Download = speedtest.ValueUnitPair{
				Value: (8.0 * float64(dl.Client.AppInfo.NumBytes)) /
					elapsed / (1000.0 * 1000.0),
				Unit: "Mbit/s",
			}
		}
		if dl.Server.TCPInfo != nil {
			if dl.Server.TCPInfo.BytesSent > 0 {
				s.DownloadRetrans = speedtest.ValueUnitPair{
					Value: float64(dl.Server.TCPInfo.BytesRetrans) / float64(dl.Server.TCPInfo.BytesSent) * 100,
					Unit:  "%",
				}
			}
			s.MinRTT = speedtest.ValueUnitPair{
				Value: float64(dl.Server.TCPInfo.MinRTT) / 1000,
				Unit:  "ms",
			}
		}
	}
	// Upload comes from the client-side Measurement during the upload test.
	if ul, ok := results[spec.TestUpload]; ok {
		if ul.Client.AppInfo != nil && ul.Client.AppInfo.ElapsedTime > 0 {
			elapsed := float64(ul.Client.AppInfo.ElapsedTime) / 1e06
			s.Upload = speedtest.ValueUnitPair{
				Value: (8.0 * float64(ul.Client.AppInfo.NumBytes)) /
					elapsed / (1000.0 * 1000.0),
				Unit: "Mbit/s",
			}
		}
	}

	return s
}
