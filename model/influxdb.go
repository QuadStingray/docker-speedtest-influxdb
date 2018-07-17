package model

import (
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
)

func SaveToInfluxDb(statistics SpeedTestStatistics, settings Settings) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     settings.InfluxDbSettings.dbUrl,
		Username: settings.InfluxDbSettings.username,
		Password: settings.InfluxDbSettings.password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  settings.InfluxDbSettings.db,
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

	pt, err := client.NewPoint("speedtest", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatalf("could not write to influx db. check connection to %v and db %s with user %v with pwd %s", settings.InfluxDbSettings.dbUrl, settings.InfluxDbSettings.db, settings.InfluxDbSettings.username, settings.InfluxDbSettings.password)
	}
}
