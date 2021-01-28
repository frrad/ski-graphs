package ikon

import (
	"log"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(t time.Time, resort string) lift.Lift {
	resp := lift.Lift{
		MeasurementTime: t,
		Name:            l.Name,
		Resort:          resort,
		AreaName:        l.MountainAreaName,
		Status:          l.Status.AsStatus(),
	}

	if ans, ok := l.WaitTime.Get(); ok {
		wts := time.Second * time.Duration(ans)

		resp.WaitTime = &wts
	}

	return resp
}

func (s LiftStatus) AsStatus() lift.Status {
	switch s {
	case LiftStatusScheduled:
		return lift.StatusScheduled
	case LiftStatusClosed:
		return lift.StatusClosed
	case LiftStatusWindHold:
		return lift.StatusClosed
	case LiftStatusWindClosure:
		return lift.StatusClosed
	case LiftStatusPatrolHold:
		return lift.StatusClosed
	case LiftStatusOpen:
		return lift.StatusOpen
	case LiftStatusMechanicalHold:
		return lift.StatusClosed
	case LiftStatusAnticipatedWeatherImpact:
		return lift.StatusClosed
	case LiftStatusDelayed:
		return lift.StatusClosed
	case LiftStatusHold:
		return lift.StatusClosed
	case LiftStatusMechanicalClosure:
		return lift.StatusClosed
	case LiftStatusMidStationOnly:
		return lift.StatusOpen
	case LiftStatusDownloadOnly:
		return lift.StatusClosed
	case LiftStatusWindyConditions:
		return lift.StatusOpen
	}
	log.Fatalf("don't know how to convert ikon status \"%s\"", s)
	return 0
}
