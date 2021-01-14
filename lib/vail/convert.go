package vail

import (
	"log"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort, area string, t time.Time) lift.Lift {
	loc := resortLocation(resort)

	ans := lift.Lift{
		Name:            l.Name,
		Resort:          resort,
		AreaName:        area,
		Status:          l.State.AsStatus(t, loc, l.StartHour, l.EndHour),
		MeasurementTime: t,
	}

	if waitTimeMinutes, ok := l.WaitMinutes.Get(); ok {
		waitTime := time.Minute * time.Duration(waitTimeMinutes)
		ans.WaitTime = &waitTime
	}

	return ans
}

func resortLocation(resortName string) *time.Location {
	tzName := ""
	switch resortName {
	case "Northstar":
		tzName = "America/Los_Angeles"
	default:
		log.Fatalf("unknown tz for %s", resortName)
	}

	// hardcode for now
	l, err := time.LoadLocation(tzName)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

func (liftState State) AsStatus(measuredAt time.Time, resortLoc *time.Location, startHour, endHour string) lift.Status {
	timeOfDayStr := measuredAt.In(resortLoc).Format("15:04")
	resortOpen := startHour <= timeOfDayStr && timeOfDayStr <= endHour

	switch {
	case liftState == StateOpen && resortOpen:
		return lift.StatusOpen
	case liftState == StateOpen && !resortOpen:
		return lift.StatusScheduled
	case liftState == StateClosed && resortOpen:
		return lift.StatusClosed
	case liftState == StateClosed && !resortOpen:
		return lift.StatusClosed
	}

	log.Fatalf("how convert %s", liftState)
	return -1
}
