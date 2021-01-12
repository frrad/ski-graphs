package lift

import (
	"log"
	"time"
)

type Lift struct {
	MeasurementTime time.Time
	Name            string
	Resort          string
	AreaName        string
	Status          Status
	WaitTime        *time.Duration
}

type Status int

const (
	StatusScheduled Status = iota
	StatusOpen
	StatusClosed
	StatusHold
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
	case StatusOpen:
		return "Open"
	case StatusClosed:
		return "Closed"
	case StatusHold:
		return "Hold"
	}

	log.Fatalf("how string %d", s)
	return ""
}
