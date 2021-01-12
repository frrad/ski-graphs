package ikon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/frrad/ski-graphs/lib/util"
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

		util.RoundtripBytes(t, b, roundTrip)
	}
}

func roundTrip(b []byte) ([]byte, error) {
	data := Response{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return []byte{}, err
	}

	encodeTo := &bytes.Buffer{}
	enc := json.NewEncoder(encodeTo)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	err = enc.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	outB := encodeTo.Bytes()
	return outB, nil
}
