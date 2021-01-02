package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"testing"
)

//go:embed testinput/*
var testFiles embed.FS

const dirName = "testinput"

func TestAbs(t *testing.T) {
	dir, err := testFiles.ReadDir(dirName)
	if err != nil {
		t.Error(err)
	}

	for _, x := range dir {
		fileName := x.Name()
		path := fmt.Sprintf("%s/%s", dirName, fileName)

		b, err := testFiles.ReadFile(path)
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

	outB, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Error(err)
	}

	CompareBytes(t, b, append(outB, []byte("\n")...))
}

func CompareBytes(t *testing.T, a, b []byte) {
	if len(a) != len(b) {
		t.Fatalf("different lens %d != %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Error(a[i], b[i])
		}
	}
}
