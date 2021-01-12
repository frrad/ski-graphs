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
		Name:   "Cool Lift",
		Resort: "Sweet Resort",
		Status: lift.StatusOpen,
	}
	out := in.AsLift("Sweet Resort", time.Time{})

	if out != expectedOut {
		t.Errorf("\n%+v\nnot same as expectation\n%+v", out, expectedOut)
	}

}
