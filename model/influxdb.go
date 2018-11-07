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
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  settings.InfluxDbSettings.Db,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"host": settings.Host}

	fields := map[string]interface{}{
		"download_mbs": statistics.Down_Mbs,
		"upload_mbs":   statistics.Up_Mbs,
		"ping":         statistics.Ping,
		"distance":     statistics.Distance,
		"serverid":     statistics.ServerId,
		"location":     statistics.Location,
	}

	if settings.ShowMyIp {
		fields = map[string]interface{}{
			"download_mbs": statistics.Down_Mbs,
			"upload_mbs":   statistics.Up_Mbs,
			"ping":         statistics.Ping,
			"distance":     statistics.Distance,
			"serverid":     statistics.ServerId,
			"location":     statistics.Location,
			"external_ip":  statistics.External_Ip,
		}
	}

	pt, err := client.NewPoint("speedtest", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatalf("could not write to influx Db. check connection to %v and Db %s with user %v with pwd %s", settings.InfluxDbSettings.Db_Url, settings.InfluxDbSettings.Db, settings.InfluxDbSettings.Username, settings.InfluxDbSettings.Password)
	}
}
