package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api "github.com/influxdata/influxdb-client-go/v2/api"
	apiWrite "github.com/influxdata/influxdb-client-go/v2/api/write"
)

const (
	organization = "organization"
	bucket       = "bucket"
	token        = "zI_gNzqimDn58hwhA1HtiJaSmFpYkThP68zD23yGp8_Q8YzepH5nXasCi8eY5XJcCfF17u7Re18JEoc36UHeLw=="
	influxURL    = "http://localhost:8086"
)

func setupInfluxClient() (api.WriteAPI, func()) {
	client := influxdb2.NewClientWithOptions(influxURL, token, influxdb2.DefaultOptions())
	writeAPI := client.WriteAPI(organization, bucket)

	cleanup := func() {
		writeAPI.Flush()
		client.Close()
	}

	return writeAPI, cleanup
}

func main() {
	writeClient, cleanup := setupInfluxClient()
	defer cleanup()

	files, err := filepath.Glob("*.json")
	if err != nil {
		log.Fatal(err)
	}

	processFiles(files, writeClient)
}

func processFiles(files []string, influxClient api.WriteAPI) {
	for _, f := range files {
		x, err := parseFile(f)
		if err != nil {
			log.Fatalf("error parsing %s: %s", f, err)
		}

		for _, area := range x.Response.MountainAreas {
			for _, l := range area.Lifts {
				ps := pointFromLift(x.Time, x.Response.Name, l)

				for _, p := range ps {
					influxClient.WritePoint(p)
				}
			}
		}
	}
}

func pointFromLift(t time.Time, resort string, l Lift) []*apiWrite.Point {

	ans := []*apiWrite.Point{}

	for statusName, val := range l.Status.OneHot() {
		tags := map[string]string{
			"AreaName": l.MountainAreaName,
			"LiftName": l.Name,
			"Resort":   resort,
			"Status":   statusName,
		}
		fields := map[string]interface{}{"count": val}

		log.Println(tags, fields)

		ans = append(ans, influxdb2.NewPoint(
			"lift",
			tags,
			fields,
			t,
		))

	}

	if l.WaitTime.set {
		tags := map[string]string{
			"AreaName": l.MountainAreaName,
			"LiftName": l.Name,
			"Resort":   resort,
		}
		fields := map[string]interface{}{"wait": l.WaitTime.val}

		log.Println(tags, fields)

		ans = append(ans, influxdb2.NewPoint(
			"lift-wait",
			tags,
			fields,
			t,
		))
	}

	return ans
}

func parseFile(f string) (ResortTime, error) {
	rt, err := parseName(f)
	if err != nil {
		return ResortTime{}, err
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		return ResortTime{}, err
	}
	d := Response{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return ResortTime{}, err
	}

	rt.Response = d
	return rt, nil
}

type ResortTime struct {
	Resort   int
	Time     time.Time
	Response Response
}

func parseName(filename string) (ResortTime, error) {
	sps := strings.Split(filename, ".")
	if len(sps) != 3 {
		return ResortTime{}, fmt.Errorf("expected three parts %s", filename)
	}

	if sps[2] != "json" {
		return ResortTime{}, fmt.Errorf("expected last part to be json %s", filename)
	}

	resort, err := parseResort(sps[0])
	if err != nil {
		return ResortTime{}, err
	}

	t, err := time.Parse("2006-01-02 15:04:05-07:00", sps[1])
	if err != nil {
		return ResortTime{}, err
	}

	return ResortTime{
		Resort: resort,
		Time:   t,
	}, nil
}

func parseResort(resortStr string) (int, error) {
	sps := strings.Split(resortStr, "-")

	if len(sps) != 2 {
		return 0, fmt.Errorf("expected two parts %s", resortStr)
	}

	if sps[0] != "resort" {
		return 0, fmt.Errorf("expected resort founde %s", sps[0])
	}

	return strconv.Atoi(sps[1])
}
