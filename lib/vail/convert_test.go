package vail

import (
	"testing"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func TestConvertLift(t *testing.T) {
	in := Lift{
		Name:  "Cool Lift",
		State: StateClosed,
	}
	expectedOut := lift.Lift{
		Name:     "Cool Lift",
		AreaName: "Neat Area",
		Resort:   "Sweet Resort",
		Status:   lift.StatusOpen,
	}
	out := in.AsLift("Sweet Resort", "Neat Area", time.Time{})

	if out != expectedOut {
		t.Errorf("\n%+v\nnot same as expectation\n%+v", out, expectedOut)
	}

}
