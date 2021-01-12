package lift

import (
	"log"
)

type Lift struct {
	Name            string
	Resort          string
	AreaName        string
	Status          Status
	WaitTimeSeconds *int
}

type Status int

const (
	StatusScheduled Status = iota
	StatusMax
)

func (s Status) OneHot() map[string]interface{} {
	ans := map[string]interface{}{}
	for i := Status(0); i < StatusMax; i++ {
		ans[i.String()] = 0
		if i == s {
			ans[i.String()] = 1
		}
	}
	return ans
}

func (s Status) String() string {
	switch s {
	case StatusScheduled:
		return "Scheduled"
	}

	log.Fatalf("how string %d", s)
	return ""
}