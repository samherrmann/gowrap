package main

import "strings"

func NewConfig() *Config {
	t := &[]Target{
		Target(goOS() + "-" + goArch()),
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
