package main

import (
	"strings"

	"github.com/samherrmann/gowrap/gotools"
)

func NewConfig() *Config {
	goos, err := gotools.GetGoOS()
	panicIf(err)

	goarch, err := gotools.GetGoArch()
	panicIf(err)

	t := &[]Target{
		Target(goos + "-" + goarch),
	}
	return &Config{
		Targets: t,
	}
}

type Config struct {
	Targets *[]Target `json:"targets"`
}

type Target string

func (t *Target) Parse() (goos string, goarch string) {
	s := strings.Split(string(*t), "-")
	return s[0], s[1]
}
