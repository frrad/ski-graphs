package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
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

func processIkonFiles(x ResortTime, b []byte) ([]*apiWrite.Point, error) {
	pts := []*apiWrite.Point{}

	resp := ikon.Response{}
	err := json.Unmarshal(b, &resp)
	if err != nil {
		return pts, err
	}

	for _, area := range resp.MountainAreas {
		for _, l := range area.Lifts {
			ps := pointFromLift(l.AsLift(x.Time, resp.Name))

			for _, p := range ps {
				pts = append(pts, p)
			}
		}
	}

	return pts, nil
}

type FileProcessor func(ResortTime, []byte) ([]*apiWrite.Point, error)

func processFiles(files []string, s *Seen, fu FileProcessor, influxClient api.WriteAPI) {
	skipped, processed := 0, 0

	for _, fileName := range files {
		if s.Saw(fileName) {
			skipped++
			continue
		}

		x, err := parseName(fileName)
		if err != nil {
			log.Fatalf("error parsing name %s: %s", fileName, err)
		}

		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalf("error reading file %s: %s", fileName, err)
		}
		pts, err := fu(x, b)
		if err != nil {
			log.Fatalf("error processing file %s: %s", fileName, err)
		}
		for _, p := range pts {
			influxClient.WritePoint(p)
		}

		processed++
		s.Mark(fileName)
	}

	log.Printf("%d/%d files skipped\n", skipped, len(files))
	log.Printf("%d/%d files processed\n", processed, len(files))
}

func processEpicFile(x ResortTime, b []byte) ([]*apiWrite.Point, error) {
	pts := []*apiWrite.Point{}

	resp := vail.Response{}
	err := json.Unmarshal(b, &resp)
	if err != nil {
		return pts, err
	}

	for _, sta := range resp.Data.Stations {
		for _, state := range sta.States {
			for _, lift := range state.Skilifts {

				ps := pointFromLift(lift.AsLift(resp.Data.Name, sta.Name, x.Time))

				for _, p := range ps {
					pts = append(pts, p)
				}

			}
		}

	}

	return pts, nil
}

func pointFromLift(l lift.Lift) []*apiWrite.Point {
	ans := []*apiWrite.Point{}

	for statusName, val := range l.Status.OneHot() {
		tags := map[string]string{
			"AreaName": l.AreaName,
			"LiftName": l.Name,
			"Resort":   l.Resort,
			"Status":   statusName,
		}
		fields := map[string]interface{}{"Count": val}
		ans = append(ans, influxdb2.NewPoint(
			"lift-status-count",
			tags,
			fields,
			l.MeasurementTime,
		))
	}

	tags := map[string]string{
		"AreaName": l.AreaName,
		"LiftName": l.Name,
		"Resort":   l.Resort,
		"Status":   l.Status.String(),
	}
	fields := map[string]interface{}{"Status": l.Status.String()}
	ans = append(ans, influxdb2.NewPoint(
		"lift-status",
		tags,
		fields,
		l.MeasurementTime,
	))

	if wt := l.WaitTime; wt != nil {
		tags := map[string]string{
			"AreaName": l.AreaName,
			"LiftName": l.Name,
			"Resort":   l.Resort,
		}
		fields := map[string]interface{}{"Wait": wt.Seconds()}
		ans = append(ans, influxdb2.NewPoint(
			"lift-wait",
			tags,
			fields,
			l.MeasurementTime,
		))
	}

	return ans
}

type ResortTime struct {
	Resort int
	Time   time.Time
}

func parseName(filepath string) (ResortTime, error) {
	_, filename := path.Split(filepath)

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
