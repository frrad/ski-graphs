package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api "github.com/influxdata/influxdb-client-go/v2/api"
	apiWrite "github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/frrad/ski-graphs/lib/ikon"
	"github.com/frrad/ski-graphs/lib/lift"
	"github.com/frrad/ski-graphs/lib/vail"
)

func processIkonFiles(files []string, influxClient api.WriteAPI) {
	for _, f := range files {
		x, err := parseName(f)
		if err != nil {
			log.Fatalf("error parsing name %s: %s", f, err)
		}

		resp, err := parseIkon(f)
		if err != nil {
			log.Fatalf("error parsing %s: %s", f, err)
		}

		for _, area := range resp.MountainAreas {
			for _, l := range area.Lifts {
				ps := pointFromLift(x.Time, l.AsLift(resp.Name))

				for _, p := range ps {
					influxClient.WritePoint(p)
				}
			}
		}
	}
}

func processEpicFiles(files []string, influxClient api.WriteAPI) {
	for _, f := range files {
		x, err := parseName(f)
		if err != nil {
			log.Fatalf("error parsing name %s: %s", f, err)
		}

		resp, err := parseEpic(f)
		if err != nil {
			log.Fatalf("error parsing %s: %s", f, err)
		}

		for _, sta := range resp.Data.Stations {
			for _, state := range sta.States {
				for _, lift := range state.Skilifts {

					ps := pointFromLift(x.Time, lift.AsLift(resp.Data.Name))

					for _, p := range ps {
						influxClient.WritePoint(p)
					}

				}
			}
		}
	}
}

func pointFromLift(t time.Time, l lift.Lift) []*apiWrite.Point {
	ans := []*apiWrite.Point{}

	for statusName, val := range l.Status.OneHot() {
		tags := map[string]string{
			"AreaName": l.AreaName,
			"LiftName": l.Name,
			"Resort":   l.Resort,
			"Status":   statusName,
		}
		fields := map[string]interface{}{"Count": val}
		log.Println(tags, fields)
		ans = append(ans, influxdb2.NewPoint(
			"lift-status-count",
			tags,
			fields,
			t,
		))
	}

	tags := map[string]string{
		"AreaName": l.AreaName,
		"LiftName": l.Name,
		"Resort":   l.Resort,
		"Status":   l.Status.String(),
	}
	fields := map[string]interface{}{"Status": l.Status.String()}
	log.Println(tags, fields)
	ans = append(ans, influxdb2.NewPoint(
		"lift-status",
		tags,
		fields,
		t,
	))

	if wt := l.WaitTimeSeconds; wt != nil {
		tags := map[string]string{
			"AreaName": l.AreaName,
			"LiftName": l.Name,
			"Resort":   l.Resort,
		}
		fields := map[string]interface{}{"Wait": *wt}
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

func parseIkon(f string) (ikon.Response, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return ikon.Response{}, err
	}
	d := ikon.Response{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return ikon.Response{}, err
	}

	return d, nil
}

func parseEpic(f string) (vail.Response, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return vail.Response{}, err
	}
	d := vail.Response{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return vail.Response{}, err
	}

	return d, nil
}

type ResortTime struct {
	Resort int
	Time   time.Time
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
