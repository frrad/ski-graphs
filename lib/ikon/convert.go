package ikon

import (
	"log"

	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort string) lift.Lift {
	resp := lift.Lift{
		Name:     l.Name,
		Resort:   resort,
		AreaName: l.MountainAreaName,
		Status:   l.Status.AsStatus(),
	}

	if ans, ok := l.WaitTime.Get(); ok {
		resp.WaitTimeSeconds = &ans
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
