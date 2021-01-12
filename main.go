package main

import (
	"log"
	"path/filepath"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api "github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	organization = "organization"
	bucket       = "bucket"
	token        = "zI_gNzqimDn58hwhA1HtiJaSmFpYkThP68zD23yGp8_Q8YzepH5nXasCi8eY5XJcCfF17u7Re18JEoc36UHeLw=="
	influxURL    = "http://localhost:8086"
)

func main() {
	writeClient, cleanup := setupInfluxClient()
	defer cleanup()

	files, err := filepath.Glob("*.json")
	if err != nil {
		log.Fatal(err)
	}

	processFiles(files, writeClient)
}

func setupInfluxClient() (api.WriteAPI, func()) {
	client := influxdb2.NewClientWithOptions(influxURL, token, influxdb2.DefaultOptions())
	writeAPI := client.WriteAPI(organization, bucket)

	cleanup := func() {
		writeAPI.Flush()
		client.Close()
	}

	return writeAPI, cleanup
}
