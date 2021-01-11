package util

import (
	"testing"
)

func RoundtripBytes(t *testing.T, b []byte, f func([]byte) ([]byte, error)) {
	outB, err := f(b)
	if err != nil {
		t.Errorf("failed to round trip bytes: %s", err)
	}

	CompareBytes(t, b, outB)
}

func CompareBytes(t *testing.T, a, b []byte) {
	if len(a) != len(b) {
		t.Errorf("different lens %d != %d", len(a), len(b))
	}

	min := len(a)
	if len(b) < min {
		min = len(b)
	}

	brokenIx := -1

	for i := 0; i < min; i++ {
		if a[i] != b[i] {
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

		t.Errorf("snippet\na=\n%s\nb=\n%s", string(a[from:to]), string(b[from:to]))
	}
}
