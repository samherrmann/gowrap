package main

import "os/exec"

// cmd returns a Cmd struct
func cmd(name string, args ...string) *Cmd {
	cmd := exec.Command(name, args...)
	c := &Cmd{}
	c.Cmd = cmd
	return c
}

// Cmd extends exec.Cmd
// from Go's standard library.
type Cmd struct {
	*exec.Cmd
}

// Run calls Go's cmd.Run() and panics
// if an error occurs.
func (c *Cmd) Run() {
	err := c.Cmd.Run()
	panicIf(err)
}

// Output calls Go's cmd.Output() and panics
// if an error occurs.
func (c *Cmd) Output() []byte {
	out, err := c.Cmd.Output()
	panicIf(err)
	return out
}
