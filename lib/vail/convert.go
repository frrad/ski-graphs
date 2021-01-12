package vail

import (
	"github.com/frrad/ski-graphs/lib/lift"
)

func (l Lift) AsLift(resort string) lift.Lift {
	return lift.Lift{
		Name:            l.Name,
		Resort:          resort,
		AreaName:        l.SectorID,
		Status:          l.State.AsState(),
		WaitTimeSeconds: 60 * l.Wait,
	}
}
