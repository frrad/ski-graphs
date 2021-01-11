package ikon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	dir, err := filepath.Glob("testinput/*.json")
	if err != nil {
		t.Error(err)
	}

	for _, fileName := range dir {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			t.Error(err)
		}

		RoundtripBytes(t, b)
	}
}

func RoundtripBytes(t *testing.T, b []byte) {
	data := Response{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		t.Error(err)
	}

	encodeTo := &bytes.Buffer{}
	enc := json.NewEncoder(encodeTo)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	err = enc.Encode(data)
	if err != nil {
		t.Error(err)
	}

	outB := encodeTo.Bytes()

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
