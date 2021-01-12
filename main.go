package main

import (
	"flag"
	"log"
	"path/filepath"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api "github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	bucket = "bucket"
)

func main() {
	influxToken := flag.String("influx-token", "", "influxDB token")
	influxURL := flag.String("influx-url", "http://localhost:8086", "influxDB url")
	influxOrg := flag.String("influx-org", "", "influxDB url")
	flag.Parse()

	writeClient, cleanup := setupInfluxClient(*influxURL, *influxToken, *influxOrg)
	defer cleanup()

	files, err := filepath.Glob("*.json")
	if err != nil {
		log.Fatal(err)
	}

	go func(writeClient api.WriteAPI) {
		errChan := writeClient.Errors()
		for err := range errChan {
			if err != nil {
				log.Fatal(err)
			}
		}
	}(writeClient)

	processFiles(files, writeClient)

}

func setupInfluxClient(url, token, org string) (api.WriteAPI, func()) {
	client := influxdb2.NewClientWithOptions(url, token, influxdb2.DefaultOptions())
	writeAPI := client.WriteAPI(org, bucket)

	cleanup := func() {
		writeAPI.Flush()
		client.Close()
	}

	return writeAPI, cleanup
}
