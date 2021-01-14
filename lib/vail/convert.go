package vail

import (
	"fmt"
	"log"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort, area string, t time.Time) lift.Lift {
	loc := resortLocation(resort)

	isOpen := isOpen(t, loc, l.StartHour, l.EndHour)
	stat, err := l.State.AsStatus(isOpen)
	if err != nil {
		log.Fatal(err)
	}

	ans := lift.Lift{
		Name:            l.Name,
		Resort:          resort,
		AreaName:        area,
		Status:          stat,
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

func isOpen(measuredAt time.Time, resortLoc *time.Location, startHour, endHour string) bool {
	timeOfDayStr := measuredAt.In(resortLoc).Format("15:04")
	return startHour <= timeOfDayStr && timeOfDayStr <= endHour
}

func (liftState State) AsStatus(resortOpen bool) (lift.Status, error) {
	switch {
	case liftState == StateOpen && resortOpen:
		return lift.StatusOpen, nil
	case liftState == StateOpen && !resortOpen:
		return lift.StatusScheduled, nil
	case liftState == StateClosed && resortOpen:
		return lift.StatusClosed, nil
	case liftState == StateClosed && !resortOpen:
		return lift.StatusClosed, nil
	case liftState == StateP && resortOpen:
		return lift.StatusUnknown, nil
	case liftState == StateP && !resortOpen:
		return lift.StatusUnknown, nil
	}

	return -1, fmt.Errorf("don't know how to convert %s", liftState)
}
