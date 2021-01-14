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
	StatusScheduled Status = iota // not open right now but will be in the morning
	StatusClosed                  // during the day "not open" at night "won't be open tomorrow"
	StatusOpen                    // open!
	statusMax                     // special max value
)

func (s Status) OneHot() map[string]interface{} {
	ans := map[string]interface{}{}
	for i := Status(0); i < statusMax; i++ {
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
	}

	log.Fatalf("how string %d", s)
	return ""
}
