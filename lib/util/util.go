package util

import (
	"testing"
)

func RoundtripBytes(t *testing.T, before []byte, f func([]byte) ([]byte, error)) {
	after, err := f(before)
	if err != nil {
		t.Errorf("failed to round trip bytes: %s", err)
	}

	CompareBytes(t, before, after)
}

func CompareBytes(t *testing.T, before, after []byte) {
	if len(before) != len(after) {
		t.Errorf("different lens %d != %d", len(before), len(after))
	}

	min := len(before)
	if len(after) < min {
		min = len(after)
	}

	brokenIx := -1

	for i := 0; i < min; i++ {
		if before[i] != after[i] {
			t.Errorf("first mismatch at offset %d", i)
			brokenIx = i
			break
		}
	}

	if brokenIx != -1 {

		width := 50
		from, to := brokenIx-width, brokenIx+width
		if from < 0 {
			from = 0
		}
		if to < 0 {
			from = 0
		}
		if from > min {
			from = min
		}
		if to > min {
			from = min
		}

		t.Errorf("snippet\nbefore=\n%s\nafter=\n%s", string(before[from:to]), string(after[from:to]))
	}
}
