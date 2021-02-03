package model

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

func SaveToInfluxDb(statistics SpeedTestStatistics, settings Settings) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     settings.InfluxDbSettings.Db_Url,
		Username: settings.InfluxDbSettings.Username,
		Password: settings.InfluxDbSettings.Password,
	})
	if err != nil {
		log.Printf("error creating http client: %v", err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  settings.InfluxDbSettings.Db,
		Precision: "s",
	})
	if err != nil {
		log.Printf("error new batch point: %v", err)
	}

	// Create a point and add to batch
	tags := map[string]string{"host": settings.Host}

	fields := map[string]interface{}{
		"download_mbs":   statistics.Down_Mbs,
		"upload_mbs":     statistics.Up_Mbs,
		"ping":           statistics.Ping,
		"distance":       statistics.Server.Distance,
		"serverid":       statistics.Server.Name,
		"location":       statistics.Server.City + ", " + statistics.Server.Country,
		"clientProvider": statistics.Client.Provider,
	}

	if settings.ShowMyIp {
		fields = map[string]interface{}{
			"download_mbs":   statistics.Down_Mbs,
			"upload_mbs":     statistics.Up_Mbs,
			"ping":           statistics.Ping,
			"distance":       statistics.Server.Distance,
			"serverid":       statistics.Server.Name,
			"location":       statistics.Server.City + ", " + statistics.Server.Country,
			"external_ip":    statistics.Client.ExternalIp,
			"clientProvider": statistics.Client.Provider,
		}
	}

	pt, err := client.NewPoint("speedtest", tags, fields, time.Now())
	if err != nil {
		log.Printf("error creating messure point: %v", err)
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Printf("could not write to influx Db. check connection to %v and Db %s with user %v with pwd %s (error: %s)", settings.InfluxDbSettings.Db_Url, settings.InfluxDbSettings.Db, settings.InfluxDbSettings.Username, settings.InfluxDbSettings.Password, err)
		time.Sleep(time.Duration(10) * time.Second)
		SaveToInfluxDb(statistics, settings)
	}
}
