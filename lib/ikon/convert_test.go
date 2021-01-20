package ikon

import (
	"testing"
)

func TestConvertStatus(t *testing.T) {
	for i := LiftStatus(0); i < LiftStatusMax; i++ {
		i.AsStatus()
	}
}
