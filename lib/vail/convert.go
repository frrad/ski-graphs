package vail

import (
	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort string) lift.Lift {

	ans := lift.Lift{
		Name:     l.Name,
		Resort:   resort,
		AreaName: l.SectorID,
		Status:   l.State.AsStatus(),
	}

	if waitTimeMinutes, ok := l.WaitMinutes.Get(); ok {
		waitTimeMinutes *= 60
		ans.WaitTimeSeconds = &waitTimeMinutes
	}

	return ans
}

func (s State) AsStatus() lift.Status {
	return lift.StatusScheduled
}
