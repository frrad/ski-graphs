package vail

import (
	"log"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort string, t time.Time) lift.Lift {
	ans := lift.Lift{
		Name:     l.Name,
		Resort:   resort,
		AreaName: l.SectorID,
		Status:   l.State.AsStatus(t, l.StartHour, l.EndHour),
	}

	if waitTimeMinutes, ok := l.WaitMinutes.Get(); ok {
		waitTimeMinutes *= 60
		ans.WaitTimeSeconds = &waitTimeMinutes
	}

	return ans
}

func (s State) AsStatus(now time.Time, startHour, endHour string) lift.Status {
	switch s {
	case StateOpen:
		return lift.StatusScheduled
	case StateClosed:
		return lift.StatusOpen
	}

	log.Fatalf("how convert %s", s)
	return -1
}
