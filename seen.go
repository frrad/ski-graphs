package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Seen struct {
	files       map[string]struct{}
	backingPath string
}

func (s *Seen) Saw(path string) bool {
	_, ok := s.files[path]
	return ok
}

func (s *Seen) Mark(path string) {
	s.files[path] = struct{}{}

	bs, err := json.MarshalIndent(s.files, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(s.backingPath, bs, perms)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	perms     = 0744
	emptyFile = "{}"
)

func NewSeen(backingPath string) (*Seen, error) {
	bs, err := ioutil.ReadFile(backingPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("can't open backing file for seen-file: %w", err)
		}

		err := ioutil.WriteFile(backingPath, []byte(emptyFile), perms)
		if err != nil {
			return nil, fmt.Errorf("error creating initial seen-file \"%s\": %w", backingPath, err)
		}
		bs = []byte(emptyFile)
	}

	s := map[string]struct{}{}

	err = json.Unmarshal(bs, &s)
	if err != nil {
		return nil, err
	}

	return &Seen{
		files:       s,
		backingPath: backingPath,
	}, nil
}
