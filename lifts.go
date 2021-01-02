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
)

func main() {
	files, err := filepath.Glob("*.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		x, err := parseFile(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, area := range x.Response.MountainAreas {
			for _, l := range area.Lifts {
				fmt.Println(x.Time, l.Name, l.Status, l.UpdateDate, l.WaitTime)
			}
		}
	}
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
