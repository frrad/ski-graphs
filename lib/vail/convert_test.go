package vail

import (
	"testing"
	"time"

	"github.com/frrad/ski-graphs/lib/lift"
)

func TestConvertLift(t *testing.T) {
	losAngeles, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Error(err)
	}
	observationTime := time.Date(2021, time.January, 1, 9, 45, 0, 0, losAngeles)

	ins := []Lift{
		{
			Name:      "Cool Lift",
			State:     StateOpen,
			StartHour: "10:00",
			EndHour:   "16:00",
		},
		{
			Name:      "ASDF Lift",
			State:     StateOpen,
			StartHour: "08:00",
			EndHour:   "16:00",
		},
	}
	expectedOut := []lift.Lift{
		{
			MeasurementTime: observationTime,
			Name:            "Cool Lift",
			AreaName:        "Neat Area",
			Resort:          "Northstar",
			Status:          lift.StatusScheduled,
		},
		{
			MeasurementTime: observationTime,
			Name:            "ASDF Lift",
			AreaName:        "Neat Area",
			Resort:          "Northstar",
			Status:          lift.StatusOpen,
		},
	}

	for i := 0; i < len(ins); i++ {
		out := ins[i].AsLift(
			"Northstar", "Neat Area", observationTime,
		)

		if out != expectedOut[i] {
			t.Errorf(
				"%d\n%+v\nnot same as expectation\n%+v",
				i, out, expectedOut[i],
			)
		}
	}
}

func TestConvertState(t *testing.T) {
	for i := State(0); i < stateMax; i++ {
		tf := true
		_, err := i.AsStatus(tf)
		if err != nil {
			t.Errorf("error converting state %s with %t: %s", i, tf, err)
		}

		tf = false
		_, err = i.AsStatus(tf)
		if err != nil {
			t.Errorf("error converting state %s with %t: %s", i, tf, err)
		}

	}

}
