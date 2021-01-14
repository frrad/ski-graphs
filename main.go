package main

import (
	"flag"
	"log"
	"path/filepath"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api "github.com/influxdata/influxdb-client-go/v2/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	influxToken := flag.String("influx-token", "", "influxDB token")
	influxURL := flag.String("influx-url", "http://localhost:8086", "influxDB url")
	influxOrg := flag.String("influx-org", "", "influxDB url")
	influxBucket := flag.String("influx-bucket", "", "influxDB bucket")

	epicGlob := flag.String("epic-glob", "", "epic glob")
	ikonGlob := flag.String("ikon-glob", "", "ikon glob")

	seenFile := flag.String("seen-file", "", "seen file")
	flag.Parse()

	seen, err := NewSeen(*seenFile)
	if err != nil {
		log.Fatal(err)
	}

	writeClient, cleanup := setupInfluxClient(*influxURL, *influxToken, *influxOrg, *influxBucket)
	defer cleanup()

	go func(writeClient api.WriteAPI) {
		errChan := writeClient.Errors()
		for err := range errChan {
			if err != nil {
				log.Fatal(err)
			}
		}
	}(writeClient)

	if *ikonGlob != "" {
		files, err := filepath.Glob(*ikonGlob)
		if err != nil {
			log.Fatal(err)
		}
		processFiles(files, seen, processIkonFiles, writeClient)
	}

	if *epicGlob != "" {
		files, err := filepath.Glob(*epicGlob)
		if err != nil {
			log.Fatal(err)
		}
		processFiles(files, seen, processEpicFiles, writeClient)
	}
}

func setupInfluxClient(url, token, org, bucket string) (api.WriteAPI, func()) {
	client := influxdb2.NewClientWithOptions(url, token, influxdb2.DefaultOptions())
	writeAPI := client.WriteAPI(org, bucket)

	cleanup := func() {
		writeAPI.Flush()
		client.Close()
	}

	return writeAPI, cleanup
}
