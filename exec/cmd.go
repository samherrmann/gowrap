package exec

import (
	"os/exec"
	"strings"
)

// Command returns a Cmd struct
func Command(name string, args ...string) *Cmd {
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

// OutputLine calls Output on the command and
// returns the first line of the standard output.
func (c *Cmd) OutputLine() (string, error) {
	out, err := c.Cmd.Output()
	return strings.TrimRight(string(out), "\n"), err
}
