package vail

import (
	"testing"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func TestConvertLift(t *testing.T) {
	ins := []Lift{{
		Name:  "Cool Lift",
		State: StateClosed,
	}}
	expectedOut := []lift.Lift{{
		Name:     "Cool Lift",
		AreaName: "Neat Area",
		Resort:   "Northstar",
		Status:   lift.StatusClosed,
	}}

	for i := 0; i < len(ins); i++ {
		out := ins[i].AsLift("Northstar", "Neat Area", time.Time{})

		if out != expectedOut[i] {
			t.Errorf(
				"%d\n%+v\nnot same as expectation\n%+v",
				i, out, expectedOut[i],
			)
		}
	}
}
