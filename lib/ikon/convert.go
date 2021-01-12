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
	}
	log.Fatalf("how convert %s", s)
	return 0
}
