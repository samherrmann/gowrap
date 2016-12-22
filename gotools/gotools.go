package gotools

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
)

const (
	envVarKeyGOOS       = "GOOS"
	envVarKeyGOARCH     = "GOARCH"
	envVarKeyGOHOSTOS   = "GOHOSTOS"
	envVarKeyGOHOSTARCH = "GOHOSTARCH"
)

// GoTools represents a Go tools environment
type GoTools struct{}

// New returns an instance of a GoTools struct
func New() *GoTools {
	return new(GoTools)
}

// GOHOSTOS returns the value of GOHOSTOS
func (gt *GoTools) GOHOSTOS() (string, error) {
	return goEnvVar(envVarKeyGOHOSTOS)
}

// GOHOSTARCH returns the value of GOHOSTOS
func (gt *GoTools) GOHOSTARCH() (string, error) {
	return goEnvVar(envVarKeyGOHOSTARCH)
}

// GOOS returns the value of GOOS
func (gt *GoTools) GOOS() (string, error) {
	return goEnvVar(envVarKeyGOOS)
}

// SetGOOS sets the value of GOOS
func (gt *GoTools) SetGOOS(goos string) error {
	return os.Setenv(envVarKeyGOOS, goos)
}

// GOARCH returns the value of GOARCH
func (gt *GoTools) GOARCH() (string, error) {
	return goEnvVar(envVarKeyGOARCH)
}

// SetGOARCH sets the value of GOARCH
func (gt *GoTools) SetGOARCH(goarch string) error {
	return os.Setenv(envVarKeyGOARCH, goarch)
}

// SupportedPlatforms returns a list of all the target operating
// systems and compilation architectures.
//
// This function requires at least Go version go1.7 to be installed.
// The 'go tool dist list' command that this function calls was introduced
// in the following commit:
// https://github.com/golang/go/commit/c3ecded729214abf8a146902741cd6f9d257f68c
func (gt *GoTools) SupportedPlatforms() (*[]Platform, error) {
	p := &[]Platform{}

	out, err := exec.Command("go", "tool", "dist", "list", "-json").CombinedOutput()
	if err != nil {
		return p, errors.New("Error while getting list of supported platforms: " +
			err.Error() +
			". Are you running at least Go version go1.7?")
	}

	err = json.Unmarshal(out, p)
	return p, err
}

// Generate executes the command "go generate"
func (gt *GoTools) Generate() error {
	_, err := cmdOutput("go", "generate")
	return err
}

// Build executes the command "go build" for the desired
// target OS and architecture, and writes the generated
// executable to the 'outDir' directory.
func (gt *GoTools) Build(args ...string) error {
	args = append([]string{"build"}, args...)
	_, err := cmdOutput("go", args...)
	return err
}

// ExeSuffix returns ".exe" if the GOOS
// environment variable is set to "windows".
func (gt *GoTools) ExeSuffix() (string, error) {
	goos, err := gt.GOOS()
	if err != nil {
		return "", err
	}
	if goos == "windows" {
		return ".exe", nil
	}
	return "", nil
}

// goEnvVar returns the value of the provided Go
// environment variable
func goEnvVar(key string) (string, error) {
	return cmdOutput("go", "env", key)
}

// cmdOutput executes the command specified by name and
// returns the first line of the standard output.
func cmdOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(out))
	}

	return strings.TrimRight(string(out), "\n"), err
}
